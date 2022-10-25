package client

import (
	"net/http"
	"time"
)

type capmonsterClient struct {
	httpClient http.Client
	clientKey  string
}

func New(clientKey string) *capmonsterClient {
	return &capmonsterClient{
		httpClient: http.Client{
			Transport: &http.Transport{
				MaxIdleConns:    10,
				IdleConnTimeout: 30 * time.Second,
			},
			Timeout: 21 * time.Second,
		},
		clientKey: clientKey,
	}
}
