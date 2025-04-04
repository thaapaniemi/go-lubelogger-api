package odometer

import (
	"time"

	"github.com/thaapaniemi/go-lubelogger-api/parser"
)

type OdometerRecord struct {
	ID              int64     `json:"id,omitempty"`
	Date            time.Time `json:"date,omitempty"`
	InitialOdometer int64     `json:"initialOdometer,omitempty"`
	Odometer        int64     `json:"odometer,omitempty"`
	Notes           string    `json:"notes,omitempty"`
	Tags            string    `json:"tags,omitempty"`
	ExtraFields     []string  `json:"extraFields,omitempty"`
	Files           []string  `json:"files,omitempty"`
}

func convertSingle(in map[string]interface{}) OdometerRecord {
	return OdometerRecord{
		ID:              parser.ParseInt(in["id"]),
		Date:            parser.ParseDateISO8601(in["date"]),
		InitialOdometer: parser.ParseInt(in["initialOdometer"]),
		Odometer:        parser.ParseInt(in["odometer"]),
		Notes:           parser.ParseString(in["notes"]),
		Tags:            parser.ParseString(in["tags"]),
		ExtraFields:     parser.ParseStringSlice(in["extraFields"]),
		Files:           parser.ParseStringSlice(in["files"]),
	}
}

func convertAll(inx []map[string]interface{}) ([]OdometerRecord, error) {
	out := make([]OdometerRecord, 0)

	for _, in := range inx {

		out = append(out, convertSingle(in))

	}

	return out, nil
}
