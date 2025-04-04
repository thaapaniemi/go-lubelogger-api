package vehicles

import (
	"time"

	"github.com/thaapaniemi/go-lubelogger-api/parser"
)

type VehicleData struct {
	ID                    int64     `json:"id"`
	ImageLocation         string    `json:"imageLocation"`
	Year                  int64     `json:"year"`
	Make                  string    `json:"make"`
	Model                 string    `json:"model"`
	LicensePlate          string    `json:"licensePlate"`
	PurchaseDate          time.Time `json:"purchaseDate,omitempty"`
	SoldDate              time.Time `json:"soldDate,omitempty"`
	PurchasePrice         float64   `json:"purchasePrice"`
	SoldPrice             float64   `json:"soldPrice"`
	IsElectric            bool      `json:"isElectric"`
	IsDiesel              bool      `json:"isDiesel"`
	UseHours              bool      `json:"useHours"`
	OdometerOptional      bool      `json:"odometerOptional"`
	ExtraFields           []string  `json:"extraFields"`
	Tags                  []string  `json:"tags"`
	HasOdometerAdjustment bool      `json:"hasOdometerAdjustment"`
	OdometerMultiplier    int64     `json:"odometerMultiplier"`
	OdometerDifference    int64     `json:"odometerDifference"`
	DashboardMetrics      []int64   `json:"dashboardMetrics"`
	VehicleIdentifier     string    `json:"vehicleIdentifier"`
}

func convertAll(inx []map[string]interface{}) ([]VehicleData, error) {
	out := make([]VehicleData, 0)

	for _, in := range inx {
		out = append(out, ConvertSingle(in))
	}

	return out, nil
}

func ConvertSingle(in map[string]interface{}) VehicleData {
	return VehicleData{
		ID:                    parser.ParseInt(in["id"]),
		ImageLocation:         parser.ParseString(in["imageLocation"]),
		Year:                  parser.ParseInt(in["year"]),
		Make:                  parser.ParseString(in["make"]),
		Model:                 parser.ParseString(in["model"]),
		LicensePlate:          parser.ParseString(in["licensePlate"]),
		PurchaseDate:          parser.ParseDateISO8601(in["purchaseDate"]),
		SoldDate:              parser.ParseDateISO8601(in["soldDate"]),
		PurchasePrice:         parser.ParseFloat(in["purchasePrice"]),
		SoldPrice:             parser.ParseFloat(in["soldPrice"]),
		IsElectric:            parser.ParseBool(in["isElectric"]),
		IsDiesel:              parser.ParseBool(in["isDiesel"]),
		UseHours:              parser.ParseBool(in["useHours"]),
		OdometerOptional:      parser.ParseBool(in["odometerOptional"]),
		ExtraFields:           parser.ParseStringSlice(in["extraFields"]),
		Tags:                  parser.ParseStringSlice(in["tags"]),
		HasOdometerAdjustment: parser.ParseBool(in["hasOdometerAdjustment"]),
		OdometerMultiplier:    parser.ParseInt(in["odometerMultiplier"]),
		OdometerDifference:    parser.ParseInt(in["odometerDifference"]),
		DashboardMetrics:      parser.ParseIntSlice(in["dashboardMetrics"]),
		VehicleIdentifier:     parser.ParseString(in["vehicleIdentifier"]),
	}
}
