package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/tandem97/capmonstercloud-client-go/pkg/tasks"
)

type resulter interface {
	getErrorId() int
	getErrorCode() string
	getStatus() string
}

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
	req.Header.Add("UserAgent", "Zennolab.CapMonsterCloud.Client/0.0.1")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("http request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusServiceUnavailable {
		return errServiceServiceUnavailable
	}
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return fmt.Errorf("response status code: %v", resp.StatusCode)
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read response body: %w", err)
	}

	if err := json.Unmarshal(respBody, &result); err != nil {
		return fmt.Errorf("unmarshal response payload: %w", err)
	}

	return nil
}
