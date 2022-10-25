package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/tandem97/capmonstercloud-client-go/pkg/tasks"
)

func (c *capmonsterClient) SendImageToTextTask(task tasks.ImageToTextTask, callbackUrl *string) (int, error) {
	return c.sendTask(task, callbackUrl)
}

func (c *capmonsterClient) SendNoCaptchaTaskProxyless(task tasks.NoCaptchaTaskProxyless, callbackUrl *string) (int, error) {
	return c.sendTask(task, callbackUrl)
}

func (c *capmonsterClient) SendNoCaptchaTask(task tasks.NoCaptchaTask, callbackUrl *string) (int, error) {
	return c.sendTask(task, callbackUrl)
}

func (c *capmonsterClient) SendRecaptchaV3TaskProxyless(task tasks.RecaptchaV3TaskProxyless, callbackUrl *string) (int, error) {
	return c.sendTask(task, callbackUrl)
}

func (c *capmonsterClient) SendRecaptchaV2EnterpriseTask(task tasks.RecaptchaV2EnterpriseTask, callbackUrl *string) (int, error) {
	return c.sendTask(task, callbackUrl)
}

func (c *capmonsterClient) SendRecaptchaV2EnterpriseTaskProxyless(task tasks.RecaptchaV2EnterpriseTaskProxyless, callbackUrl *string) (int, error) {
	return c.sendTask(task, callbackUrl)
}

func (c *capmonsterClient) SendFunCaptchaTask(task tasks.FunCaptchaTask, callbackUrl *string) (int, error) {
	return c.sendTask(task, callbackUrl)
}

func (c *capmonsterClient) SendFunCaptchaTaskProxyless(task tasks.FunCaptchaTaskProxyless, callbackUrl *string) (int, error) {
	return c.sendTask(task, callbackUrl)
}

func (c *capmonsterClient) SendHCaptchaTask(task tasks.HCaptchaTask, callbackUrl *string) (int, error) {
	return c.sendTask(task, callbackUrl)
}

func (c *capmonsterClient) SendHCaptchaTaskProxyless(task tasks.FunCaptchaTaskProxyless, callbackUrl *string) (int, error) {
	return c.sendTask(task, callbackUrl)
}

func (c *capmonsterClient) SendGeeTestTask(task tasks.GeeTestTask, callbackUrl *string) (int, error) {
	return c.sendTask(task, callbackUrl)
}

func (c *capmonsterClient) SendGeeTestTaskProxyless(task tasks.GeeTestTaskProxyless, callbackUrl *string) (int, error) {
	return c.sendTask(task, callbackUrl)
}

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

func (c *capmonsterClient) sendTask(task interface{}, callbackUrl *string) (int, error) {
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
