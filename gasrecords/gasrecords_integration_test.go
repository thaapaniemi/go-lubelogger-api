package gasrecords_test

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/thaapaniemi/go-lubelogger-api"
	"github.com/thaapaniemi/go-lubelogger-api/document"
	"github.com/thaapaniemi/go-lubelogger-api/gasrecords"
	"github.com/thaapaniemi/go-lubelogger-api/vehicles"
)

func TestGasRecordFilesAgainstLiveServer(t *testing.T) {
	if os.Getenv("LUBELOGGER_INTEGRATION") != "1" {
		t.Skip("set LUBELOGGER_INTEGRATION=1 to run live server integration test")
	}

	// Set timeout for live API calls
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	c := lubelogger.NewClient("http://127.0.0.1:8080", "test", "1234")

	vv, err := vehicles.GetRecords(ctx, c)
	if err != nil {
		t.Fatalf("GetRecords vehicles failed: %v", err)
	}
	if len(vv) == 0 {
		t.Fatal("expected at least one vehicle in test instance")
	}

	vehicleID := vv[0].ID

	// Generate unique test data per run
	uniqueID := fmt.Sprintf("%d-%d", time.Now().Unix(), time.Now().Nanosecond())
	testDocKey := fmt.Sprintf("gasrecord-files-live-test-%s.txt", uniqueID)
	testOdometer := int64(999000) // Fixed odometer, uniqueness via description/notes
	testDescription := fmt.Sprintf("gas files integration test [%s]", uniqueID)
	testNotes := fmt.Sprintf("created by integration test [%s]", uniqueID)

	doc := document.Document{
		Key:         testDocKey,
		Description: "integration test upload",
		Type:        "file",
		Src:         []byte("gas record attachment integration test"),
	}

	location, err := doc.Upload(ctx, c)
	if err != nil {
		t.Fatalf("document upload failed: %v", err)
	}
	if location == "" {
		t.Fatal("document upload returned empty location")
	}

	record := gasrecords.GasRecord{
		Date:         time.Now().UTC(),
		Odometer:     testOdometer,
		FuelConsumed: 9.25,
		IsFillToFull: true,
		MissedFuelUp: false,
		Description:  testDescription,
		Notes:        testNotes,
		Cost:         23.45,
		Files: []gasrecords.UploadedFile{
			{
				Name:      doc.Key,
				Location:  location,
				IsPending: false,
			},
		},
	}

	if err := record.Add(ctx, c, vehicleID); err != nil {
		t.Fatalf("gas record add with files failed: %v", err)
	}

	// Register cleanup defer immediately after successful Add, before any assertions
	// that could abort the test (which would prevent the defer from running).
	var created *gasrecords.GasRecord
	defer func() {
		if created != nil {
			if err := created.Delete(ctx, c); err != nil {
				t.Errorf("gas record delete failed: %v", err)
			}
		}
	}()

	records, err := gasrecords.GetRecords(ctx, c, vehicleID)
	if err != nil {
		t.Fatalf("gas record get failed: %v", err)
	}

	for i := range records {
		if records[i].Description == testDescription && records[i].Notes == testNotes && records[i].Odometer == testOdometer {
			created = &records[i]
			break
		}
	}
	if created == nil {
		t.Fatalf("created gas record not found in live server response (expected description=%q, notes=%q, odometer=%d)", testDescription, testNotes, testOdometer)
	}

	if len(created.Files) != 1 {
		t.Fatalf("expected 1 attached file, got %d", len(created.Files))
	}
	if created.Files[0].Name != doc.Key {
		t.Fatalf("unexpected attached file name: %q", created.Files[0].Name)
	}
	if created.Files[0].Location != location {
		t.Fatalf("unexpected attached file location: %q", created.Files[0].Location)
	}

	created.Notes = fmt.Sprintf("updated by integration test [%s]", uniqueID)
	created.Files = []gasrecords.UploadedFile{
		{
			Name:      doc.Key,
			Location:  location,
			IsPending: false,
		},
	}
	if err := created.Update(ctx, c); err != nil {
		t.Fatalf("gas record update with files failed: %v", err)
	}

	updatedRecords, err := gasrecords.GetRecords(ctx, c, vehicleID)
	if err != nil {
		t.Fatalf("gas record get after update failed: %v", err)
	}

	var updated *gasrecords.GasRecord
	for i := range updatedRecords {
		if updatedRecords[i].ID == created.ID {
			updated = &updatedRecords[i]
			break
		}
	}
	if updated == nil {
		t.Fatal("updated gas record not found in live server response")
	}
	if updated.Notes != fmt.Sprintf("updated by integration test [%s]", uniqueID) {
		t.Fatalf("gas record update did not persist notes, got %q", updated.Notes)
	}
	if len(updated.Files) != 1 {
		t.Fatalf("expected 1 attached file after update, got %d", len(updated.Files))
	}
	if updated.Files[0].Location != location {
		t.Fatalf("attached file location changed after update: %q", updated.Files[0].Location)
	}
}
