package client

import (
	"net/url"
)

type LubeLoggerRequest interface {
	LLPath() string
	LLMethod() string
	LLPayload() []byte
	LLQuery() string
	LLContentType() string
}

type Query struct {
	Path        string
	Method      string
	Payload     []byte
	Query       url.Values
	ContentType string
}

func (q Query) LLPath() string {
	return q.Path
}

func (Q Query) LLMethod() string {
	return Q.Method
}

func (q Query) LLPayload() []byte {
	return q.Payload
}

func (q Query) LLQuery() string {
	return q.Query.Encode()
}

func (q Query) LLContentType() string {
	return q.ContentType
}
