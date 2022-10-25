package tasks

type noCaptchaTaskProxyless struct {
	Type                string  `json:"type"`
	WebsiteURL          string  `json:"websiteURL"`
	WebsiteKey          string  `json:"websiteKey"`
	RecaptchaDataSValue *string `json:"recaptchaDataSValue,omitempty"`
	UserAgent           string  `json:"userAgent"`
	Cookies             *string `json:"cookies,omitempty"`
}

func NewNoCaptchaTaskProxyless(websiteURL, websiteKey, userAgent string) noCaptchaTaskProxyless {
	return noCaptchaTaskProxyless{Type: "NoCaptchaTaskProxyless", WebsiteURL: websiteURL, WebsiteKey: websiteKey, UserAgent: userAgent}
}

func (t noCaptchaTaskProxyless) WithRecaptchaDataSValue(recaptchaDataSValue string) noCaptchaTaskProxyless {
	t.RecaptchaDataSValue = &recaptchaDataSValue
	return t
}

func (t noCaptchaTaskProxyless) WithCookies(cookies string) noCaptchaTaskProxyless {
	t.Cookies = &cookies
	return t
}
