package upgraderecords

import (
	"time"

	"github.com/thaapaniemi/go-lubelogger-api/parser"
)

type UpgradeRecord struct {
	ID          int64     `json:"id,omitempty"`
	Date        time.Time `json:"date,omitempty"`
	Odometer    int64     `json:"odometer,omitempty"`
	Description string    `json:"description,omitempty"`
	Notes       string    `json:"notes,omitempty"`
	Cost        float64   `json:"cost,omitempty"`
	Tags        string    `json:"tags,omitempty"`
	ExtraFields []string  `json:"extraFields,omitempty"`
	Files       []string  `json:"files,omitempty"`
}

func convertSingle(in map[string]interface{}) UpgradeRecord {

	return UpgradeRecord{
		ID:          parser.ParseInt(in["id"]),
		Date:        parser.ParseDateISO8601(in["date"]),
		Odometer:    parser.ParseInt(in["odometer"]),
		Description: parser.ParseString(in["description"]),
		Notes:       parser.ParseString(in["notes"]),
		Cost:        parser.ParseFloat(in["cost"]),
		Tags:        parser.ParseString(in["tags"]),
		ExtraFields: parser.ParseStringSlice(in["extraFields"]),
		Files:       parser.ParseStringSlice(in["files"]),
	}
}

func convertAll(inx []map[string]interface{}) ([]UpgradeRecord, error) {
	out := make([]UpgradeRecord, 0)

	for _, in := range inx {

		out = append(out, convertSingle(in))

	}
	return out, nil
}
