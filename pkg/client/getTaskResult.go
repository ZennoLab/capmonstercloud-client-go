package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/tandem97/capmonstercloud-client-go/pkg/tasks"
)

func (c *capmonsterClient) getImageToTextTaskResult(taskId int) (tasks.ImageToTextTaskResult, error) {
	var result tasks.ImageToTextTaskResult
	return result, c.getTaskResult(taskId, &result)
}

func (c *capmonsterClient) getNoCaptchaTaskResult(taskId int) (tasks.NoCaptchaTaskResult, error) {
	var result tasks.NoCaptchaTaskResult
	return result, c.getTaskResult(taskId, &result)
}

func (c *capmonsterClient) getRecaptchaV3TaskResult(taskId int) (tasks.RecaptchaV3TaskTaskResult, error) {
	var result tasks.RecaptchaV3TaskTaskResult
	return result, c.getTaskResult(taskId, &result)
}

func (c *capmonsterClient) getRecaptchaV2EnterpriseTaskResult(taskId int) (tasks.RecaptchaV2EnterpriseTaskResult, error) {
	var result tasks.RecaptchaV2EnterpriseTaskResult
	return result, c.getTaskResult(taskId, &result)
}

func (c *capmonsterClient) getFunCaptchaTaskResult(taskId int) (tasks.FunCaptchaTaskResult, error) {
	var result tasks.FunCaptchaTaskResult
	return result, c.getTaskResult(taskId, &result)
}

func (c *capmonsterClient) getHCaptchaTaskResult(taskId int) (tasks.HCaptchaTaskResult, error) {
	var result tasks.HCaptchaTaskResult
	return result, c.getTaskResult(taskId, &result)
}

func (c *capmonsterClient) getGeeTestTaskResult(taskId int) (tasks.GeeTestTaskResult, error) {
	var result tasks.GeeTestTaskResult
	return result, c.getTaskResult(taskId, &result)
}

type getTaskResultPayload struct {
	ClientKey string `json:"clientKey"`
	TaskId    int    `json:"taskId"`
}

func (c *capmonsterClient) getTaskResult(taskId int, result interface{}) error {
	body, err := json.Marshal(getTaskResultPayload{
		ClientKey: c.clientKey,
		TaskId:    taskId,
	})
	if err != nil {
		return fmt.Errorf("marshal payload for request: %w", err)
	}

	bodyReader := bytes.NewReader(body)
	req, err := http.NewRequest("POST", getTaskResultUrl, bodyReader)
	if err != nil {
		return fmt.Errorf("create http request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("http request: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read response body: %w", err)
	}

	if err := json.Unmarshal(respBody, result); err != nil {
		return fmt.Errorf("unmarshal responce payload: %w", err)
	}

	return nil
}
