package client

import (
	"encoding/json"
	"fmt"
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

	respBody, err := c.invokeRequest(body, getBalanceUrl)
	if err != nil {
		return 0, fmt.Errorf("invoke request: %w", err)
	}

	var respPayload getBalanceResponsePayload
	if err := json.Unmarshal(respBody, &respPayload); err != nil {
		return 0, fmt.Errorf("unmarshal response payload: %w", err)
	}

	if respPayload.ErrorId != 0 {
		if err, ok := errMap[respPayload.ErrorCode]; ok {
			return 0, err
		}
		return 0, errUnknown
	}

	return respPayload.Balance, nil
}
