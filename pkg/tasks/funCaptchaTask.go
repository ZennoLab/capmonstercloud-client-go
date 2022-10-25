package tasks

type FunCaptchaTaskProxyless struct {
	Type                     string  `json:"type"`
	WebsiteURL               string  `json:"websiteURL"`
	FuncaptchaApiJSSubdomain *string `json:"funcaptchaApiJSSubdomain"`
	WebsitePublicKey         string  `json:"websitePublicKey"`
	Data                     *string `json:"data"`
}

func NewFunCaptchaTaskProxyless(websiteURL, websitePublicKey string) FunCaptchaTaskProxyless {
	return FunCaptchaTaskProxyless{
		Type:             "FunCaptchaTaskProxyless",
		WebsiteURL:       websiteURL,
		WebsitePublicKey: websitePublicKey,
	}
}

func (t FunCaptchaTaskProxyless) WithFuncaptchaApiJSSubdomain(funcaptchaApiJSSubdomain string) FunCaptchaTaskProxyless {
	t.FuncaptchaApiJSSubdomain = &funcaptchaApiJSSubdomain
	return t
}

func (t FunCaptchaTaskProxyless) WithData(data string) FunCaptchaTaskProxyless {
	t.Data = &data
	return t
}

type FunCaptchaTask struct {
	FunCaptchaTaskProxyless
	taskProxy
	userAgentAndCookies
}

func NewFunCaptchaTask(websiteURL, websitePublicKey, proxyType, proxyAddress, userAgent string, proxyPort int) FunCaptchaTask {
	return FunCaptchaTask{
		FunCaptchaTaskProxyless: FunCaptchaTaskProxyless{
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
