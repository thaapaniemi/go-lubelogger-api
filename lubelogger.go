package lubelogger

import "github.com/thaapaniemi/go-lubelogger-api/client"

const TEST = ""

func NewClient(endpoint, username, password string) client.Client {
	return client.New(endpoint, username, password)
}
