package tasks

import (
	"fmt"
	"math"
	"net/url"
)

type FunCaptchaTaskProxyless struct {
	Type                     string  `json:"type"`
	WebsiteURL               string  `json:"websiteURL"`
	FuncaptchaApiJSSubdomain *string `json:"funcaptchaApiJSSubdomain"`
	WebsitePublicKey         string  `json:"websitePublicKey"`
	Data                     *string `json:"data"`
}

func NewFunCaptchaTaskProxyless(websiteURL, websitePublicKey string) FunCaptchaTaskProxyless {
	return FunCaptchaTaskProxyless{
		Type:             "FunCaptchaTaskProxyless",
		WebsiteURL:       websiteURL,
		WebsitePublicKey: websitePublicKey,
	}
}

func (t FunCaptchaTaskProxyless) WithFuncaptchaApiJSSubdomain(funcaptchaApiJSSubdomain string) FunCaptchaTaskProxyless {
	t.FuncaptchaApiJSSubdomain = &funcaptchaApiJSSubdomain
	return t
}

func (t FunCaptchaTaskProxyless) WithData(data string) FunCaptchaTaskProxyless {
	t.Data = &data
	return t
}

func (t FunCaptchaTaskProxyless) Validate() error {
	if _, err := url.ParseRequestURI(t.WebsiteURL); err != nil {
		return fmt.Errorf("parse WebsiteURL: %w", err)
	}
	if len(t.WebsitePublicKey) < 1 || len(t.WebsitePublicKey) > math.MaxInt {
		return fmt.Errorf("WebsitePublicKey len error")
	}
	return nil
}

type FunCaptchaTask struct {
	FunCaptchaTaskProxyless
	taskProxy
	userAgentAndCookies
}

func NewFunCaptchaTask(websiteURL, websitePublicKey, proxyType, proxyAddress, userAgent string, proxyPort int) FunCaptchaTask {
	return FunCaptchaTask{
		FunCaptchaTaskProxyless: FunCaptchaTaskProxyless{
			Type:             "FunCaptchaTask",
			WebsiteURL:       websiteURL,
			WebsitePublicKey: websitePublicKey,
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

func (t FunCaptchaTask) Validate() error {
	if err := t.FunCaptchaTaskProxyless.Validate(); err != nil {
		return err
	}
	if err := t.taskProxy.validate(); err != nil {
		return err
	}
	return nil
}

type FunCaptchaTaskSolution struct {
	Token   string            `json:"Token"`
	Cookies map[string]string `json:"cookies"`
}
