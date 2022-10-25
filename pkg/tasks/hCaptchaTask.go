package tasks

type HCaptchaTaskProxyless struct {
	Type        string  `json:"type"`
	WebsiteURL  string  `json:"websiteURL"`
	WebsiteKey  string  `json:"websiteKey"`
	IsInvisible *bool   `json:"isInvisible,omitempty"`
	Data        *string `json:"data,omitempty"`
	userAgentAndCookies
}

func NewHCaptchaTaskProxyless(websiteURL, websiteKey, userAgent string) HCaptchaTaskProxyless {
	return HCaptchaTaskProxyless{
		Type:       "HCaptchaTaskProxyless",
		WebsiteURL: websiteURL,
		WebsiteKey: websiteKey,
		userAgentAndCookies: userAgentAndCookies{
			UserAgent: &userAgent,
		},
	}
}

func (t HCaptchaTaskProxyless) WithIsInvisible(isInvisible bool) HCaptchaTaskProxyless {
	t.IsInvisible = &isInvisible
	return t
}

func (t HCaptchaTaskProxyless) WithData(data string) HCaptchaTaskProxyless {
	t.Data = &data
	return t
}

type HCaptchaTask struct {
	HCaptchaTaskProxyless
	taskProxy
}

func NewHCaptchaTask(websiteURL, websiteKey, proxyType, proxyAddress string, proxyPort int) HCaptchaTask {
	return HCaptchaTask{
		HCaptchaTaskProxyless: HCaptchaTaskProxyless{
			Type:       "HCaptchaTask",
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

type HCaptchaTaskResult struct {
	result
	Solution struct {
		GRecaptchaResponse string            `json:"gRecaptchaResponse"`
		Cookies            map[string]string `json:"cookies"`
	} `json:"solution"`
}
