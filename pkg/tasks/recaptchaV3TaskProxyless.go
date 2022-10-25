package tasks

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

type RecaptchaV3TaskTaskSolution struct {
	GRecaptchaResponse string            `json:"gRecaptchaResponse"`
	Cookies            map[string]string `json:"cookies"`
}

type RecaptchaV3TaskTaskResult struct {
	result
	Solution RecaptchaV3TaskTaskSolution `json:"solution"`
}
