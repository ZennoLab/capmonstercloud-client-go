package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type createTaskRequestPayload struct {
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
		ClientKey:   c.clientKey,
		Task:        task,
		CallbackUrl: callbackUrl,
	})
	if err != nil {
		return 0, fmt.Errorf("marshal payload for request: %w", err)
	}

	bodyReader := bytes.NewReader(body)
	req, err := http.NewRequest("POST", createTaskUrl, bodyReader)
	if err != nil {
		return 0, fmt.Errorf("create http request: %w", err)
	}
	req.Header.Add("UserAgent", "Zennolab.CapMonsterCloud.Client/0.0.1")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return 0, fmt.Errorf("http request: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, fmt.Errorf("read response body: %w", err)
	}

	var respPayload createTaskRequestResponsePayload
	if err := json.Unmarshal(respBody, &respPayload); err != nil {
		return 0, fmt.Errorf("unmarshal responce payload: %w", err)
	}

	if respPayload.ErrorId != 0 {
		if err, ok := errMap[respPayload.ErrorCode]; ok {
			return 0, err
		}
		return 0, errUnknown
	}

	return respPayload.TaskId, nil
}
