package tasks

import (
	"fmt"
	"math"
	"net/url"
)

type RecaptchaV2TaskProxyless struct {
	Type                string  `json:"type"`
	WebsiteURL          string  `json:"websiteURL"`
	WebsiteKey          string  `json:"websiteKey"`
	RecaptchaDataSValue *string `json:"recaptchaDataSValue,omitempty"`
	userAgentAndCookies
}

func NewRecaptchaV2TaskProxyless(websiteURL, websiteKey string) RecaptchaV2TaskProxyless {
	return RecaptchaV2TaskProxyless{
		Type:       "NoCaptchaTaskProxyless",
		WebsiteURL: websiteURL,
		WebsiteKey: websiteKey,
	}
}

func (t RecaptchaV2TaskProxyless) WithRecaptchaDataSValue(recaptchaDataSValue string) RecaptchaV2TaskProxyless {
	t.RecaptchaDataSValue = &recaptchaDataSValue
	return t
}

func (t RecaptchaV2TaskProxyless) Validate() error {
	if _, err := url.ParseRequestURI(t.WebsiteURL); err != nil {
		return fmt.Errorf("parse WebsiteURL: %w", err)
	}

	if len(t.WebsiteKey) < 1 || len(t.WebsiteKey) > math.MaxInt {
		return fmt.Errorf("WebsiteKey len error")
	}
	return nil
}

type RecaptchaV2Task struct {
	RecaptchaV2TaskProxyless
	taskProxy
}

func NewRecaptchaV2Task(websiteURL, websiteKey, proxyType, proxyAddress string, proxyPort int) RecaptchaV2Task {
	return RecaptchaV2Task{
		RecaptchaV2TaskProxyless: RecaptchaV2TaskProxyless{
			Type:       "NoCaptchaTask",
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

func (t RecaptchaV2Task) Validate() error {
	if err := t.RecaptchaV2TaskProxyless.Validate(); err != nil {
		return err
	}
	if err := t.taskProxy.validate(); err != nil {
		return err
	}
	return nil
}

type RecaptchaV2TaskSolution struct {
	GRecaptchaResponse string            `json:"gRecaptchaResponse"`
	Cookies            map[string]string `json:"cookies"`
}
