package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/tandem97/capmonstercloud-client-go/pkg/tasks"
)

type result struct {
	ErrorId   int    `json:"errorId"`
	ErrorCode string `json:"errorCode"`
	Status    string `json:"status"`
}

type funCaptchaTaskResult struct {
	result
	Solution tasks.FunCaptchaTaskSolution `json:"solution"`
}

type geeTestTaskResult struct {
	result
	Solution tasks.GeeTestTaskSolution `json:"solution"`
}

type hCaptchaTaskResult struct {
	result
	Solution tasks.HCaptchaTaskSolution `json:"solution"`
}

type imageToTextTaskResult struct {
	result
	Solution tasks.ImageToTextTaskSolution `json:"solution"`
}

type noCaptchaTaskResult struct {
	result
	Solution tasks.NoCaptchaTaskSolution `json:"solution"`
}

type recaptchaV3TaskTaskResult struct {
	result
	Solution tasks.RecaptchaV3TaskTaskSolution `json:"solution"`
}

type recaptchaV2EnterpriseTaskResult struct {
	result
	Solution tasks.RecaptchaV2EnterpriseTaskSolution `json:"solution"`
}

func (c *capmonsterClient) getImageToTextTaskResult(taskId int) (imageToTextTaskResult, error) {
	var result imageToTextTaskResult
	return result, c.getTaskResult(taskId, &result)
}

func (c *capmonsterClient) getNoCaptchaTaskResult(taskId int) (noCaptchaTaskResult, error) {
	var result noCaptchaTaskResult
	return result, c.getTaskResult(taskId, &result)
}

func (c *capmonsterClient) getRecaptchaV3TaskResult(taskId int) (recaptchaV3TaskTaskResult, error) {
	var result recaptchaV3TaskTaskResult
	return result, c.getTaskResult(taskId, &result)
}

func (c *capmonsterClient) getRecaptchaV2EnterpriseTaskResult(taskId int) (recaptchaV2EnterpriseTaskResult, error) {
	var result recaptchaV2EnterpriseTaskResult
	return result, c.getTaskResult(taskId, &result)
}

func (c *capmonsterClient) getFunCaptchaTaskResult(taskId int) (funCaptchaTaskResult, error) {
	var result funCaptchaTaskResult
	return result, c.getTaskResult(taskId, &result)
}

func (c *capmonsterClient) getHCaptchaTaskResult(taskId int) (hCaptchaTaskResult, error) {
	var result hCaptchaTaskResult
	return result, c.getTaskResult(taskId, &result)
}

func (c *capmonsterClient) getGeeTestTaskResult(taskId int) (geeTestTaskResult, error) {
	var result geeTestTaskResult
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
