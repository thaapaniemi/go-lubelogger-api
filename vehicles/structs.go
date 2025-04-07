package vehicles

import (
	"time"

	"github.com/thaapaniemi/go-lubelogger-api/parser"
)

type VehicleData struct {
	ID                    int64     `json:"id,omitempty"`
	ImageLocation         string    `json:"imageLocation,omitempty"`
	Year                  int64     `json:"year,omitempty"`
	Make                  string    `json:"make,omitempty"`
	Model                 string    `json:"model,omitempty"`
	LicensePlate          string    `json:"licensePlate,omitempty"`
	PurchaseDate          time.Time `json:"purchaseDate,omitempty"`
	SoldDate              time.Time `json:"soldDate,omitempty"`
	PurchasePrice         float64   `json:"purchasePrice,omitempty"`
	SoldPrice             float64   `json:"soldPrice,omitempty"`
	IsElectric            bool      `json:"isElectric,omitempty"`
	IsDiesel              bool      `json:"isDiesel,omitempty"`
	UseHours              bool      `json:"useHours,omitempty"`
	OdometerOptional      bool      `json:"odometerOptional,omitempty"`
	ExtraFields           []string  `json:"extraFields,omitempty"`
	Tags                  []string  `json:"tags,omitempty"`
	HasOdometerAdjustment bool      `json:"hasOdometerAdjustment,omitempty"`
	OdometerMultiplier    int64     `json:"odometerMultiplier,omitempty"`
	OdometerDifference    int64     `json:"odometerDifference,omitempty"`
	DashboardMetrics      []int64   `json:"dashboardMetrics,omitempty"`
	VehicleIdentifier     string    `json:"vehicleIdentifier,omitempty"`
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
