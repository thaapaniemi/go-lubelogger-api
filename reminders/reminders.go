package reminders

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/thaapaniemi/go-lubelogger-api/client"
)

type Urgency string

const URGENCY_NOT_URGENT Urgency = "NotUrgent"
const URGENCY_URGENT Urgency = "Urgent"
const URGENCY_VERY_URGENT Urgency = "VeryUrgent"
const URGENCY_PAST_DUE Urgency = "PastDue"

func GetRecords(ctx context.Context, c client.Client, vehicleId int64) ([]Reminder, error) {
	q := client.Query{
		Path:    "/api/vehicle/reminders",
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

func SendReminderEmails(ctx context.Context, c client.Client, urgencies []Urgency) error {
	q := client.Query{
		Path:    "/api/vehicle/reminders/send",
		Method:  http.MethodGet,
		Payload: nil,
		Query:   make(url.Values, 0),
	}

	for _, urgency := range urgencies {
		q.Query.Add("urgencies", string(urgency))
	}

	_, err := c.DoRequest(ctx, q)
	return err
}
