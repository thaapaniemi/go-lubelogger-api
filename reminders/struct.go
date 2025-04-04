package reminders

import (
	"time"

	"github.com/thaapaniemi/go-lubelogger-api/parser"
)

type Reminder struct {
	Description string    `json:"description"`
	Urgency     string    `json:"urgency"`
	Metric      string    `json:"metric"`
	Notes       string    `json:"notes"`
	DueDate     time.Time `json:"dueDate"`
	DueOdometer int64     `json:"dueOdometer"`
}

func ConvertSingle(x interface{}) Reminder {

	if x == nil {
		return Reminder{}
	}

	in := x.(map[string]interface{})

	return Reminder{
		Description: parser.ParseString(in["description"]),
		Urgency:     parser.ParseString(in["urgency"]),
		Metric:      parser.ParseString(in["metric"]),
		Notes:       parser.ParseString(in["notes"]),
		DueDate:     parser.ParseDateISO8601(in["dueDate"]),
		DueOdometer: parser.ParseInt(in["dueOdometer"]),
	}
}

func convertAll(inx []map[string]interface{}) ([]Reminder, error) {
	out := make([]Reminder, 0)

	for _, in := range inx {

		out = append(out, ConvertSingle(in))

	}
	return out, nil
}
