package tasks

type noCaptchaTask struct {
	Type                string  `json:"type"`
	WebsiteURL          string  `json:"websiteURL"`
	WebsiteKey          string  `json:"websiteKey"`
	RecaptchaDataSValue *string `json:"recaptchaDataSValue,omitempty"`
	ProxyType           string  `json:"proxyType"`
	ProxyAddress        string  `json:"proxyAddress"`
	ProxyPort           int     `json:"proxyPort"`
	ProxyLogin          *string `json:"proxyLogin,omitempty"`
	ProxyPassword       *string `json:"proxyPassword,omitempty"`
	UserAgent           *string `json:"userAgent,omitempty"`
	Cookies             *string `json:"cookies,omitempty"`
}

func NewNoCaptchaTask(websiteURL, websiteKey, proxyType, proxyAddress string, proxyPort int) noCaptchaTask {
	return noCaptchaTask{Type: "NoCaptchaTask", WebsiteURL: websiteURL, WebsiteKey: websiteKey,
		ProxyType: proxyType, ProxyAddress: proxyAddress, ProxyPort: proxyPort}
}

func (t noCaptchaTask) WithRecaptchaDataSValue(recaptchaDataSValue string) noCaptchaTask {
	t.RecaptchaDataSValue = &recaptchaDataSValue
	return t
}

func (t noCaptchaTask) WithProxyLogin(proxyLogin string) noCaptchaTask {
	t.ProxyLogin = &proxyLogin
	return t
}

func (t noCaptchaTask) WithProxyPassword(proxyPassword string) noCaptchaTask {
	t.ProxyPassword = &proxyPassword
	return t
}

func (t noCaptchaTask) WithCookies(cookies string) noCaptchaTask {
	t.Cookies = &cookies
	return t
}
