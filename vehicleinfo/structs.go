package vehicleinfo

import (
	"github.com/thaapaniemi/go-lubelogger-api/parser"
	"github.com/thaapaniemi/go-lubelogger-api/reminders"
	"github.com/thaapaniemi/go-lubelogger-api/vehicles"
)

type VehicleInfo struct {
	VehicleData               vehicles.VehicleData `json:"vehicleData,omitempty"`
	VeryUrgentReminderCount   int64                `json:"veryUrgentReminderCount,omitempty"`
	UrgentReminderCount       int64                `json:"urgentReminderCount,omitempty"`
	NotUrgentReminderCount    int64                `json:"notUrgentReminderCount,omitempty"`
	PastDueReminderCount      int64                `json:"pastDueReminderCount,omitempty"`
	NextReminder              reminders.Reminder   `json:"nextReminder,omitempty"`
	ServiceRecordCount        int64                `json:"serviceRecordCount,omitempty"`
	ServiceRecordCost         float64              `json:"serviceRecordCost,omitempty"`
	RepairRecordCount         int64                `json:"repairRecordCount,omitempty"`
	RepairRecordCost          float64              `json:"repairRecordCost,omitempty"`
	UpgradeRecordCount        int64                `json:"upgradeRecordCount,omitempty"`
	UpgradeRecordCost         float64              `json:"upgradeRecordCost,omitempty"`
	TaxRecordCount            int64                `json:"taxRecordCount,omitempty"`
	TaxRecordCost             float64              `json:"taxRecordCost,omitempty"`
	GasRecordCount            int64                `json:"gasRecordCount,omitempty"`
	GasRecordCost             float64              `json:"gasRecordCost,omitempty"`
	LastReportedOdometer      int64                `json:"lastReportedOdometer,omitempty"`
	PlanRecordBackLogCount    int64                `json:"planRecordBackLogCount,omitempty"`
	PlanRecordInProgressCount int64                `json:"planRecordInProgressCount,omitempty"`
	PlanRecordTestingCount    int64                `json:"planRecordTestingCount,omitempty"`
	PlanRecordDoneCount       int64                `json:"planRecordDoneCount,omitempty"`
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
