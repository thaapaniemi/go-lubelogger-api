package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/thaapaniemi/go-lubelogger-api"
	"github.com/thaapaniemi/go-lubelogger-api/calendar"
	"github.com/thaapaniemi/go-lubelogger-api/client"
	"github.com/thaapaniemi/go-lubelogger-api/debuglog"
	"github.com/thaapaniemi/go-lubelogger-api/document"
	"github.com/thaapaniemi/go-lubelogger-api/gasrecords"

	"github.com/thaapaniemi/go-lubelogger-api/odometer"
	"github.com/thaapaniemi/go-lubelogger-api/reminders"
	"github.com/thaapaniemi/go-lubelogger-api/repairrecords"
	"github.com/thaapaniemi/go-lubelogger-api/servicerecord"
	"github.com/thaapaniemi/go-lubelogger-api/taxrecords"
	"github.com/thaapaniemi/go-lubelogger-api/upgraderecords"
	"github.com/thaapaniemi/go-lubelogger-api/vehicleinfo"
	"github.com/thaapaniemi/go-lubelogger-api/vehicles"
)

func main() {
	debuglog.Enabled = true
	c := lubelogger.NewClient("http://127.0.0.1:8080", "test", "1234")
	vehicleId := testVehicleInfo(c)
	testOdometerRecords(c, vehicleId)
	testServiceRecords(c, vehicleId)
	testRepairRecords(c, vehicleId)
	testUpgradeRecords(c, vehicleId)
	testTaxRecords(c, vehicleId)
	testGasRecords(c, vehicleId)
	testCalendar(c)
	testReminders(c, vehicleId)
	testUpload(c)
}

func testVehicleInfo(c client.Client) int64 {
	fmt.Println("--- testVehicleInfo ---")
	ctx := context.Background()
	vv, err := vehicles.GetRecords(ctx, c)
	if err != nil {
		panic(err)
	}

	fmt.Printf("found %d vechicles\n", len(vv))

	firstVehicleId := vv[0].ID

	ww, err := vehicleinfo.GetRecords(ctx, c, firstVehicleId)
	if err != nil {
		panic(err)
	}

	fmt.Printf("found %d vechicleInfos\n", len(ww))
	return firstVehicleId
}

func testOdometerRecords(c client.Client, vehicleId int64) {
	fmt.Println("--- testOdometerRecords ---")
	ctx := context.Background()

	new := odometer.OdometerRecord{
		Odometer: 999000,
		Date:     time.Now(),
		Notes:    "test note",
	}
	err := new.Add(context.Background(), c, vehicleId)
	if err != nil {
		panic(err)
	}

	latest, err := odometer.GetLatestValue(ctx, c, vehicleId)
	if err != nil {
		panic(err)
	}
	fmt.Println(latest)

	v, err := odometer.GetRecords(ctx, c, vehicleId)
	if err != nil {
		panic(err)
	}

	data, err := json.Marshal(&v)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))

	last := v[len(v)-1]
	last.Notes = last.Notes + " updated"
	err = last.Update(ctx, c)
	if err != nil {
		panic(err)
	}

	err = last.Delete(ctx, c)
	if err != nil {
		panic(err)
	}
}

func testServiceRecords(c client.Client, vehicleId int64) {
	fmt.Println("--- testServiceRecords ---")
	ctx := context.Background()

	new := servicerecord.ServiceRecord{
		Odometer:    999000,
		Date:        time.Now(),
		Description: "description",
		Notes:       "test note",
		Cost:        25.99,
	}
	err := new.Add(ctx, c, vehicleId)
	if err != nil {
		panic(err)
	}

	v, err := servicerecord.GetRecords(ctx, c, vehicleId)
	if err != nil {
		panic(err)
	}

	data, err := json.Marshal(&v)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))

	last := v[len(v)-1]
	last.Description = last.Description + " updated"
	err = last.Update(ctx, c)
	if err != nil {
		panic(err)
	}

	err = last.Delete(ctx, c)
	if err != nil {
		panic(err)
	}
}

func testRepairRecords(c client.Client, vehicleId int64) {
	fmt.Println("--- testRepairRecords ---")
	ctx := context.Background()

	new := repairrecords.RepairRecord{
		Odometer:    999000,
		Date:        time.Now(),
		Description: "description",
		Notes:       "test note",
		Cost:        31.99,
	}
	err := new.Add(ctx, c, vehicleId)
	if err != nil {
		panic(err)
	}

	v, err := repairrecords.GetRecords(ctx, c, vehicleId)
	if err != nil {
		panic(err)
	}

	data, err := json.Marshal(&v)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))

	last := v[len(v)-1]
	last.Description = last.Description + " updated"
	err = last.Update(ctx, c)
	if err != nil {
		panic(err)
	}

	err = last.Delete(ctx, c)
	if err != nil {
		panic(err)
	}
}

func testUpgradeRecords(c client.Client, vehicleId int64) {
	fmt.Println("--- testUpgradeRecords ---")
	ctx := context.Background()

	new := upgraderecords.UpgradeRecord{
		Odometer:    999000,
		Date:        time.Now(),
		Description: "description",
		Notes:       "test note",
		Cost:        55.44,
	}
	err := new.Add(ctx, c, vehicleId)
	if err != nil {
		panic(err)
	}

	v, err := upgraderecords.GetRecords(ctx, c, vehicleId)
	if err != nil {
		panic(err)
	}

	data, err := json.Marshal(&v)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))

	last := v[len(v)-1]
	last.Description = last.Description + " updated"
	err = last.Update(ctx, c)
	if err != nil {
		panic(err)
	}

	err = last.Delete(ctx, c)
	if err != nil {
		panic(err)
	}
}

func testTaxRecords(c client.Client, vehicleId int64) {
	fmt.Println("--- testTaxRecords ---")
	ctx := context.Background()

	new := taxrecords.TaxRecord{
		Date:        time.Now(),
		Description: "description",
		Notes:       "test note",
		Cost:        99.00,
	}
	err := new.Add(ctx, c, vehicleId)
	if err != nil {
		panic(err)
	}

	v, err := taxrecords.GetRecords(ctx, c, vehicleId)
	if err != nil {
		panic(err)
	}

	data, err := json.Marshal(&v)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))

	last := v[len(v)-1]
	last.Description = last.Description + " updated"
	err = last.Update(ctx, c)
	if err != nil {
		panic(err)
	}

	err = last.Delete(ctx, c)
	if err != nil {
		panic(err)
	}
}

func testGasRecords(c client.Client, vehicleId int64) {
	fmt.Println("--- testGasRecords ---")
	ctx := context.Background()

	new := gasrecords.GasRecord{
		Odometer:     999000,
		FuelConsumed: 7,
		IsFillToFull: true,
		MissedFuelUp: true,
		Cost:         15,
		Date:         time.Now(),
		Description:  "description",
		Notes:        "test note",
	}
	err := new.Add(ctx, c, vehicleId)
	if err != nil {
		panic(err)
	}

	v, err := gasrecords.GetRecords(ctx, c, vehicleId)
	if err != nil {
		panic(err)
	}

	data, err := json.Marshal(&v)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))

	last := v[len(v)-1]
	last.Description = last.Description + " updated"
	err = last.Update(ctx, c)
	if err != nil {
		panic(err)
	}

	err = last.Delete(ctx, c)
	if err != nil {
		panic(err)
	}
}

func testCalendar(c client.Client) {
	fmt.Println("--- testCalendar ---")
	ctx := context.Background()

	data, err := calendar.GetCalendar(ctx, c)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))

}

func testReminders(c client.Client, vehicleId int64) {
	fmt.Println("--- testReminders ---")
	ctx := context.Background()

	v, err := reminders.GetRecords(ctx, c, vehicleId)
	if err != nil {
		panic(err)
	}

	data, err := json.Marshal(&v)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))

	err = reminders.SendReminderEmails(ctx, c, []reminders.Urgency{reminders.URGENCY_NOT_URGENT, reminders.URGENCY_URGENT, reminders.URGENCY_VERY_URGENT, reminders.URGENCY_PAST_DUE})
	if err != nil {
		panic(err)
	}
}

func testUpload(c client.Client) {
	fmt.Println("--- testUpload ---")
	ctx := context.Background()

	item := document.Document{
		Key:         "small.txt",
		Description: "sample",
		Type:        "file",
		Src:         []byte("small file"),
	}

	location, err := item.Upload(ctx, c)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(location))

}
