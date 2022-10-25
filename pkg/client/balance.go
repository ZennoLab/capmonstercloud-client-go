package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type getBalanceRequestPayload struct {
	ClientKey string `json:"clientKey"`
}

type getBalanceResponsePayload struct {
	ErrorId   int     `json:"errorId"`
	ErrorCode string  `json:"errorCode"`
	Balance   float64 `json:"balance"`
}

func (c *capmonsterClient) GetBalance() (float64, error) {
	body, err := json.Marshal(getBalanceRequestPayload{ClientKey: c.clientKey})
	if err != nil {
		return 0, fmt.Errorf("marshal payload for request: %w", err)
	}

	bodyReader := bytes.NewReader(body)
	req, err := http.NewRequest("GET", "https://api.capmonster.cloud/getBalance", bodyReader)
	if err != nil {
		return 0, fmt.Errorf("create http request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return 0, fmt.Errorf("http request: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, fmt.Errorf("read response body: %w", err)
	}

	var respPayload getBalanceResponsePayload
	if err := json.Unmarshal(respBody, &respPayload); err != nil {
		return 0, fmt.Errorf("unmarshal responce payload: %w", err)
	}

	if respPayload.ErrorId != 0 {
		return 0, fmt.Errorf("%v", respPayload.ErrorCode)
	}

	return respPayload.Balance, nil
}
