package vehicles

import (
	"context"
	"net/http"
	"net/url"

	"github.com/thaapaniemi/go-lubelogger-api/client"
)

func GetRecords(ctx context.Context, c client.Client) ([]VehicleData, error) {
	q := client.Query{
		Path:    "/api/vehicles",
		Method:  http.MethodGet,
		Payload: nil,
		Query:   make(url.Values, 0),
	}

	raw, err := c.DoRequest(ctx, q)

	if err != nil {
		return nil, err
	}

	var response []map[string]interface{}
	c.Decode(raw, &response)

	out, err := convertAll(response)
	return out, err
}
