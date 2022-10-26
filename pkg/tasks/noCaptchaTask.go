package tasks

import (
	"fmt"
	"math"
	"net/url"
)

type NoCaptchaTaskProxyless struct {
	Type                string  `json:"type"`
	WebsiteURL          string  `json:"websiteURL"`
	WebsiteKey          string  `json:"websiteKey"`
	RecaptchaDataSValue *string `json:"recaptchaDataSValue,omitempty"`
	userAgentAndCookies
}

func NewNoCaptchaTaskProxyless(websiteURL, websiteKey, userAgent string) NoCaptchaTaskProxyless {
	return NoCaptchaTaskProxyless{
		Type:       "NoCaptchaTaskProxyless",
		WebsiteURL: websiteURL,
		WebsiteKey: websiteKey,
		userAgentAndCookies: userAgentAndCookies{
			UserAgent: &userAgent,
		},
	}
}

func (t NoCaptchaTaskProxyless) WithRecaptchaDataSValue(recaptchaDataSValue string) NoCaptchaTaskProxyless {
	t.RecaptchaDataSValue = &recaptchaDataSValue
	return t
}

func (t NoCaptchaTaskProxyless) Validate() error {
	if _, err := url.ParseRequestURI(t.WebsiteURL); err != nil {
		return fmt.Errorf("parse WebsiteURL: %w", err)
	}

	if len(t.WebsiteKey) < 1 || len(t.WebsiteKey) > math.MaxInt {
		return fmt.Errorf("WebsiteKey len error")
	}
	return nil
}

type NoCaptchaTask struct {
	NoCaptchaTaskProxyless
	taskProxy
}

func NewNoCaptchaTask(websiteURL, websiteKey, proxyType, proxyAddress string, proxyPort int) NoCaptchaTask {
	return NoCaptchaTask{
		NoCaptchaTaskProxyless: NoCaptchaTaskProxyless{
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

func (t NoCaptchaTask) Validate() error {
	if err := t.NoCaptchaTaskProxyless.Validate(); err != nil {
		return err
	}
	if err := t.taskProxy.validate(); err != nil {
		return err
	}
	return nil
}

type NoCaptchaTaskSolution struct {
	GRecaptchaResponse string            `json:"gRecaptchaResponse"`
	Cookies            map[string]string `json:"cookies"`
}
