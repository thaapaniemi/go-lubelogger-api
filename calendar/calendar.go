package calendar

import (
	"context"
	"net/http"
	"net/url"

	"github.com/thaapaniemi/go-lubelogger-api/client"
)

func GetCalendar(ctx context.Context, c client.Client) (string, error) {
	q := client.Query{
		Path:    "/api/calendar",
		Method:  http.MethodGet,
		Payload: nil,
		Query:   make(url.Values, 0),
	}

	raw, err := c.DoRequest(ctx, q)
	return string(raw), err
}
