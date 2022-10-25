package tasks

type funCaptchaTaskProxyless struct {
	Type                     string  `json:"type"`
	WebsiteURL               string  `json:"websiteURL"`
	FuncaptchaApiJSSubdomain *string `json:"funcaptchaApiJSSubdomain"`
	WebsitePublicKey         string  `json:"websitePublicKey"`
	Data                     *string `json:"data"`
}

func NewFunCaptchaTaskProxyless(websiteURL, websitePublicKey string) funCaptchaTaskProxyless {
	return funCaptchaTaskProxyless{
		Type:             "FunCaptchaTaskProxyless",
		WebsiteURL:       websiteURL,
		WebsitePublicKey: websitePublicKey,
	}
}

func (t funCaptchaTaskProxyless) WithFuncaptchaApiJSSubdomain(funcaptchaApiJSSubdomain string) funCaptchaTaskProxyless {
	t.FuncaptchaApiJSSubdomain = &funcaptchaApiJSSubdomain
	return t
}

func (t funCaptchaTaskProxyless) WithData(data string) funCaptchaTaskProxyless {
	t.Data = &data
	return t
}

type funCaptchaTask struct {
	funCaptchaTaskProxyless
	taskProxy
	userAgentAndCookies
}

func NewFunCaptchaTask(websiteURL, websitePublicKey, proxyType, proxyAddress, userAgent string, proxyPort int) funCaptchaTask {
	return funCaptchaTask{
		funCaptchaTaskProxyless: funCaptchaTaskProxyless{
			Type:             "FunCaptchaTask",
			WebsiteURL:       websiteURL,
			WebsitePublicKey: websitePublicKey,
		},
		taskProxy: taskProxy{
			ProxyType:    proxyType,
			ProxyAddress: proxyAddress,
			ProxyPort:    proxyPort,
		},
		userAgentAndCookies: userAgentAndCookies{
			UserAgent: &userAgent,
		},
	}
}
