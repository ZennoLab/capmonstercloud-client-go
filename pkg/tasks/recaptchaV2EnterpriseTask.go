package tasks

type RecaptchaV2EnterpriseTaskProxyless struct {
	Type              string  `json:"type"`
	WebsiteURL        string  `json:"websiteURL"`
	WebsiteKey        string  `json:"websiteKey"`
	EnterprisePayload *string `json:"enterprisePayload,omitempty"`
	ApiDomain         *string `json:"apiDomain,omitempty"`
}

func NewRecaptchaV2EnterpriseTaskProxyless(websiteURL, websiteKey string) RecaptchaV2EnterpriseTaskProxyless {
	return RecaptchaV2EnterpriseTaskProxyless{
		Type:       "RecaptchaV2EnterpriseTaskProxyless",
		WebsiteURL: websiteURL,
		WebsiteKey: websiteKey,
	}
}

func (t RecaptchaV2EnterpriseTaskProxyless) WithEnterprisePayload(enterprisePayload string) RecaptchaV2EnterpriseTaskProxyless {
	t.EnterprisePayload = &enterprisePayload
	return t
}

func (t RecaptchaV2EnterpriseTaskProxyless) WithApiDomain(apiDomain string) RecaptchaV2EnterpriseTaskProxyless {
	t.ApiDomain = &apiDomain
	return t
}

type RecaptchaV2EnterpriseTask struct {
	RecaptchaV2EnterpriseTaskProxyless
	taskProxy
	userAgentAndCookies
}

func NewRecaptchaV2EnterpriseTask(websiteURL, websiteKey, proxyType, proxyAddress string, proxyPort int) RecaptchaV2EnterpriseTask {
	return RecaptchaV2EnterpriseTask{
		RecaptchaV2EnterpriseTaskProxyless: RecaptchaV2EnterpriseTaskProxyless{
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

type RecaptchaV2EnterpriseTaskResult struct {
	GRecaptchaResponse string `json:"gRecaptchaResponse"`
}
