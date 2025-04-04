package repairrecords

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/thaapaniemi/go-lubelogger-api/client"
)

func GetRecords(ctx context.Context, c client.Client, vehicleId int64) ([]RepairRecord, error) {
	q := client.Query{
		Path:    "/api/vehicle/repairrecords",
		Method:  http.MethodGet,
		Payload: nil,
		Query:   make(url.Values, 0),
	}
	q.Query.Set("vehicleId", fmt.Sprintf("%d", vehicleId))

	raw, err := c.DoRequest(ctx, q)

	if err != nil {
		return nil, err
	}

	var response []map[string]interface{}

	err = c.Decode(raw, &response)
	if err != nil {
		return nil, err
	}

	out, err := convertAll(response)
	return out, err
}

func (o RepairRecord) Add(ctx context.Context, c client.Client, vehicleId int64) error {
	o.ID = 0

	payload, err := json.Marshal(&o)
	if err != nil {
		return err
	}

	q := client.Query{
		Path:        "/api/vehicle/repairrecords/add",
		Method:      http.MethodPost,
		Payload:     payload,
		Query:       make(url.Values, 0),
		ContentType: "application/json",
	}
	q.Query.Set("vehicleId", fmt.Sprintf("%d", vehicleId))

	_, err = c.DoRequest(ctx, q)
	return err
}

func (o RepairRecord) Update(ctx context.Context, c client.Client) error {
	payload, err := json.Marshal(&o)
	if err != nil {
		return err
	}

	q := client.Query{
		Path:        "/api/vehicle/repairrecords/update",
		Method:      http.MethodPut,
		Payload:     payload,
		Query:       make(url.Values, 0),
		ContentType: "application/json",
	}

	_, err = c.DoRequest(ctx, q)
	return err
}

func (o RepairRecord) Delete(ctx context.Context, c client.Client) error {
	q := client.Query{
		Path:    "/api/vehicle/repairrecords/delete",
		Method:  http.MethodDelete,
		Payload: nil,
		Query:   make(url.Values, 0),
	}
	q.Query.Set("id", fmt.Sprintf("%d", o.ID))

	_, err := c.DoRequest(ctx, q)
	return err
}
