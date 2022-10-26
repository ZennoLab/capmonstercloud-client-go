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

func (c *capmonsterClient) ResolveImageToText(task tasks.ImageToTextTask, callbackUrl *string) (*tasks.ImageToTextTaskSolution, error) {
	taskId, err := c.createTask(task, callbackUrl)
	if err != nil {
		return nil, fmt.Errorf("create task: %w", err)
	}

	time.Sleep(imageToTextTimings.firstRequestDelay)
	retryTicker := time.NewTicker(imageToTextTimings.requestsInterval)
	timeoutTicker := time.NewTicker(imageToTextTimings.timeout)

	var result imageToTextTaskResult
	if err := c.resolve(retryTicker.C, timeoutTicker.C, taskId, &result); err != nil {
		return nil, fmt.Errorf("resolve: %w", err)
	}
	return &result.Solution, nil
}

func (c *capmonsterClient) resolveNoCaptcha(task interface{}, callbackUrl *string) (*tasks.NoCaptchaTaskSolution, error) {
	taskId, err := c.createTask(task, callbackUrl)
	if err != nil {
		return nil, fmt.Errorf("create task: %w", err)
	}

	time.Sleep(noCaptchaTaskTimings.firstRequestDelay)
	retryTicker := time.NewTicker(noCaptchaTaskTimings.requestsInterval)
	timeoutTicker := time.NewTicker(noCaptchaTaskTimings.timeout)

	var result noCaptchaTaskResult
	if err := c.resolve(retryTicker.C, timeoutTicker.C, taskId, &result); err != nil {
		return nil, fmt.Errorf("resolve: %w", err)
	}
	return &result.Solution, nil
}

func (c *capmonsterClient) ResolveNoCaptcha(task tasks.NoCaptchaTask, callbackUrl *string) (*tasks.NoCaptchaTaskSolution, error) {
	return c.resolveNoCaptcha(task, callbackUrl)
}

func (c *capmonsterClient) ResolveNoCaptchaProxyless(task tasks.NoCaptchaTaskProxyless, callbackUrl *string) (*tasks.NoCaptchaTaskSolution, error) {
	return c.resolveNoCaptcha(task, callbackUrl)
}

func (c *capmonsterClient) ResolveRecaptchaV3Proxyless(task tasks.RecaptchaV3TaskProxyless, callbackUrl *string) (*tasks.RecaptchaV3TaskTaskSolution, error) {
	taskId, err := c.createTask(task, callbackUrl)
	if err != nil {
		return nil, fmt.Errorf("create task: %w", err)
	}

	time.Sleep(recaptchaV3Timings.firstRequestDelay)
	retryTicker := time.NewTicker(recaptchaV3Timings.requestsInterval)
	timeoutTicker := time.NewTicker(recaptchaV3Timings.timeout)

	var result recaptchaV3TaskTaskResult
	if err := c.resolve(retryTicker.C, timeoutTicker.C, taskId, &result); err != nil {
		return nil, fmt.Errorf("resolve: %w", err)
	}
	return &result.Solution, nil
}

func (c *capmonsterClient) resolveRecaptchaV2Enterprise(task interface{}, callbackUrl *string) (*tasks.RecaptchaV2EnterpriseTaskSolution, error) {
	taskId, err := c.createTask(task, callbackUrl)
	if err != nil {
		return nil, fmt.Errorf("create task: %w", err)
	}

	time.Sleep(recaptchaV2EnterpriseTimings.firstRequestDelay)
	retryTicker := time.NewTicker(recaptchaV2EnterpriseTimings.requestsInterval)
	timeoutTicker := time.NewTicker(recaptchaV2EnterpriseTimings.timeout)

	var result recaptchaV2EnterpriseTaskResult
	if err := c.resolve(retryTicker.C, timeoutTicker.C, taskId, &result); err != nil {
		return nil, fmt.Errorf("resolve: %w", err)
	}
	return &result.Solution, nil
}

func (c *capmonsterClient) ResolveRecaptchaV2Enterprise(task tasks.RecaptchaV2EnterpriseTask, callbackUrl *string) (*tasks.RecaptchaV2EnterpriseTaskSolution, error) {
	return c.resolveRecaptchaV2Enterprise(task, callbackUrl)
}

func (c *capmonsterClient) ResolveRecaptchaV2EnterpriseProxyless(task tasks.RecaptchaV2EnterpriseTaskProxyless, callbackUrl *string) (*tasks.RecaptchaV2EnterpriseTaskSolution, error) {
	return c.resolveRecaptchaV2Enterprise(task, callbackUrl)
}

func (c *capmonsterClient) resolveFunCaptcha(task interface{}, callbackUrl *string) (*tasks.FunCaptchaTaskSolution, error) {
	taskId, err := c.createTask(task, callbackUrl)
	if err != nil {
		return nil, fmt.Errorf("create task: %w", err)
	}

	time.Sleep(funCaptchaTimings.firstRequestDelay)
	retryTicker := time.NewTicker(funCaptchaTimings.requestsInterval)
	timeoutTicker := time.NewTicker(funCaptchaTimings.timeout)

	var result funCaptchaTaskResult
	if err := c.resolve(retryTicker.C, timeoutTicker.C, taskId, &result); err != nil {
		return nil, fmt.Errorf("resolve: %w", err)
	}
	return &result.Solution, nil
}

func (c *capmonsterClient) ResolveFunCaptcha(task tasks.RecaptchaV2EnterpriseTask, callbackUrl *string) (*tasks.FunCaptchaTaskSolution, error) {
	return c.resolveFunCaptcha(task, callbackUrl)
}

func (c *capmonsterClient) ResolveFunCaptchaProxyless(task tasks.RecaptchaV2EnterpriseTaskProxyless, callbackUrl *string) (*tasks.FunCaptchaTaskSolution, error) {
	return c.resolveFunCaptcha(task, callbackUrl)
}

func (c *capmonsterClient) resolveHCaptcha(task interface{}, callbackUrl *string) (*tasks.HCaptchaTaskSolution, error) {
	taskId, err := c.createTask(task, callbackUrl)
	if err != nil {
		return nil, fmt.Errorf("create task: %w", err)
	}

	time.Sleep(hCaptchaTimings.firstRequestDelay)
	retryTicker := time.NewTicker(hCaptchaTimings.requestsInterval)
	timeoutTicker := time.NewTicker(hCaptchaTimings.timeout)

	var result hCaptchaTaskResult
	if err := c.resolve(retryTicker.C, timeoutTicker.C, taskId, &result); err != nil {
		return nil, fmt.Errorf("resolve: %w", err)
	}
	return &result.Solution, nil
}

func (c *capmonsterClient) ResolveHCaptcha(task tasks.RecaptchaV2EnterpriseTask, callbackUrl *string) (*tasks.HCaptchaTaskSolution, error) {
	return c.resolveHCaptcha(task, callbackUrl)
}

func (c *capmonsterClient) ResolveHCaptchaProxyless(task tasks.RecaptchaV2EnterpriseTaskProxyless, callbackUrl *string) (*tasks.HCaptchaTaskSolution, error) {
	return c.resolveHCaptcha(task, callbackUrl)
}

func (c *capmonsterClient) resolveGeeTest(task interface{}, callbackUrl *string) (*tasks.GeeTestTaskSolution, error) {
	taskId, err := c.createTask(task, callbackUrl)
	if err != nil {
		return nil, fmt.Errorf("create task: %w", err)
	}

	time.Sleep(geeTestTimings.firstRequestDelay)
	retryTicker := time.NewTicker(geeTestTimings.requestsInterval)
	timeoutTicker := time.NewTicker(geeTestTimings.timeout)

	var result geeTestTaskResult
	if err := c.resolve(retryTicker.C, timeoutTicker.C, taskId, &result); err != nil {
		return nil, fmt.Errorf("resolve: %w", err)
	}
	return &result.Solution, nil
}

func (c *capmonsterClient) ResolveGeeTest(task tasks.RecaptchaV2EnterpriseTask, callbackUrl *string) (*tasks.GeeTestTaskSolution, error) {
	return c.resolveGeeTest(task, callbackUrl)
}

func (c *capmonsterClient) ResolveGeeTestProxyless(task tasks.RecaptchaV2EnterpriseTaskProxyless, callbackUrl *string) (*tasks.GeeTestTaskSolution, error) {
	return c.resolveGeeTest(task, callbackUrl)
}

func (c *capmonsterClient) resolve(retryTickerCh <-chan time.Time, timeoutTickerCh <-chan time.Time, taskId int, result interface{}) error {
	for {
		select {
		case <-retryTickerCh:
			var result geeTestTaskResult
			err := c.getTaskResult(taskId, result)
			switch {
			case err != nil:
				return fmt.Errorf("get task result: %w", err)
			case result.ErrorId != 0 && result.ErrorCode != "CAPTCHA_NOT_READY":
				if err, ok := errMap[result.ErrorCode]; ok {
					return fmt.Errorf("get task result: %w", err)
				}
				return fmt.Errorf("get task result: %w", errUnknown)
			case result.Status == "ready":
				return nil

			}
		case <-timeoutTickerCh:
			return errTimeout
		}
	}
}
