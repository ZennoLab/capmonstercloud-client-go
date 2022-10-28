package client

import (
	"net/http"
	"time"
)

const (
	getBalanceUrl    = "https://api.capmonster.cloud/getBalance"
	createTaskUrl    = "https://api.capmonster.cloud/createTask"
	getTaskResultUrl = "https://api.capmonster.cloud/getTaskResult/"
	softId           = 58
)

var (
	reqHeaders = map[string][]string{
		"UserAgent": {"Zennolab.CapMonsterCloud.Client.Go/0.0.1"},
	}

	httpClient = &http.Client{
		Transport: &http.Transport{
			MaxIdleConns:    10,
			IdleConnTimeout: 30 * time.Second,
		},
		Timeout: 21 * time.Second,
	}
)

type capmonsterClient struct {
	httpClient *http.Client
	clientKey  string
}

func New(clientKey string) *capmonsterClient {
	return &capmonsterClient{
		httpClient: httpClient,
		clientKey:  clientKey,
	}
}
