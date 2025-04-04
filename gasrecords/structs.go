package gasrecords

import (
	"time"

	"github.com/thaapaniemi/go-lubelogger-api/parser"
)

type GasRecord struct {
	ID           int64     `json:"id,omitempty"`
	Date         time.Time `json:"date"`
	Odometer     int64     `json:"odometer"`
	FuelConsumed float64   `json:"fuelConsumed"`
	FuelEconomy  float64   `json:"fuelEconomy"`
	IsFillToFull bool      `json:"isFillToFull"`
	MissedFuelUp bool      `json:"missedFuelUp"`
	Description  string    `json:"description"`
	Notes        string    `json:"notes"`
	Cost         float64   `json:"cost"`
	Tags         string    `json:"tags"`
	ExtraFields  []string  `json:"extraFields"`
	Files        []string  `json:"files"`
}

func convertSingle(in map[string]interface{}) GasRecord {

	return GasRecord{
		ID:           parser.ParseInt(in["id"]),
		Date:         parser.ParseDateISO8601(in["date"]),
		Odometer:     parser.ParseInt(in["odometer"]),
		FuelConsumed: parser.ParseFloat(in["fuelConsumed"]),
		FuelEconomy:  parser.ParseFloat(in["fuelEconomy"]),
		IsFillToFull: parser.ParseBool(in["isFillToFull"]),
		MissedFuelUp: parser.ParseBool(in["missedFuelUp"]),
		Description:  parser.ParseString(in["description"]),
		Notes:        parser.ParseString(in["notes"]),
		Cost:         parser.ParseFloat(in["cost"]),
		Tags:         parser.ParseString(in["tags"]),
		ExtraFields:  parser.ParseStringSlice(in["extraFields"]),
		Files:        parser.ParseStringSlice(in["files"]),
	}
}

func convertAll(inx []map[string]interface{}) ([]GasRecord, error) {
	out := make([]GasRecord, 0)

	for _, in := range inx {

		out = append(out, convertSingle(in))

	}
	return out, nil
}
