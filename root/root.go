package root

import (
	"context"
	"net/http"
	"net/url"

	"github.com/thaapaniemi/go-lubelogger-api/client"
)

func MakeBackup(ctx context.Context, c client.Client) ([]byte, error) {
	q := make(url.Values, 0)

	query := client.Query{
		Path:    "/api/makebackup",
		Method:  http.MethodGet,
		Payload: nil,
		Query:   q,
	}

	raw, err := c.DoRequest(ctx, query)
	return raw, err
}

func Cleanup(ctx context.Context, c client.Client) ([]byte, error) {
	q := make(url.Values, 0)
	q.Set("deepClean", "true")

	query := client.Query{
		Path:    "/api/cleanup",
		Method:  http.MethodGet,
		Payload: nil,
		Query:   q,
	}

	raw, err := c.DoRequest(ctx, query)
	return raw, err
}
