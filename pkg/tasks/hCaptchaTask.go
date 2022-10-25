package tasks

type hCaptchaTaskProxyless struct {
	Type        string  `json:"type"`
	WebsiteURL  string  `json:"websiteURL"`
	WebsiteKey  string  `json:"websiteKey"`
	IsInvisible *bool   `json:"isInvisible,omitempty"`
	Data        *string `json:"data,omitempty"`
	userAgentAndCookies
}

func NewHCaptchaTaskProxyless(websiteURL, websiteKey, userAgent string) hCaptchaTaskProxyless {
	return hCaptchaTaskProxyless{
		Type:       "HCaptchaTaskProxyless",
		WebsiteURL: websiteURL,
		WebsiteKey: websiteKey,
		userAgentAndCookies: userAgentAndCookies{
			UserAgent: &userAgent,
		},
	}
}

func (t hCaptchaTaskProxyless) WithIsInvisible(isInvisible bool) hCaptchaTaskProxyless {
	t.IsInvisible = &isInvisible
	return t
}

func (t hCaptchaTaskProxyless) WithData(data string) hCaptchaTaskProxyless {
	t.Data = &data
	return t
}

type hCaptchaTask struct {
	hCaptchaTaskProxyless
	taskProxy
}

func NewHCaptchaTask(websiteURL, websiteKey, proxyType, proxyAddress string, proxyPort int) hCaptchaTask {
	return hCaptchaTask{
		hCaptchaTaskProxyless: hCaptchaTaskProxyless{
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
