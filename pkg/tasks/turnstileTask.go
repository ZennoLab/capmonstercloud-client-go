package tasks

import (
	"math"
	"net/url"
)

type TurnstileTaskProxyless struct {
	Type               string  `json:"type"`
	WebsiteURL         string  `json:"websiteURL"`
	WebsiteKey         string  `json:"websiteKey"`
	CloudflareTaskType *string `json:"cloudflareTaskType,omitempty"`
	PageAction         *string `json:"PageAction,omitempty"`
	Data               *string `json:"Data,omitempty"`
	PageData           *string `json:"PageData,omitempty"`
	HtmlPageBase64     *string `json:"htmlPageBase64,omitempty"`
	userAgentAndCookies
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

	CfClearanceType := string("cf_clearance")
	TokenType := string("token")

	if t.CloudflareTaskType != nil && *t.CloudflareTaskType != CfClearanceType && *t.CloudflareTaskType != TokenType {
		return ErrCloudflareTaskType
	}

	if t.CloudflareTaskType != nil && t.UserAgent == nil {
		return ErrUserAgentRequired
	}

	if t.CloudflareTaskType != nil && *t.CloudflareTaskType == TokenType && t.PageData == nil {
		return ErrCloudflarePageData
	}

	if t.CloudflareTaskType != nil && *t.CloudflareTaskType == TokenType && t.PageAction == nil {
		return ErrCloudflarePageAction
	}

	if t.CloudflareTaskType != nil && *t.CloudflareTaskType == TokenType && t.Data == nil {
		return ErrCloudflareData
	}

	if t.CloudflareTaskType != nil && *t.CloudflareTaskType == CfClearanceType && t.HtmlPageBase64 == nil {
		return ErrCloudflareHtmlPageBase64
	}

	return nil
}

type TurnstileTask struct {
	TurnstileTaskProxyless
	taskProxy
	userAgentAndCookies
}

func (t TurnstileTaskProxyless) WithCloudflareTaskType(cloudflareTaskType string) TurnstileTaskProxyless {
	t.CloudflareTaskType = &cloudflareTaskType
	return t
}

func (t TurnstileTaskProxyless) WithPageAction(pageAction string) TurnstileTaskProxyless {
	t.PageAction = &pageAction
	return t
}

func (t TurnstileTaskProxyless) WithPageData(pageData string) TurnstileTaskProxyless {
	t.PageData = &pageData
	return t
}

func (t TurnstileTaskProxyless) WithData(data string) TurnstileTaskProxyless {
	t.Data = &data
	return t
}

func (t TurnstileTaskProxyless) WithHtmlPageBase64(htmlPageBase64 string) TurnstileTaskProxyless {
	t.HtmlPageBase64 = &htmlPageBase64
	return t
}

func (t TurnstileTaskProxyless) WithUserAgent(userAgent string) TurnstileTaskProxyless {
	t.UserAgent = &userAgent
	return t
}

func NewTurnstileTask(websiteURL, websiteKey, proxyType, proxyAddress string, proxyPort int, userAgent string) TurnstileTask {
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
		userAgentAndCookies: userAgentAndCookies{
			UserAgent: &userAgent,
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

type TurnstileTaskSolution struct {
	Token string `json:"token"`
}
