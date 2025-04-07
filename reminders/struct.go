package reminders

import (
	"time"

	"github.com/thaapaniemi/go-lubelogger-api/parser"
)

type Reminder struct {
	Description string    `json:"description,omitempty"`
	Urgency     string    `json:"urgency,omitempty"`
	Metric      string    `json:"metric,omitempty"`
	Notes       string    `json:"notes,omitempty"`
	DueDate     time.Time `json:"dueDate,omitempty"`
	DueOdometer int64     `json:"dueOdometer,omitempty"`
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
