package gasrecords

import (
	"fmt"
	"time"

	"github.com/thaapaniemi/go-lubelogger-api/parser"
)

// Re-export UploadedFile for backwards compatibility
type UploadedFile = parser.UploadedFile

type GasRecord struct {
	ID           int64          `json:"id,omitempty"`
	Date         time.Time      `json:"date,omitempty"`
	Odometer     int64          `json:"odometer,omitempty"`
	FuelConsumed float64        `json:"fuelConsumed,omitempty"`
	FuelEconomy  float64        `json:"fuelEconomy,omitempty"`
	IsFillToFull bool           `json:"isFillToFull"`
	MissedFuelUp bool           `json:"missedFuelUp"`
	Description  string         `json:"description,omitempty"`
	Notes        string         `json:"notes,omitempty"`
	Cost         float64        `json:"cost,omitempty"`
	Tags         string         `json:"tags,omitempty"`
	ExtraFields  []string       `json:"extraFields,omitempty"`
	Files        []UploadedFile `json:"files,omitempty"`
}

func convertSingle(in map[string]interface{}) (GasRecord, error) {
	files, err := parser.ConvertUploadedFiles(in["files"])
	if err != nil {
		return GasRecord{}, fmt.Errorf("failed to convert files: %v", err)
	}

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
		Files:        files,
	}, nil
}

func convertAll(inx []map[string]interface{}) ([]GasRecord, error) {
	out := make([]GasRecord, 0)

	for i, in := range inx {
		record, err := convertSingle(in)
		if err != nil {
			return nil, fmt.Errorf("failed to convert record %d: %v", i, err)
		}

		out = append(out, record)
	}
	return out, nil
}
