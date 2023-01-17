package tasks

import (
	"math"
	"net/url"
)

type TurnstileTaskProxyless struct {
	Type       string `json:"type"`
	WebsiteURL string `json:"websiteURL"`
	WebsiteKey string `json:"websiteKey"`
}

func NewTurnstileTaskProxyless(websiteURL, websiteKey string) TurnstileTaskProxyless {
	return TurnstileTaskProxyless{
		Type:       "TurnstileTaskProxyless",
		WebsiteURL: websiteURL,
		WebsiteKey: websiteKey,
	}
}

func (t TurnstileTaskProxyless) Validate() error {
	if _, err := url.ParseRequestURI(t.WebsiteURL); err != nil {
		return ErrInvalidWebsiteUrl
	}

	if len(t.WebsiteKey) < 1 || len(t.WebsiteKey) > math.MaxInt {
		return ErrInvalidWebSiteKey
	}
	return nil
}

type TurnstileTask struct {
	TurnstileTaskProxyless
	taskProxy
}

func NewTurnstileTask(websiteURL, websiteKey, proxyType, proxyAddress string, proxyPort int) TurnstileTask {
	return TurnstileTask{
		TurnstileTaskProxyless: TurnstileTaskProxyless{
			Type:       "TurnstileTask",
			WebsiteURL: websiteURL,
			WebsiteKey: websiteKey,
		},
		taskProxy: taskProxy{
			ProxyType:    proxyType,
			ProxyAddress: proxyAddress,
			ProxyPort:    proxyPort,
		},
	}
}

func (t TurnstileTask) Validate() error {
	if err := t.TurnstileTaskProxyless.Validate(); err != nil {
		return err
	}
	if err := t.taskProxy.validate(); err != nil {
		return err
	}
	return nil
}

type turnstileTaskSolution struct {
	token string `json:"token"`
}
