package vehicles_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/thaapaniemi/go-lubelogger-api/client"
	"github.com/thaapaniemi/go-lubelogger-api/vehicles"
)

const ApiReply = `[{"id":1,"imageLocation":"/defaults/noimage.png","year":1992,"make":"Jeep","model":"Cherokee","licensePlate":"MOMWAGON","purchaseDate":"2025-01-15","soldDate":"2025-03-31","purchasePrice":0,"soldPrice":0,"isElectric":false,"isDiesel":false,"useHours":false,"odometerOptional":false,"extraFields":[],"tags":[],"hasOdometerAdjustment":false,"odometerMultiplier":1,"odometerDifference":0,"dashboardMetrics":[],"vehicleIdentifier":"LicensePlate"}]`

func TestGetVehicles(t *testing.T) {
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { fmt.Fprint(w, ApiReply) }))
	defer svr.Close()

	c := client.New(svr.URL, "", "")

	purchaseDate, _ := time.Parse("2006-01-02", "2025-01-15")
	soldDate, _ := time.Parse("2006-01-02", "2025-03-31")

	type args struct {
		ctx    context.Context
		client client.Client
	}
	tests := []struct {
		name    string
		args    args
		want    []vehicles.VehicleData
		wantErr bool
	}{
		{name: "test1", args: args{ctx: context.Background(), client: c}, want: []vehicles.VehicleData{vehicles.VehicleData{ID: 1, Model: "Cherokee", PurchaseDate: purchaseDate, SoldDate: soldDate}}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := vehicles.GetRecords(tt.args.ctx, tt.args.client)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetVehicles() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if len(got) != len(tt.want) {
				t.Errorf("found devices amount don't match")
			}

			for i, _ := range got {
				if got[i].ID != tt.want[i].ID {
					t.Errorf("GetVehicles() ID got = %v, want = %v", got[i].ID, tt.want[i].ID)
				}

				if got[i].Model != tt.want[i].Model {
					t.Errorf("GetVehicles() Model got = %v, want = %v", got[i].Model, tt.want[i].Model)
				}

				if got[i].PurchaseDate.Format("2006-01-02") != tt.want[i].PurchaseDate.Format("2006-01-02") {
					t.Errorf("GetVehicles() PurchaseDate got = %v, want = %v", got[i].PurchaseDate, tt.want[i].PurchaseDate)
				}

				if got[i].SoldDate.Format("2006-01-02") != tt.want[i].SoldDate.Format("2006-01-02") {
					t.Errorf("GetVehicles() SoldDate got = %v, want = %v", got[i].SoldDate, tt.want[i].SoldDate)
				}
			}

		})
	}
}
