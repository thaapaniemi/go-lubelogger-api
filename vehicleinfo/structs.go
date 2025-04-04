package vehicleinfo

import (
	"github.com/thaapaniemi/go-lubelogger-api/parser"
	"github.com/thaapaniemi/go-lubelogger-api/reminders"
	"github.com/thaapaniemi/go-lubelogger-api/vehicles"
)

type VehicleInfo struct {
	VehicleData               vehicles.VehicleData `json:"vehicleData"`
	VeryUrgentReminderCount   int64                `json:"veryUrgentReminderCount"`
	UrgentReminderCount       int64                `json:"urgentReminderCount"`
	NotUrgentReminderCount    int64                `json:"notUrgentReminderCount"`
	PastDueReminderCount      int64                `json:"pastDueReminderCount"`
	NextReminder              reminders.Reminder   `json:"nextReminder"`
	ServiceRecordCount        int64                `json:"serviceRecordCount"`
	ServiceRecordCost         float64              `json:"serviceRecordCost"`
	RepairRecordCount         int64                `json:"repairRecordCount"`
	RepairRecordCost          float64              `json:"repairRecordCost"`
	UpgradeRecordCount        int64                `json:"upgradeRecordCount"`
	UpgradeRecordCost         float64              `json:"upgradeRecordCost"`
	TaxRecordCount            int64                `json:"taxRecordCount"`
	TaxRecordCost             float64              `json:"taxRecordCost"`
	GasRecordCount            int64                `json:"gasRecordCount"`
	GasRecordCost             float64              `json:"gasRecordCost"`
	LastReportedOdometer      int64                `json:"lastReportedOdometer"`
	PlanRecordBackLogCount    int64                `json:"planRecordBackLogCount"`
	PlanRecordInProgressCount int64                `json:"planRecordInProgressCount"`
	PlanRecordTestingCount    int64                `json:"planRecordTestingCount"`
	PlanRecordDoneCount       int64                `json:"planRecordDoneCount"`
}

func convertSingle(in map[string]interface{}) VehicleInfo {
	return VehicleInfo{
		VehicleData:               vehicles.ConvertSingle(in["vehicleData"].(map[string]interface{})),
		VeryUrgentReminderCount:   parser.ParseInt(in["veryUrgentReminderCount"]),
		UrgentReminderCount:       parser.ParseInt(in["urgentReminderCount"]),
		NotUrgentReminderCount:    parser.ParseInt(in["notUrgentReminderCount"]),
		PastDueReminderCount:      parser.ParseInt(in["pastDueReminderCount"]),
		NextReminder:              reminders.ConvertSingle(in["nextReminder"]),
		ServiceRecordCount:        parser.ParseInt(in["serviceRecordCount"]),
		ServiceRecordCost:         parser.ParseFloat(in["serviceRecordCost"]),
		RepairRecordCount:         parser.ParseInt(in["repairRecordCount"]),
		RepairRecordCost:          parser.ParseFloat(in["repairRecordCost"]),
		UpgradeRecordCount:        parser.ParseInt(in["upgradeRecordCount"]),
		UpgradeRecordCost:         parser.ParseFloat(in["upgradeRecordCost"]),
		TaxRecordCount:            parser.ParseInt(in["taxRecordCount"]),
		TaxRecordCost:             parser.ParseFloat(in["taxRecordCost"]),
		GasRecordCount:            parser.ParseInt(in["gasRecordCount"]),
		GasRecordCost:             parser.ParseFloat(in["gasRecordCost"]),
		LastReportedOdometer:      parser.ParseInt(in["lastReportedOdometer"]),
		PlanRecordBackLogCount:    parser.ParseInt(in["planRecordBackLogCount"]),
		PlanRecordInProgressCount: parser.ParseInt(in["planRecordInProgressCount"]),
		PlanRecordTestingCount:    parser.ParseInt(in["planRecordTestingCount"]),
		PlanRecordDoneCount:       parser.ParseInt(in["planRecordDoneCount"]),
	}
}

func convertAll(inx []map[string]interface{}) ([]VehicleInfo, error) {
	out := make([]VehicleInfo, 0)

	for _, in := range inx {

		out = append(out, convertSingle(in))

	}

	return out, nil
}
