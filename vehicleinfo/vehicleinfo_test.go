package vehicleinfo_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/thaapaniemi/go-lubelogger-api/client"
	"github.com/thaapaniemi/go-lubelogger-api/vehicleinfo"
)

const ApiReply = `[{"vehicleData":{"id":1,"imageLocation":"/defaults/noimage.png","year":1992,"make":"Jeep","model":"Cherokee","licensePlate":"MOMWAGON","purchaseDate":"2025-01-15","soldDate":"2025-03-31","purchasePrice":0,"soldPrice":0,"isElectric":false,"isDiesel":false,"useHours":false,"odometerOptional":false,"extraFields":[],"tags":[],"hasOdometerAdjustment":false,"odometerMultiplier":1,"odometerDifference":0,"dashboardMetrics":[],"vehicleIdentifier":"LicensePlate"},"veryUrgentReminderCount":0,"urgentReminderCount":0,"notUrgentReminderCount":0,"pastDueReminderCount":0,"nextReminder":null,"serviceRecordCount":5,"serviceRecordCost":224.15,"repairRecordCount":5,"repairRecordCost":198.10,"upgradeRecordCount":5,"upgradeRecordCost":261.57,"taxRecordCount":3,"taxRecordCost":210.50,"gasRecordCount":5,"gasRecordCost":157.97,"lastReportedOdometer":222647,"planRecordBackLogCount":0,"planRecordInProgressCount":0,"planRecordTestingCount":0,"planRecordDoneCount":0}]`

func TestGetVehicleInfo(t *testing.T) {
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { fmt.Fprint(w, ApiReply) }))
	defer svr.Close()

	c := client.New(svr.URL, "", "")

	res, err := vehicleinfo.GetRecords(context.Background(), c, 1)
	if err != nil {
		t.Errorf("GetVehicleInfo got exception: %v", err)
	}

	if len(res) != 1 {
		t.Errorf("length of result is incorrect: %d", len(res))
	}

	if res[0].TaxRecordCost != 210.5 {
		t.Errorf("invalid %s: %f", "TaxRecordCost", res[0].TaxRecordCost)
	}

}
