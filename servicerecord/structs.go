package servicerecord

import (
	"time"

	"github.com/thaapaniemi/go-lubelogger-api/parser"
)

type ServiceRecord struct {
	ID          int64     `json:"id,omitempty"`
	Date        time.Time `json:"date"`
	Odometer    int64     `json:"odometer"`
	Description string    `json:"description"`
	Notes       string    `json:"notes"`
	Cost        float64   `json:"cost"`
	Tags        string    `json:"tags"`
	ExtraFields []string  `json:"extraFields"`
	Files       []string  `json:"files"`
}

func convertSingle(in map[string]interface{}) ServiceRecord {

	return ServiceRecord{
		ID:          parser.ParseInt(in["id"]),
		Date:        parser.ParseDateISO8601(in["date"]),
		Odometer:    parser.ParseInt(in["odometer"]),
		Description: parser.ParseString(in["description"]),
		Cost:        parser.ParseFloat(in["cost"]),
		Notes:       parser.ParseString(in["notes"]),
		Tags:        parser.ParseString(in["tags"]),
		ExtraFields: parser.ParseStringSlice(in["extraFields"]),
		Files:       parser.ParseStringSlice(in["files"]),
	}
}

func convertAll(inx []map[string]interface{}) ([]ServiceRecord, error) {
	out := make([]ServiceRecord, 0)

	for _, in := range inx {

		out = append(out, convertSingle(in))

	}
	return out, nil
}
