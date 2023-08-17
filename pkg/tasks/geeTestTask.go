package tasks

import (
	"math"
	"net/url"
)

type GeeTestTaskProxyless struct {
	Type                      string       `json:"type"`
	WebsiteURL                string       `json:"websiteURL"`
	Gt                        string       `json:"gt"`
	Challenge                 *string      `json:"challenge,omitempty"`
	Version                   *int         `json:"version,omitempty"`
	InitParameters            *interface{} `json:"initParameters,omitempty"`
	GeetestApiServerSubdomain *string      `json:"geetestApiServerSubdomain,omitempty"`
	GeetestGetLib             *string      `json:"geetestGetLib,omitempty"`
	userAgentAndCookies
}

func NewGeeTestTaskProxyless(websiteURL string, gt string) GeeTestTaskProxyless {
	return GeeTestTaskProxyless{
		Type:       "GeeTestTaskProxyless",
		WebsiteURL: websiteURL,
		Gt:         gt,
	}
}

func (t GeeTestTaskProxyless) WithChallenge(challenge string) GeeTestTaskProxyless {
	t.Challenge = &challenge
	return t
}

func (t GeeTestTaskProxyless) WithGeetestApiServerSubdomain(geetestApiServerSubdomain string) GeeTestTaskProxyless {
	t.GeetestApiServerSubdomain = &geetestApiServerSubdomain
	return t
}

func (t GeeTestTaskProxyless) WithGeetestGetLib(geetestGetLib string) GeeTestTaskProxyless {
	t.GeetestGetLib = &geetestGetLib
	return t
}

func (t GeeTestTaskProxyless) WithVersion(version int) GeeTestTaskProxyless {
	t.Version = &version
	return t
}

func (t GeeTestTaskProxyless) WithInitParametres(initParameters interface{}) GeeTestTaskProxyless {
	t.InitParameters = &initParameters
	return t
}

func (t GeeTestTaskProxyless) WithUserAgent(userAgent string) GeeTestTaskProxyless {
	t.UserAgent = &userAgent
	return t
}

func (t GeeTestTaskProxyless) Validate() error {
	if _, err := url.ParseRequestURI(t.WebsiteURL); err != nil {
		return ErrInvalidWebsiteUrl
	}
	if len(t.Gt) < 1 || len(t.Gt) > math.MaxInt {
		return ErrInvalidGt
	}
	if t.Version == nil && t.Challenge == nil {
		return ErrChallenge
	}
	return nil
}

type GeeTestTask struct {
	GeeTestTaskProxyless
	taskProxy
	userAgentAndCookies
}

func NewGeeTestTask(websiteURL, gt, proxyType, proxyAddress string, proxyPort int, userAgent string) GeeTestTask {
	return GeeTestTask{
		GeeTestTaskProxyless: GeeTestTaskProxyless{
			Type: "GeeTestTask",
			Gt:   gt,
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
	Challenge     *string            `json:"challenge"`
	Validate      *string            `json:"validate"`
	Seccode       *string            `json:"seccode"`
	Cookies       *map[string]string `json:"cookies"`
	CaptchaId     *string            `json:"captcha_id"`
	LotNumber     *string            `json:"lot_number"`
	PassToken     *string            `json:"pass_token"`
	GenTime       *string            `json:"gen_time"`
	CaptchaOutput *string            `json:"captcha_output"`
}
