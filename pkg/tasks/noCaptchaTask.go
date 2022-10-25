package tasks

type NoCaptchaTaskProxyless struct {
	Type                string  `json:"type"`
	WebsiteURL          string  `json:"websiteURL"`
	WebsiteKey          string  `json:"websiteKey"`
	RecaptchaDataSValue *string `json:"recaptchaDataSValue,omitempty"`
	userAgentAndCookies
}

func NewNoCaptchaTaskProxyless(websiteURL, websiteKey, userAgent string) NoCaptchaTaskProxyless {
	return NoCaptchaTaskProxyless{
		Type:       "NoCaptchaTaskProxyless",
		WebsiteURL: websiteURL,
		WebsiteKey: websiteKey,
		userAgentAndCookies: userAgentAndCookies{
			UserAgent: &userAgent,
		},
	}
}

func (t NoCaptchaTaskProxyless) WithRecaptchaDataSValue(recaptchaDataSValue string) NoCaptchaTaskProxyless {
	t.RecaptchaDataSValue = &recaptchaDataSValue
	return t
}

type NoCaptchaTask struct {
	NoCaptchaTaskProxyless
	taskProxy
}

func NewNoCaptchaTask(websiteURL, websiteKey, proxyType, proxyAddress string, proxyPort int) NoCaptchaTask {
	return NoCaptchaTask{
		NoCaptchaTaskProxyless: NoCaptchaTaskProxyless{
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

type NoCaptchaTaskResult struct {
	result
	Solution struct {
		GRecaptchaResponse string            `json:"gRecaptchaResponse"`
		Cookies            map[string]string `json:"cookies"`
	} `json:"solution"`
}
