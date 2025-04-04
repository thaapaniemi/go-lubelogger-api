package document

import (
	"bytes"
	"context"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"

	"github.com/thaapaniemi/go-lubelogger-api/client"
)

func (o Document) Upload(ctx context.Context, c client.Client) (string, error) {
	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	formFile, err := writer.CreateFormFile("documents", o.Key)
	if err != nil {
		return "", err
	}

	// Write the decoded file content to the form field
	_, err = io.Copy(formFile, bytes.NewReader(o.Src))
	if err != nil {
		return "", err
	}

	// Close the writer to finalize the multipart form data
	err = writer.Close()
	if err != nil {
		return "", err
	}

	query := client.Query{
		Path:        "/api/documents/upload",
		Method:      http.MethodPost,
		Payload:     buf.Bytes(),
		Query:       make(url.Values, 0),
		ContentType: writer.FormDataContentType(),
	}

	response, err := c.DoRequest(ctx, query)
	if err != nil {
		return "", err
	}

	var u []uploadResponse
	err = c.Decode(response, &u)

	var location string
	if len(u) > 0 {
		location = u[0].Location
	}

	return location, err
}
