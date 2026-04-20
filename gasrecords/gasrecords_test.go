package gasrecords

import (
	"encoding/json"
	"strings"
	"testing"
	"time"
)

// TestGasRecordFilesMarshaling verifies that GasRecord.Files are marshaled
// as objects with name, location, and isPending fields, not as strings.
func TestGasRecordFilesMarshaling(t *testing.T) {
	record := GasRecord{
		ID:           1,
		Date:         time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC),
		Odometer:     50000,
		FuelConsumed: 10.5,
		FuelEconomy:  25.0,
		IsFillToFull: true,
		MissedFuelUp: false,
		Description:  "Tank refill",
		Notes:        "Full tank",
		Cost:         45.50,
		Tags:         "routine",
		Files: []UploadedFile{
			{
				Name:      "receipt.pdf",
				Location:  "https://example.com/files/receipt.pdf",
				IsPending: false,
			},
			{
				Name:      "invoice.pdf",
				Location:  "https://example.com/files/invoice.pdf",
				IsPending: false,
			},
		},
	}

	payload, err := json.Marshal(&record)
	if err != nil {
		t.Fatalf("Failed to marshal gas record: %v", err)
	}

	var result map[string]interface{}
	err = json.Unmarshal(payload, &result)
	if err != nil {
		t.Fatalf("Failed to unmarshal marshaled payload: %v", err)
	}

	filesInterface, ok := result["files"]
	if !ok {
		t.Fatal("files field not present in marshaled payload")
	}

	filesArray, ok := filesInterface.([]interface{})
	if !ok {
		t.Fatalf("files is not an array, got type: %T", filesInterface)
	}

	if len(filesArray) != 2 {
		t.Fatalf("Expected 2 files, got %d", len(filesArray))
	}

	// Check first file object structure
	firstFile, ok := filesArray[0].(map[string]interface{})
	if !ok {
		t.Fatalf("First file is not an object, got type: %T", filesArray[0])
	}

	if name, ok := firstFile["name"].(string); !ok || name != "receipt.pdf" {
		t.Errorf("First file name mismatch: %v", firstFile["name"])
	}

	if location, ok := firstFile["location"].(string); !ok || location != "https://example.com/files/receipt.pdf" {
		t.Errorf("First file location mismatch: %v", firstFile["location"])
	}

	if isPending, ok := firstFile["isPending"].(bool); !ok || isPending != false {
		t.Errorf("First file isPending mismatch: %v", firstFile["isPending"])
	}
}

// TestGasRecordFilesUnmarshaling verifies that API responses with file objects
// are correctly converted to GasRecord without panics or data loss.
func TestGasRecordFilesUnmarshaling(t *testing.T) {
	// Simulate API response with file objects
	apiResponse := []map[string]interface{}{
		{
			"id":           json.Number("1"),
			"date":         "2024-01-15",
			"odometer":     json.Number("50000"),
			"fuelConsumed": json.Number("10.5"),
			"fuelEconomy":  json.Number("25.0"),
			"isFillToFull": true,
			"missedFuelUp": false,
			"description":  "Tank refill",
			"notes":        "Full tank",
			"cost":         json.Number("45.50"),
			"tags":         "routine",
			"extraFields":  nil,
			"files": []interface{}{
				map[string]interface{}{
					"name":      "receipt.pdf",
					"location":  "https://example.com/files/receipt.pdf",
					"isPending": false,
				},
				map[string]interface{}{
					"name":      "invoice.pdf",
					"location":  "https://example.com/files/invoice.pdf",
					"isPending": true,
				},
			},
		},
	}

	records, err := convertAll(apiResponse)
	if err != nil {
		t.Fatalf("convertAll failed: %v", err)
	}

	if len(records) != 1 {
		t.Fatalf("Expected 1 record, got %d", len(records))
	}

	record := records[0]
	if len(record.Files) != 2 {
		t.Fatalf("Expected 2 files, got %d", len(record.Files))
	}

	if record.Files[0].Name != "receipt.pdf" {
		t.Errorf("First file name mismatch: %s", record.Files[0].Name)
	}

	if record.Files[0].Location != "https://example.com/files/receipt.pdf" {
		t.Errorf("First file location mismatch: %s", record.Files[0].Location)
	}

	if record.Files[0].IsPending != false {
		t.Errorf("First file isPending mismatch: %v", record.Files[0].IsPending)
	}

	if record.Files[1].IsPending != true {
		t.Errorf("Second file isPending mismatch: %v", record.Files[1].IsPending)
	}
}

// TestGasRecordNilAndEmptyFiles verifies that nil or empty files arrays
// are handled correctly without adding spurious fields.
func TestGasRecordNilAndEmptyFiles(t *testing.T) {
	// Test with nil files
	recordNil := GasRecord{
		ID:       1,
		Date:     time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC),
		Odometer: 50000,
		Files:    nil,
	}

	payload, err := json.Marshal(&recordNil)
	if err != nil {
		t.Fatalf("Failed to marshal gas record with nil files: %v", err)
	}

	var result map[string]interface{}
	err = json.Unmarshal(payload, &result)
	if err != nil {
		t.Fatalf("Failed to unmarshal marshaled payload: %v", err)
	}

	// files field should not be present when nil (due to omitempty tag)
	if _, ok := result["files"]; ok && result["files"] != nil {
		t.Errorf("Expected no files field for nil files, but got: %v", result["files"])
	}

	// Test with empty files array
	recordEmpty := GasRecord{
		ID:       1,
		Date:     time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC),
		Odometer: 50000,
		Files:    []UploadedFile{},
	}

	payload, err = json.Marshal(&recordEmpty)
	if err != nil {
		t.Fatalf("Failed to marshal gas record with empty files: %v", err)
	}

	err = json.Unmarshal(payload, &result)
	if err != nil {
		t.Fatalf("Failed to unmarshal marshaled payload: %v", err)
	}

	// files field should not be present when empty (due to omitempty tag)
	if _, ok := result["files"]; ok {
		t.Errorf("Expected no files field for empty files array, but field is present with value: %v", result["files"])
	}

	// Test unmarshaling with nil files in API response
	apiResponse := []map[string]interface{}{
		{
			"id":           json.Number("1"),
			"date":         "2024-01-15",
			"odometer":     json.Number("50000"),
			"fuelConsumed": json.Number("0"),
			"fuelEconomy":  json.Number("0"),
			"isFillToFull": false,
			"missedFuelUp": false,
			"description":  "",
			"notes":        "",
			"cost":         json.Number("0"),
			"tags":         "",
			"extraFields":  nil,
			"files":        nil,
		},
	}

	records, err := convertAll(apiResponse)
	if err != nil {
		t.Fatalf("convertAll failed with nil files: %v", err)
	}

	if len(records) != 1 {
		t.Fatalf("Expected 1 record, got %d", len(records))
	}

	if records[0].Files != nil && len(records[0].Files) > 0 {
		t.Errorf("Expected nil/empty files for API nil files, got: %v", records[0].Files)
	}
}

// TestGasRecordMalformedAttachmentDetection verifies that callers can detect
// conversion failures when attachment payloads have invalid shapes.
func TestGasRecordMalformedAttachmentDetection(t *testing.T) {
	tests := []struct {
		name    string
		input   []map[string]interface{}
		wantErr bool
		errMsg  string
	}{
		{
			name: "wrong_type_name_field",
			input: []map[string]interface{}{
				{
					"id":           json.Number("1"),
					"date":         "2024-01-15",
					"odometer":     json.Number("50000"),
					"isFillToFull": false,
					"missedFuelUp": false,
					"cost":         json.Number("0"),
					"files": []interface{}{
						map[string]interface{}{
							"name":      123, // wrong type
							"location":  "https://example.com/test.pdf",
							"isPending": false,
						},
					},
				},
			},
			wantErr: true,
			errMsg:  "unmarshal",
		},
		{
			name: "missing_required_name_field",
			input: []map[string]interface{}{
				{
					"id":           json.Number("1"),
					"date":         "2024-01-15",
					"odometer":     json.Number("50000"),
					"isFillToFull": false,
					"missedFuelUp": false,
					"cost":         json.Number("0"),
					"files": []interface{}{
						map[string]interface{}{
							"location":  "https://example.com/test.pdf",
							"isPending": false,
						},
					},
				},
			},
			wantErr: true,
			errMsg:  "missing required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := convertAll(tt.input)

			if tt.wantErr && err == nil {
				t.Fatalf("Expected error but got none")
			}
			if !tt.wantErr && err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}
			if tt.wantErr && !strings.Contains(strings.ToLower(err.Error()), strings.ToLower(tt.errMsg)) {
				t.Fatalf("Error message should contain '%s', got: %s", tt.errMsg, err.Error())
			}
		})
	}
}

// TestGasRecordMalformedFilesHandling verifies that malformed attachment field types
// are detected and return errors instead of panicking or silently losing data.
func TestGasRecordMalformedFilesHandling(t *testing.T) {
	tests := []struct {
		name    string
		input   []map[string]interface{}
		wantErr bool
	}{
		{
			name: "missing_name_field",
			input: []map[string]interface{}{
				{
					"id":           json.Number("1"),
					"date":         "2024-01-15",
					"odometer":     json.Number("50000"),
					"isFillToFull": false,
					"missedFuelUp": false,
					"cost":         json.Number("0"),
					"files": []interface{}{
						map[string]interface{}{
							// name is missing
							"location":  "https://example.com/files/test.pdf",
							"isPending": false,
						},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "wrong_type_name_string",
			input: []map[string]interface{}{
				{
					"id":           json.Number("1"),
					"date":         "2024-01-15",
					"odometer":     json.Number("50000"),
					"isFillToFull": false,
					"missedFuelUp": false,
					"cost":         json.Number("0"),
					"files": []interface{}{
						map[string]interface{}{
							"name":      123, // should be string
							"location":  "https://example.com/files/test.pdf",
							"isPending": false,
						},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "wrong_type_isPending_bool",
			input: []map[string]interface{}{
				{
					"id":           json.Number("1"),
					"date":         "2024-01-15",
					"odometer":     json.Number("50000"),
					"isFillToFull": false,
					"missedFuelUp": false,
					"cost":         json.Number("0"),
					"files": []interface{}{
						map[string]interface{}{
							"name":      "test.pdf",
							"location":  "https://example.com/files/test.pdf",
							"isPending": "yes", // should be bool
						},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "file_not_object",
			input: []map[string]interface{}{
				{
					"id":           json.Number("1"),
					"date":         "2024-01-15",
					"odometer":     json.Number("50000"),
					"isFillToFull": false,
					"missedFuelUp": false,
					"cost":         json.Number("0"),
					"files": []interface{}{
						"this is not an object",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "files_not_array",
			input: []map[string]interface{}{
				{
					"id":           json.Number("1"),
					"date":         "2024-01-15",
					"odometer":     json.Number("50000"),
					"isFillToFull": false,
					"missedFuelUp": false,
					"cost":         json.Number("0"),
					"files":        "not an array",
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			records, err := convertAll(tt.input)

			if tt.wantErr && err == nil {
				t.Errorf("convertAll should return error for %s, but got nil", tt.name)
			}

			if !tt.wantErr && err != nil {
				t.Errorf("convertAll should not return error for %s, but got: %v", tt.name, err)
			}

			// When error is returned, records should be nil
			if tt.wantErr && records != nil {
				t.Errorf("convertAll should return nil records on error for %s, but got: %v", tt.name, records)
			}
		})
	}
}

// TestEmptyFilesArrayPreservation verifies that empty files array from API
// is preserved as non-nil empty slice, distinguishing from null/absent files.
func TestEmptyFilesArrayPreservation(t *testing.T) {
	// Test empty array is preserved (non-nil)
	apiResponseEmptyArray := []map[string]interface{}{
		{
			"id":           json.Number("1"),
			"date":         "2024-01-15",
			"odometer":     json.Number("50000"),
			"fuelConsumed": json.Number("0"),
			"fuelEconomy":  json.Number("0"),
			"isFillToFull": false,
			"missedFuelUp": false,
			"description":  "",
			"notes":        "",
			"cost":         json.Number("0"),
			"tags":         "",
			"extraFields":  nil,
			"files":        []interface{}{}, // Empty array from API
		},
	}

	records, err := convertAll(apiResponseEmptyArray)
	if err != nil {
		t.Fatalf("convertAll failed with empty files array: %v", err)
	}

	if len(records) != 1 {
		t.Fatalf("Expected 1 record, got %d", len(records))
	}

	if records[0].Files == nil {
		t.Errorf("Expected non-nil empty slice for empty files array from API, got nil")
	}

	if records[0].Files != nil && len(records[0].Files) != 0 {
		t.Errorf("Expected empty files slice, got %d files", len(records[0].Files))
	}

	// Test null files is also handled (returns nil)
	apiResponseNullFiles := []map[string]interface{}{
		{
			"id":           json.Number("2"),
			"date":         "2024-01-15",
			"odometer":     json.Number("50000"),
			"fuelConsumed": json.Number("0"),
			"fuelEconomy":  json.Number("0"),
			"isFillToFull": false,
			"missedFuelUp": false,
			"description":  "",
			"notes":        "",
			"cost":         json.Number("0"),
			"tags":         "",
			"extraFields":  nil,
			"files":        nil, // Null files from API
		},
	}

	records, err = convertAll(apiResponseNullFiles)
	if err != nil {
		t.Fatalf("convertAll failed with nil files: %v", err)
	}

	if len(records) != 1 {
		t.Fatalf("Expected 1 record, got %d", len(records))
	}

	if records[0].Files != nil {
		t.Errorf("Expected nil for null files from API, got: %v", records[0].Files)
	}

	// Verify they are distinguishable
	emptyArrayRecord := apiResponseEmptyArray[0]
	nullRecord := apiResponseNullFiles[0]

	emptyRecs, _ := convertAll([]map[string]interface{}{emptyArrayRecord})
	nullRecs, _ := convertAll([]map[string]interface{}{nullRecord})

	isEmptyNil := emptyRecs[0].Files == nil
	isNullNil := nullRecs[0].Files == nil

	if isEmptyNil == isNullNil {
		t.Errorf("Empty array and null should be distinguishable: empty=%v, null=%v", isEmptyNil, isNullNil)
	}
}

// TestUploadedFileEmptyFieldsSerialization verifies that UploadedFile.Name
// and UploadedFile.Location with empty strings are always present in JSON,
// not omitted due to omitempty directive.
func TestUploadedFileEmptyFieldsSerialization(t *testing.T) {
	record := GasRecord{
		ID:       1,
		Date:     time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC),
		Odometer: 50000,
		Files: []UploadedFile{
			{
				Name:      "", // Empty Name
				Location:  "", // Empty Location
				IsPending: false,
			},
		},
	}

	payload, err := json.Marshal(&record)
	if err != nil {
		t.Fatalf("Failed to marshal gas record: %v", err)
	}

	var result map[string]interface{}
	err = json.Unmarshal(payload, &result)
	if err != nil {
		t.Fatalf("Failed to unmarshal marshaled payload: %v", err)
	}

	filesInterface, ok := result["files"]
	if !ok {
		t.Fatal("files field not present in marshaled payload")
	}

	filesArray, ok := filesInterface.([]interface{})
	if !ok {
		t.Fatalf("files is not an array, got type: %T", filesInterface)
	}

	if len(filesArray) != 1 {
		t.Fatalf("Expected 1 file, got %d", len(filesArray))
	}

	firstFile, ok := filesArray[0].(map[string]interface{})
	if !ok {
		t.Fatalf("First file is not an object, got type: %T", filesArray[0])
	}

	// Verify "name" field is present with empty string value
	name, hasName := firstFile["name"]
	if !hasName {
		t.Errorf("Expected 'name' field to be present in JSON, but it was omitted")
	}
	if name != "" {
		t.Errorf("Expected 'name' to be empty string, got: %v", name)
	}

	// Verify "location" field is present with empty string value
	location, hasLocation := firstFile["location"]
	if !hasLocation {
		t.Errorf("Expected 'location' field to be present in JSON, but it was omitted")
	}
	if location != "" {
		t.Errorf("Expected 'location' to be empty string, got: %v", location)
	}
}

// TestErrorWrappingWithPercenW verifies that errors.Is can unwrap
// errors returned from convertAll due to proper %w wrapping.
func TestErrorWrappingWithPercentW(t *testing.T) {
	apiResponse := []map[string]interface{}{
		{
			"id":           json.Number("1"),
			"date":         "2024-01-15",
			"odometer":     json.Number("50000"),
			"isFillToFull": false,
			"missedFuelUp": false,
			"cost":         json.Number("0"),
			"files": []interface{}{
				map[string]interface{}{
					"name":      123, // Wrong type will trigger ConvertUploadedFiles error
					"location":  "https://example.com/test.pdf",
					"isPending": false,
				},
			},
		},
	}

	_, err := convertAll(apiResponse)
	if err == nil {
		t.Fatal("Expected error from convertAll but got nil")
	}

	// Verify error message structure (wrapping occurred)
	errStr := err.Error()
	if !strings.Contains(strings.ToLower(errStr), "convert record") {
		t.Errorf("Error message should show conversion failure: %s", errStr)
	}

	// Verify wrapped error chain is present
	if !strings.Contains(errStr, "failed to convert") {
		t.Errorf("Expected wrapped error chain, got: %s", errStr)
	}
}
