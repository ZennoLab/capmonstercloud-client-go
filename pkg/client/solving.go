package client

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/gachi-lord/capmonstercloud-client-go/pkg/tasks"
)

type resulter interface {
	getErrorId() int
	getErrorCode() string
	getStatus() string
}

type validater interface {
	Validate() error
}

func (c *capmonsterClient) solve(task validater, callbackUrl *string, timings resolveCapTimings, noCache bool, taskResult resulter) error {
	if err := task.Validate(); err != nil {
		return fmt.Errorf("validate: %w", err)
	}

	taskId, err := c.createTask(task, callbackUrl)
	if err != nil {
		return fmt.Errorf("create task: %w", err)
	}

	timeoutTicker := time.NewTicker(timings.timeout)
	var retryTicker *time.Ticker
	if noCache {
		retryTicker = time.NewTicker(timings.firstRequestNoCacheDelay)
	} else {
		retryTicker = time.NewTicker(timings.firstRequestDelay)
	}
	defer retryTicker.Stop()
	defer timeoutTicker.Stop()

	var setTickerForRetries sync.Once

	for {
		select {
		case <-retryTicker.C:
			setTickerForRetries.Do(func() { retryTicker.Reset(timings.requestsInterval) })
			err := c.getTaskResult(taskId, taskResult)
			switch {
			case err != nil:
				if errors.Is(err, errServiceUnavailable) {
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

func (c *capmonsterClient) SolveImageToText(task tasks.ImageToTextTask, callbackUrl *string) (*tasks.ImageToTextTaskSolution, error) {
	var result imageToTextTaskResult
	if err := c.solve(task, callbackUrl, imageToTextTimings, false, &result); err != nil {
		return nil, fmt.Errorf("resolve: %w", err)
	}
	return &result.Solution, nil
}

func (c *capmonsterClient) SolveRecaptchaV2(task tasks.RecaptchaV2Task, noCache bool, callbackUrl *string) (*tasks.RecaptchaV2TaskSolution, error) {
	var result recaptchaV2Result
	if err := c.solve(task, callbackUrl, recaptchaV2TaskTimings, noCache, &result); err != nil {
		return nil, fmt.Errorf("resolve: %w", err)
	}
	return &result.Solution, nil
}

func (c *capmonsterClient) SolveRecaptchaV2Proxyless(task tasks.RecaptchaV2TaskProxyless, noCache bool, callbackUrl *string) (*tasks.RecaptchaV2TaskSolution, error) {
	var result recaptchaV2Result
	if err := c.solve(task, callbackUrl, recaptchaV2TaskTimings, noCache, &result); err != nil {
		return nil, fmt.Errorf("resolve: %w", err)
	}
	return &result.Solution, nil
}

func (c *capmonsterClient) SolveRecaptchaV3Proxyless(task tasks.RecaptchaV3TaskProxyless, noCache bool, callbackUrl *string) (*tasks.RecaptchaV3TaskTaskSolution, error) {
	var result recaptchaV3TaskTaskResult
	if err := c.solve(task, callbackUrl, recaptchaV3Timings, noCache, &result); err != nil {
		return nil, fmt.Errorf("resolve: %w", err)
	}
	return &result.Solution, nil
}

func (c *capmonsterClient) SolveRecaptchaV2Enterprise(task tasks.RecaptchaV2EnterpriseTask, noCache bool, callbackUrl *string) (*tasks.RecaptchaV2EnterpriseTaskSolution, error) {
	var result recaptchaV2EnterpriseTaskResult
	if err := c.solve(task, callbackUrl, recaptchaV2EnterpriseTimings, noCache, &result); err != nil {
		return nil, fmt.Errorf("resolve: %w", err)
	}
	return &result.Solution, nil
}

func (c *capmonsterClient) SolveRecaptchaV2EnterpriseProxyless(task tasks.RecaptchaV2EnterpriseTaskProxyless, noCache bool, callbackUrl *string) (*tasks.RecaptchaV2EnterpriseTaskSolution, error) {
	var result recaptchaV2EnterpriseTaskResult
	if err := c.solve(task, callbackUrl, recaptchaV2EnterpriseTimings, noCache, &result); err != nil {
		return nil, fmt.Errorf("resolve: %w", err)
	}
	return &result.Solution, nil
}

func (c *capmonsterClient) SolveFunCaptcha(task tasks.FunCaptchaTask, noCache bool, callbackUrl *string) (*tasks.FunCaptchaTaskSolution, error) {
	var result funCaptchaTaskResult
	if err := c.solve(task, callbackUrl, funCaptchaTimings, noCache, &result); err != nil {
		return nil, fmt.Errorf("resolve: %w", err)
	}
	return &result.Solution, nil
}

func (c *capmonsterClient) SolveFunCaptchaProxyless(task tasks.FunCaptchaTaskProxyless, noCache bool, callbackUrl *string) (*tasks.FunCaptchaTaskSolution, error) {
	var result funCaptchaTaskResult
	if err := c.solve(task, callbackUrl, funCaptchaTimings, noCache, &result); err != nil {
		return nil, fmt.Errorf("resolve: %w", err)
	}
	return &result.Solution, nil
}

func (c *capmonsterClient) SolveHCaptcha(task tasks.HCaptchaTask, noCache bool, callbackUrl *string) (*tasks.HCaptchaTaskSolution, error) {
	var result hCaptchaTaskResult
	if err := c.solve(task, callbackUrl, hCaptchaTimings, noCache, &result); err != nil {
		return nil, fmt.Errorf("resolve: %w", err)
	}
	return &result.Solution, nil
}

func (c *capmonsterClient) SolveHCaptchaProxyless(task tasks.HCaptchaTaskProxyless, noCache bool, callbackUrl *string) (*tasks.HCaptchaTaskSolution, error) {
	var result hCaptchaTaskResult
	if err := c.solve(task, callbackUrl, hCaptchaTimings, noCache, &result); err != nil {
		return nil, fmt.Errorf("resolve: %w", err)
	}
	return &result.Solution, nil
}

func (c *capmonsterClient) SolveGeeTest(task tasks.GeeTestTask, callbackUrl *string) (*tasks.GeeTestTaskSolution, error) {
	var result geeTestTaskResult
	if err := c.solve(task, callbackUrl, geeTestTimings, false, &result); err != nil {
		return nil, fmt.Errorf("resolve: %w", err)
	}
	return &result.Solution, nil
}

func (c *capmonsterClient) SolveGeeTestProxyless(task tasks.GeeTestTaskProxyless, callbackUrl *string) (*tasks.GeeTestTaskSolution, error) {
	var result geeTestTaskResult
	if err := c.solve(task, callbackUrl, geeTestTimings, false, &result); err != nil {
		return nil, fmt.Errorf("resolve: %w", err)
	}
	return &result.Solution, nil
}

func (c *capmonsterClient) SolveTurnstile(task tasks.TurnstileTask, callbackUrl *string) (*tasks.turnstileSolution, error) {
	var result turnstileTaskResult
	if err := c.solve(task, callbackUrl, turnstileTimings, false, &result); err != nil {
		return nil, fmt.Errorf("resolve: %w", err)
	}
	return &result.Solution, nil
}

func (c *capmonsterClient) SolveTurnstileProxyless(task tasks.TurnstileTaskProxyless, callbackUrl *string) (*tasks.turnstileSolution, error) {
	var result turnstileTaskResult
	if err := c.solve(task, callbackUrl, turnstileTimings, false, &result); err != nil {
		return nil, fmt.Errorf("resolve: %w", err)
	}
	return &result.Solution, nil
}
