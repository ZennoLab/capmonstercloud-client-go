package tasks

type recaptchaV3TaskProxyless struct {
	Type       string   `json:"type"`
	WebsiteURL string   `json:"websiteURL"`
	WebsiteKey string   `json:"websiteKey"`
	MinScore   *float64 `json:"minScore,omitempty"`
	PageAction *string  `json:"pageAction,omitempty"`
}

func NewRecaptchaV3TaskProxyless(websiteURL, websiteKey string) recaptchaV3TaskProxyless {
	return recaptchaV3TaskProxyless{Type: "RecaptchaV3TaskProxyless", WebsiteURL: websiteURL, WebsiteKey: websiteKey}
}

func (t recaptchaV3TaskProxyless) WithMinScore(minScore float64) recaptchaV3TaskProxyless {
	t.MinScore = &minScore
	return t
}

func (t recaptchaV3TaskProxyless) WithPageAction(pageAction string) recaptchaV3TaskProxyless {
	t.PageAction = &pageAction
	return t
}
