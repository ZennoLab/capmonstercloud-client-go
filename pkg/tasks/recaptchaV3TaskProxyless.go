package tasks

import (
	"fmt"
	"math"
	"net/url"
)

type RecaptchaV3TaskProxyless struct {
	Type       string   `json:"type"`
	WebsiteURL string   `json:"websiteURL"`
	WebsiteKey string   `json:"websiteKey"`
	MinScore   *float64 `json:"minScore,omitempty"`
	PageAction *string  `json:"pageAction,omitempty"`
}

func NewRecaptchaV3TaskProxyless(websiteURL, websiteKey string) RecaptchaV3TaskProxyless {
	return RecaptchaV3TaskProxyless{Type: "RecaptchaV3TaskProxyless", WebsiteURL: websiteURL, WebsiteKey: websiteKey}
}

func (t RecaptchaV3TaskProxyless) WithMinScore(minScore float64) RecaptchaV3TaskProxyless {
	t.MinScore = &minScore
	return t
}

func (t RecaptchaV3TaskProxyless) WithPageAction(pageAction string) RecaptchaV3TaskProxyless {
	t.PageAction = &pageAction
	return t
}

func (t RecaptchaV3TaskProxyless) Validate() error {
	if _, err := url.ParseRequestURI(t.WebsiteURL); err != nil {
		return fmt.Errorf("parse WebsiteURL: %w", err)
	}

	if len(t.WebsiteKey) < 1 || len(t.WebsiteKey) > math.MaxInt {
		return fmt.Errorf("WebsiteKey len error")
	}

	if t.MinScore != nil && (*t.MinScore < 0.1 || *t.MinScore > 0.9) {
		return fmt.Errorf("MinScore is not in [0.1,0.9] range")
	}
	return nil
}

type RecaptchaV3TaskTaskSolution struct {
	GRecaptchaResponse string            `json:"gRecaptchaResponse"`
	Cookies            map[string]string `json:"cookies"`
}
