package calendar_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/thaapaniemi/go-lubelogger-api/calendar"
	"github.com/thaapaniemi/go-lubelogger-api/client"
)

const ApiReply = `200: BEGIN:VCALENDAR
VERSION:2.0
PRODID:lubelogger.com
CALSCALE:GREGORIAN
METHOD:PUBLISH
END:VCALENDAR

BEGIN:VCALENDAR
VERSION:2.0
PRODID:lubelogger.com
CALSCALE:GREGORIAN
METHOD:PUBLISH
END:VCALENDAR
`

func TestCalendar(t *testing.T) {
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { fmt.Fprint(w, ApiReply) }))
	defer svr.Close()

	c := client.New(svr.URL, "", "")

	res, err := calendar.GetCalendar(context.Background(), c)
	if err != nil {
		t.Errorf("got exception: %v", err)
	}

	if res != ApiReply {
		t.Errorf("GetCalendar should not touch the result")
	}
}
