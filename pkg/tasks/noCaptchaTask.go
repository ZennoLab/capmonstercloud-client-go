package tasks

type noCaptchaTaskProxyless struct {
	Type                string  `json:"type"`
	WebsiteURL          string  `json:"websiteURL"`
	WebsiteKey          string  `json:"websiteKey"`
	RecaptchaDataSValue *string `json:"recaptchaDataSValue,omitempty"`
	userAgentAndCookies
}

func NewNoCaptchaTaskProxyless(websiteURL, websiteKey, userAgent string) noCaptchaTaskProxyless {
	return noCaptchaTaskProxyless{
		Type:       "NoCaptchaTaskProxyless",
		WebsiteURL: websiteURL,
		WebsiteKey: websiteKey,
		userAgentAndCookies: userAgentAndCookies{
			UserAgent: &userAgent,
		},
	}
}

func (t noCaptchaTaskProxyless) WithRecaptchaDataSValue(recaptchaDataSValue string) noCaptchaTaskProxyless {
	t.RecaptchaDataSValue = &recaptchaDataSValue
	return t
}

type noCaptchaTask struct {
	noCaptchaTaskProxyless
	taskProxy
}

func NewNoCaptchaTask(websiteURL, websiteKey, proxyType, proxyAddress string, proxyPort int) noCaptchaTask {
	return noCaptchaTask{
		noCaptchaTaskProxyless: noCaptchaTaskProxyless{
			Type:       "NoCaptchaTask",
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
