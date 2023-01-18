package client

import (
	"encoding/json"
	"fmt"

	"github.com/gachi-lord/capmonstercloud-client-go/pkg/tasks"
)

type result struct {
	ErrorId   int    `json:"errorId"`
	ErrorCode string `json:"errorCode"`
	Status    string `json:"status"`
}

func (r result) getErrorId() int {
	return r.ErrorId
}

func (r result) getErrorCode() string {
	return r.ErrorCode
}

func (r result) getStatus() string {
	return r.Status
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

type recaptchaV2Result struct {
	result
	Solution tasks.RecaptchaV2TaskSolution `json:"solution"`
}

type recaptchaV3TaskTaskResult struct {
	result
	Solution tasks.RecaptchaV3TaskTaskSolution `json:"solution"`
}

type recaptchaV2EnterpriseTaskResult struct {
	result
	Solution tasks.RecaptchaV2EnterpriseTaskSolution `json:"solution"`
}

type turnstileTaskResult struct {
	result
	Solution tasks.TurnstileTaskSolution `json:"solution"`
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

	respBody, err := c.invokeRequest(body, getTaskResultUrl)
	if err != nil {
		return fmt.Errorf("invoke request: %w", err)
	}

	if err := json.Unmarshal(respBody, &result); err != nil {
		return fmt.Errorf("unmarshal response payload: %w", err)
	}

	return nil
}
