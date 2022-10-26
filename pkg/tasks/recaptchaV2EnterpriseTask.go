package tasks

import (
	"fmt"
	"math"
	"net/url"
)

type RecaptchaV2EnterpriseTaskProxyless struct {
	Type              string  `json:"type"`
	WebsiteURL        string  `json:"websiteURL"`
	WebsiteKey        string  `json:"websiteKey"`
	EnterprisePayload *string `json:"enterprisePayload,omitempty"`
	ApiDomain         *string `json:"apiDomain,omitempty"`
}

func NewRecaptchaV2EnterpriseTaskProxyless(websiteURL, websiteKey string) RecaptchaV2EnterpriseTaskProxyless {
	return RecaptchaV2EnterpriseTaskProxyless{
		Type:       "RecaptchaV2EnterpriseTaskProxyless",
		WebsiteURL: websiteURL,
		WebsiteKey: websiteKey,
	}
}

func (t RecaptchaV2EnterpriseTaskProxyless) WithEnterprisePayload(enterprisePayload string) RecaptchaV2EnterpriseTaskProxyless {
	t.EnterprisePayload = &enterprisePayload
	return t
}

func (t RecaptchaV2EnterpriseTaskProxyless) WithApiDomain(apiDomain string) RecaptchaV2EnterpriseTaskProxyless {
	t.ApiDomain = &apiDomain
	return t
}

func (t RecaptchaV2EnterpriseTaskProxyless) Validate() error {
	if _, err := url.ParseRequestURI(t.WebsiteURL); err != nil {
		return fmt.Errorf("parse WebsiteURL: %w", err)
	}
	if len(t.WebsiteKey) < 1 || len(t.WebsiteKey) > math.MaxInt {
		return fmt.Errorf("WebsiteKey len error")
	}
	if t.EnterprisePayload != nil && (len(*t.EnterprisePayload) < 1 || len(*t.EnterprisePayload) > math.MaxInt) {
		return fmt.Errorf("EnterprisePayload len error")
	}
	return nil
}

type RecaptchaV2EnterpriseTask struct {
	RecaptchaV2EnterpriseTaskProxyless
	taskProxy
	userAgentAndCookies
}

func NewRecaptchaV2EnterpriseTask(websiteURL, websiteKey, proxyType, proxyAddress string, proxyPort int) RecaptchaV2EnterpriseTask {
	return RecaptchaV2EnterpriseTask{
		RecaptchaV2EnterpriseTaskProxyless: RecaptchaV2EnterpriseTaskProxyless{
			Type:       "RecaptchaV2EnterpriseTask",
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

func (t RecaptchaV2EnterpriseTask) Validate() error {
	if err := t.RecaptchaV2EnterpriseTaskProxyless.Validate(); err != nil {
		return err
	}
	if err := t.taskProxy.validate(); err != nil {
		return err
	}
	return nil
}

type RecaptchaV2EnterpriseTaskSolution struct {
	GRecaptchaResponse string            `json:"gRecaptchaResponse"`
	Cookies            map[string]string `json:"cookies"`
}
