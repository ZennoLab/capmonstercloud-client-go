package tasks

type funCaptchaTaskProxyless struct {
	Type                     string  `json:"type"`
	WebsiteURL               string  `json:"websiteURL"`
	FuncaptchaApiJSSubdomain *string `json:"funcaptchaApiJSSubdomain"`
	WebsitePublicKey         string  `json:"websitePublicKey"`
	Data                     *string `json:"data"`
}

func NewFunCaptchaTaskProxyless(websiteURL, websitePublicKey string) funCaptchaTaskProxyless {
	return funCaptchaTaskProxyless{Type: "FunCaptchaTaskProxyless", WebsiteURL: websiteURL, WebsitePublicKey: websitePublicKey}
}

func (t funCaptchaTaskProxyless) WithFuncaptchaApiJSSubdomain(funcaptchaApiJSSubdomain string) funCaptchaTaskProxyless {
	t.FuncaptchaApiJSSubdomain = &funcaptchaApiJSSubdomain
	return t
}

func (t funCaptchaTaskProxyless) WithData(data string) funCaptchaTaskProxyless {
	t.Data = &data
	return t
}
