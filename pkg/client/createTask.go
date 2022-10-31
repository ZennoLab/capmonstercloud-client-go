package client

import (
	"encoding/json"
	"fmt"
)

type createTaskRequestPayload struct {
	SoftId      int         `json:"softId"`
	ClientKey   string      `json:"clientKey"`
	Task        interface{} `json:"task"`
	CallbackUrl *string     `json:"callbackUrl,omitempty"`
}

type createTaskRequestResponsePayload struct {
	ErrorId   int    `json:"errorId"`
	ErrorCode string `json:"errorCode"`
	TaskId    int    `json:"taskId"`
}

func (c *capmonsterClient) createTask(task interface{}, callbackUrl *string) (int, error) {
	body, err := json.Marshal(createTaskRequestPayload{
		SoftId:      softId,
		ClientKey:   c.clientKey,
		Task:        task,
		CallbackUrl: callbackUrl,
	})
	if err != nil {
		return 0, fmt.Errorf("marshal payload for request: %w", err)
	}

	respBody, err := c.invokeRequest(body, createTaskUrl)
	if err != nil {
		return 0, fmt.Errorf("invoke request: %w", err)
	}

	var respPayload createTaskRequestResponsePayload
	if err := json.Unmarshal(respBody, &respPayload); err != nil {
		return 0, fmt.Errorf("unmarshal response payload: %w", err)
	}

	if respPayload.ErrorId != 0 {
		if err, ok := errMap[respPayload.ErrorCode]; ok {
			return 0, err
		}
		return 0, errUnknown
	}

	return respPayload.TaskId, nil
}
