package client

import (
	"fmt"
	"net/http"
	"time"

	"github.com/tandem97/capmonstercloud-client-go/pkg/tasks"
)

const (
	getBalanceUrl    = "https://api.capmonster.cloud/getBalance"
	createTaskUrl    = "https://api.capmonster.cloud/createTask"
	getTaskResultUrl = "https://api.capmonster.cloud/getTaskResult/"
	softId           = 58
)

type capmonsterClient struct {
	httpClient http.Client
	clientKey  string
}

func New(clientKey string) *capmonsterClient {
	return &capmonsterClient{
		httpClient: http.Client{
			Transport: &http.Transport{
				MaxIdleConns:    10,
				IdleConnTimeout: 30 * time.Second,
			},
			Timeout: 21 * time.Second,
		},
		clientKey: clientKey,
	}
}

func (c *capmonsterClient) resolve(task interface{}, callbackUrl *string, timings resolveCapTimings, taskResult resulter) error {
	taskId, err := c.createTask(task, callbackUrl)
	if err != nil {
		return fmt.Errorf("create task: %w", err)
	}

	time.Sleep(timings.firstRequestDelay)
	retryTicker := time.NewTicker(timings.requestsInterval)
	timeoutTicker := time.NewTicker(timings.timeout)

	for {
		select {
		case <-retryTicker.C:
			err := c.getTaskResult(taskId, taskResult)
			switch {
			case err != nil:
				if err == errServiceServiceUnavailable {
					continue
				}
				return fmt.Errorf("get task result: %w", err)
			case taskResult.getErrorId() != 0 && taskResult.getErrorCode() != "CAPTCHA_NOT_READY":
				if err, ok := errMap[taskResult.getErrorCode()]; ok {
					return fmt.Errorf("get task result: %w", err)
				}
				return errUnknown
			case taskResult.getStatus() == "ready":
				return nil

			}
		case <-timeoutTicker.C:
			return errTimeout
		}
	}
}

func (c *capmonsterClient) ResolveImageToText(task tasks.ImageToTextTask, callbackUrl *string) (*tasks.ImageToTextTaskSolution, error) {
	if err := task.Validate(); err != nil {
		return nil, fmt.Errorf("validate: %w", err)
	}

	var result imageToTextTaskResult
	if err := c.resolve(task, callbackUrl, imageToTextTimings, &result); err != nil {
		return nil, fmt.Errorf("resolve: %w", err)
	}

	return &result.Solution, nil
}

func (c *capmonsterClient) ResolveNoCaptcha(task tasks.NoCaptchaTask, callbackUrl *string) (*tasks.NoCaptchaTaskSolution, error) {
	if err := task.Validate(); err != nil {
		return nil, fmt.Errorf("validate: %w", err)
	}

	var result noCaptchaTaskResult
	if err := c.resolve(task, callbackUrl, noCaptchaTaskTimings, &result); err != nil {
		return nil, fmt.Errorf("resolve: %w", err)
	}

	return &result.Solution, nil
}

func (c *capmonsterClient) ResolveNoCaptchaProxyless(task tasks.NoCaptchaTaskProxyless, callbackUrl *string) (*tasks.NoCaptchaTaskSolution, error) {
	if err := task.Validate(); err != nil {
		return nil, fmt.Errorf("validate: %w", err)
	}

	var result noCaptchaTaskResult
	if err := c.resolve(task, callbackUrl, noCaptchaTaskTimings, &result); err != nil {
		return nil, fmt.Errorf("resolve: %w", err)
	}

	return &result.Solution, nil
}

func (c *capmonsterClient) ResolveRecaptchaV3Proxyless(task tasks.RecaptchaV3TaskProxyless, callbackUrl *string) (*tasks.RecaptchaV3TaskTaskSolution, error) {
	if err := task.Validate(); err != nil {
		return nil, fmt.Errorf("validate: %w", err)
	}

	var result recaptchaV3TaskTaskResult
	if err := c.resolve(task, callbackUrl, recaptchaV3Timings, &result); err != nil {
		return nil, fmt.Errorf("resolve: %w", err)
	}

	return &result.Solution, nil
}

func (c *capmonsterClient) ResolveRecaptchaV2Enterprise(task tasks.RecaptchaV2EnterpriseTask, callbackUrl *string) (*tasks.RecaptchaV2EnterpriseTaskSolution, error) {
	if err := task.Validate(); err != nil {
		return nil, fmt.Errorf("validate: %w", err)
	}

	var result recaptchaV2EnterpriseTaskResult
	if err := c.resolve(task, callbackUrl, recaptchaV2EnterpriseTimings, &result); err != nil {
		return nil, fmt.Errorf("resolve: %w", err)
	}

	return &result.Solution, nil
}

func (c *capmonsterClient) ResolveRecaptchaV2EnterpriseProxyless(task tasks.RecaptchaV2EnterpriseTaskProxyless, callbackUrl *string) (*tasks.RecaptchaV2EnterpriseTaskSolution, error) {
	if err := task.Validate(); err != nil {
		return nil, fmt.Errorf("validate: %w", err)
	}

	var result recaptchaV2EnterpriseTaskResult
	if err := c.resolve(task, callbackUrl, recaptchaV2EnterpriseTimings, &result); err != nil {
		return nil, fmt.Errorf("resolve: %w", err)
	}

	return &result.Solution, nil
}

func (c *capmonsterClient) ResolveFunCaptcha(task tasks.FunCaptchaTask, callbackUrl *string) (*tasks.FunCaptchaTaskSolution, error) {
	if err := task.Validate(); err != nil {
		return nil, fmt.Errorf("validate: %w", err)
	}

	var result funCaptchaTaskResult
	if err := c.resolve(task, callbackUrl, funCaptchaTimings, &result); err != nil {
		return nil, fmt.Errorf("resolve: %w", err)
	}

	return &result.Solution, nil
}

func (c *capmonsterClient) ResolveFunCaptchaProxyless(task tasks.FunCaptchaTaskProxyless, callbackUrl *string) (*tasks.FunCaptchaTaskSolution, error) {
	if err := task.Validate(); err != nil {
		return nil, fmt.Errorf("validate: %w", err)
	}

	var result funCaptchaTaskResult
	if err := c.resolve(task, callbackUrl, funCaptchaTimings, &result); err != nil {
		return nil, fmt.Errorf("resolve: %w", err)
	}

	return &result.Solution, nil
}

func (c *capmonsterClient) ResolveHCaptcha(task tasks.HCaptchaTask, callbackUrl *string) (*tasks.HCaptchaTaskSolution, error) {
	if err := task.Validate(); err != nil {
		return nil, fmt.Errorf("validate: %w", err)
	}

	var result hCaptchaTaskResult
	if err := c.resolve(task, callbackUrl, hCaptchaTimings, &result); err != nil {
		return nil, fmt.Errorf("resolve: %w", err)
	}

	return &result.Solution, nil
}

func (c *capmonsterClient) ResolveHCaptchaProxyless(task tasks.HCaptchaTaskProxyless, callbackUrl *string) (*tasks.HCaptchaTaskSolution, error) {
	if err := task.Validate(); err != nil {
		return nil, fmt.Errorf("validate: %w", err)
	}

	var result hCaptchaTaskResult
	if err := c.resolve(task, callbackUrl, hCaptchaTimings, &result); err != nil {
		return nil, fmt.Errorf("resolve: %w", err)
	}

	return &result.Solution, nil
}

func (c *capmonsterClient) ResolveGeeTest(task tasks.GeeTestTask, callbackUrl *string) (*tasks.GeeTestTaskSolution, error) {
	if err := task.Validate(); err != nil {
		return nil, fmt.Errorf("validate: %w", err)
	}

	var result geeTestTaskResult
	if err := c.resolve(task, callbackUrl, geeTestTimings, &result); err != nil {
		return nil, fmt.Errorf("resolve: %w", err)
	}

	return &result.Solution, nil
}

func (c *capmonsterClient) ResolveGeeTestProxyless(task tasks.GeeTestTaskProxyless, callbackUrl *string) (*tasks.GeeTestTaskSolution, error) {
	if err := task.Validate(); err != nil {
		return nil, fmt.Errorf("validate: %w", err)
	}

	var result geeTestTaskResult
	if err := c.resolve(task, callbackUrl, geeTestTimings, &result); err != nil {
		return nil, fmt.Errorf("resolve: %w", err)
	}

	return &result.Solution, nil
}
