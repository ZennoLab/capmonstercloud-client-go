package tasks

type recaptchaV2EnterpriseTaskProxyless struct {
	Type              string  `json:"type"`
	WebsiteURL        string  `json:"websiteURL"`
	WebsiteKey        string  `json:"websiteKey"`
	EnterprisePayload *string `json:"enterprisePayload,omitempty"`
	ApiDomain         *string `json:"apiDomain,omitempty"`
}

func NewRecaptchaV2EnterpriseTaskProxyless(websiteURL, websiteKey string) recaptchaV2EnterpriseTaskProxyless {
	return recaptchaV2EnterpriseTaskProxyless{Type: "RecaptchaV2EnterpriseTaskProxyless", WebsiteURL: websiteURL, WebsiteKey: websiteKey}
}

func (t recaptchaV2EnterpriseTaskProxyless) WithEnterprisePayload(enterprisePayload string) recaptchaV2EnterpriseTaskProxyless {
	t.EnterprisePayload = &enterprisePayload
	return t
}

func (t recaptchaV2EnterpriseTaskProxyless) WithApiDomain(apiDomain string) recaptchaV2EnterpriseTaskProxyless {
	t.ApiDomain = &apiDomain
	return t
}
