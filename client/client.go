package client

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/thaapaniemi/go-lubelogger-api/debuglog"
)

type Client struct {
	Endpoint   string
	basicAuth  string
	httpClient *http.Client
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func New(endpoint, username, password string) Client {
	endpoint = strings.TrimSuffix(endpoint, "/")

	return Client{
		Endpoint:   endpoint,
		basicAuth:  basicAuth(username, password),
		httpClient: http.DefaultClient,
	}
}

func (c *Client) HttpClient(newClient *http.Client) {
	c.httpClient = newClient
}

func (c *Client) DoRequest(ctx context.Context, r LubeLoggerRequest) ([]byte, error) {
	url := c.Endpoint + r.LLPath()

	buf := bytes.NewBuffer(r.LLPayload())

	req, err := http.NewRequest(r.LLMethod(), url, buf)
	if err != nil {
		return nil, err
	}

	req.URL.RawQuery = r.LLQuery()

	req.Header.Set("culture-invariant", "true")
	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", c.basicAuth))

	ctype := r.LLContentType()
	if ctype != "" {
		req.Header.Set("Content-Type", ctype)
	}

	debuglog.Debugf("DoRequest payload: %s", string(r.LLPayload()))
	debuglog.Debugf("DoRequest Content-Type: %s", ctype)
	debuglog.Debugf("DoRequest url: %s", url)

	response, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	resBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	debuglog.Debugf("DoRequest response status: %d", response.StatusCode)
	debuglog.Debugf("DoRequest response: %s", string(resBody))

	if response.StatusCode != http.StatusOK {
		return resBody, fmt.Errorf("http request non-ok response status: %d, body: %s", response.StatusCode, resBody)
	}

	return resBody, nil
}

func (c *Client) Decode(in []byte, out interface{}) error {
	dec := json.NewDecoder(bytes.NewBuffer(in))
	dec.UseNumber()

	err := dec.Decode(&out)

	return err
}
