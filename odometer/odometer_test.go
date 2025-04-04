package odometer_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/thaapaniemi/go-lubelogger-api/client"
	"github.com/thaapaniemi/go-lubelogger-api/odometer"
)

const ApiReplyEnUS = `[
  {
    "id": 1,
    "date": "2020-05-07",
    "initialOdometer": 204603,
    "odometer": 204836,
    "notes": "Auto Insert From Gas Record via CSV Import.",
    "tags": "",
    "extraFields": [],
    "files": []
  },
  {
    "id": 2,
    "date": "2020-05-29",
    "initialOdometer": 204836,
    "odometer": 205056,
    "notes": "Auto Insert From Gas Record via CSV Import.",
    "tags": "",
    "extraFields": [],
    "files": []
  },
  {
    "id": 3,
    "date": "2020-06-28",
    "initialOdometer": 205056,
    "odometer": 205202,
    "notes": "Auto Insert From Gas Record via CSV Import.",
    "tags": "",
    "extraFields": [],
    "files": []
  },
  {
    "id": 4,
    "date": "2020-07-10",
    "initialOdometer": 205202,
    "odometer": 205339,
    "notes": "Auto Insert From Gas Record via CSV Import.",
    "tags": "",
    "extraFields": [],
    "files": []
  },
  {
    "id": 5,
    "date": "2020-07-24",
    "initialOdometer": 205339,
    "odometer": 205489,
    "notes": "Auto Insert From Gas Record via CSV Import.",
    "tags": "",
    "extraFields": [],
    "files": []
  },
  {
    "id": 6,
    "date": "2020-08-09",
    "initialOdometer": 205489,
    "odometer": 205698,
    "notes": "Auto Insert From Gas Record via CSV Import.",
    "tags": "",
    "extraFields": [],
    "files": []
  }]`

func TestOdometerRecords(t *testing.T) {
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { fmt.Fprint(w, ApiReplyEnUS) }))
	defer svr.Close()

	c := client.New(svr.URL, "", "")

	_, err := odometer.GetRecords(context.Background(), c, 1)
	if err != nil {
		t.Errorf("GetOdometerRecords got exception: %v", err)
	}

}
