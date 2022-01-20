package client

import (
	"net/http"
	"time"
)

type APIClient struct {
	client  http.Client
	timeout time.Duration
}

func NewClient(duration int) *APIClient {
	return &APIClient{
		client:  *http.DefaultClient,
		timeout: time.Duration(duration) * time.Second,
	}
}
