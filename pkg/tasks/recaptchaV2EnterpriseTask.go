package tasks

type recaptchaV2EnterpriseTask struct {
	Type              string  `json:"type"`
	WebsiteURL        string  `json:"websiteURL"`
	WebsiteKey        string  `json:"websiteKey"`
	EnterprisePayload *string `json:"enterprisePayload,omitempty"`
	ApiDomain         *string `json:"apiDomain,omitempty"`
	ProxyType         string  `json:"proxyType"`
	ProxyAddress      string  `json:"proxyAddress"`
	ProxyPort         int     `json:"proxyPort"`
	ProxyLogin        *string `json:"proxyLogin,omitempty"`
	ProxyPassword     *string `json:"proxyPassword,omitempty"`
	UserAgent         *string `json:"userAgent,omitempty"`
	Cookies           *string `json:"cookies,omitempty"`
}

func NewRecaptchaV2EnterpriseTask(websiteURL, websiteKey, proxyType, proxyAddress string, proxyPort int) recaptchaV2EnterpriseTask {
	return recaptchaV2EnterpriseTask{Type: "RecaptchaV2EnterpriseTask", WebsiteURL: websiteURL, WebsiteKey: websiteKey,
		ProxyAddress: proxyAddress, ProxyType: proxyType, ProxyPort: proxyPort}
}

func (t recaptchaV2EnterpriseTask) WithEnterprisePayload(enterprisePayload string) recaptchaV2EnterpriseTask {
	t.EnterprisePayload = &enterprisePayload
	return t
}

func (t recaptchaV2EnterpriseTask) WithApiDomain(apiDomain string) recaptchaV2EnterpriseTask {
	t.ApiDomain = &apiDomain
	return t
}

func (t recaptchaV2EnterpriseTask) WithProxyLogin(proxyLogin string) recaptchaV2EnterpriseTask {
	t.ProxyLogin = &proxyLogin
	return t
}

func (t recaptchaV2EnterpriseTask) WithProxyPassword(proxyPassword string) recaptchaV2EnterpriseTask {
	t.ProxyPassword = &proxyPassword
	return t
}

func (t recaptchaV2EnterpriseTask) WithUserAgent(userAgent string) recaptchaV2EnterpriseTask {
	t.UserAgent = &userAgent
	return t
}

func (t recaptchaV2EnterpriseTask) WithCookies(cookies string) recaptchaV2EnterpriseTask {
	t.Cookies = &cookies
	return t
}
