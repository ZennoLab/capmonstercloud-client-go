package tasks

import (
	"math"
	"net/url"
)

type GeeTestTaskProxyless struct {
	Type                      string  `json:"type"`
	WebsiteURL                string  `json:"websiteURL"`
	Gt                        string  `json:"gt"`
	Challenge                 string  `json:"challenge"`
	GeetestApiServerSubdomain *string `json:"geetestApiServerSubdomain,omitempty"`
	GeetestGetLib             *string `json:"geetestGetLib,omitempty"`
	userAgentAndCookies
}

func NewGeeTestTaskProxyless(websiteURL, gt, challenge string) GeeTestTaskProxyless {
	return GeeTestTaskProxyless{
		Type:       "GeeTestTaskProxyless",
		WebsiteURL: websiteURL,
		Gt:         gt,
		Challenge:  challenge,
	}
}

func (t GeeTestTaskProxyless) WithGeetestApiServerSubdomain(geetestApiServerSubdomain string) GeeTestTaskProxyless {
	t.GeetestApiServerSubdomain = &geetestApiServerSubdomain
	return t
}

func (t GeeTestTaskProxyless) WithGeetestGetLib(geetestGetLib string) GeeTestTaskProxyless {
	t.GeetestGetLib = &geetestGetLib
	return t
}

func (t GeeTestTaskProxyless) Validate() error {
	if _, err := url.ParseRequestURI(t.WebsiteURL); err != nil {
		return ErrInvalidWebsiteUrl
	}
	if len(t.Gt) < 1 || len(t.Gt) > math.MaxInt {
		return ErrInvalidGt
	}
	return nil
}

type GeeTestTask struct {
	GeeTestTaskProxyless
	taskProxy
}

func NewGeeTestTask(websiteURL, gt, challenge, proxyType, proxyAddress string, proxyPort int) GeeTestTask {
	return GeeTestTask{
		GeeTestTaskProxyless: GeeTestTaskProxyless{
			Type:      "GeeTestTask",
			Gt:        gt,
			Challenge: challenge,
		},
		taskProxy: taskProxy{
			ProxyType:    proxyType,
			ProxyAddress: proxyAddress,
			ProxyPort:    proxyPort,
		},
	}
}

func (t GeeTestTask) Validate() error {
	if err := t.GeeTestTaskProxyless.Validate(); err != nil {
		return err
	}
	if err := t.taskProxy.validate(); err != nil {
		return err
	}
	return nil
}

type GeeTestTaskSolution struct {
	Challenge string            `json:"challenge"`
	Validate  string            `json:"validate"`
	Seccode   string            `json:"seccode"`
	Cookies   map[string]string `json:"cookies"`
}
