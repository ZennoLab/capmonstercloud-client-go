package tasks

type recaptchaV2EnterpriseTaskProxyless struct {
	Type              string  `json:"type"`
	WebsiteURL        string  `json:"websiteURL"`
	WebsiteKey        string  `json:"websiteKey"`
	EnterprisePayload *string `json:"enterprisePayload,omitempty"`
	ApiDomain         *string `json:"apiDomain,omitempty"`
}

func NewRecaptchaV2EnterpriseTaskProxyless(websiteURL, websiteKey string) recaptchaV2EnterpriseTaskProxyless {
	return recaptchaV2EnterpriseTaskProxyless{
		Type:       "RecaptchaV2EnterpriseTaskProxyless",
		WebsiteURL: websiteURL,
		WebsiteKey: websiteKey,
	}
}

func (t recaptchaV2EnterpriseTaskProxyless) WithEnterprisePayload(enterprisePayload string) recaptchaV2EnterpriseTaskProxyless {
	t.EnterprisePayload = &enterprisePayload
	return t
}

func (t recaptchaV2EnterpriseTaskProxyless) WithApiDomain(apiDomain string) recaptchaV2EnterpriseTaskProxyless {
	t.ApiDomain = &apiDomain
	return t
}

type recaptchaV2EnterpriseTask struct {
	recaptchaV2EnterpriseTaskProxyless
	taskProxy
	userAgentAndCookies
}

func NewRecaptchaV2EnterpriseTask(websiteURL, websiteKey, proxyType, proxyAddress string, proxyPort int) recaptchaV2EnterpriseTask {
	return recaptchaV2EnterpriseTask{
		recaptchaV2EnterpriseTaskProxyless: recaptchaV2EnterpriseTaskProxyless{
			Type:       "RecaptchaV2EnterpriseTask",
			WebsiteURL: websiteURL,
			WebsiteKey: websiteKey,
		},
		taskProxy: taskProxy{
			ProxyType:    proxyType,
			ProxyAddress: proxyAddress,
			ProxyPort:    proxyPort,
		},
	}
}
