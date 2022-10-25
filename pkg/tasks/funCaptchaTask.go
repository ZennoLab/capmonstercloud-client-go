package tasks

type funCaptchaTask struct {
	Type                     string  `json:"type"`
	WebsiteURL               string  `json:"websiteURL"`
	FuncaptchaApiJSSubdomain *string `json:"funcaptchaApiJSSubdomain"`
	WebsitePublicKey         string  `json:"websitePublicKey"`
	Data                     *string `json:"data"`
	ProxyType                string  `json:"proxyType"`
	ProxyAddress             string  `json:"proxyAddress"`
	ProxyPort                int     `json:"proxyPort"`
	ProxyLogin               *string `json:"proxyLogin,omitempty"`
	ProxyPassword            *string `json:"proxyPassword,omitempty"`
	UserAgent                *string `json:"userAgent,omitempty"`
	Cookies                  *string `json:"cookies,omitempty"`
}

func NewFunCaptchaTask(websiteURL, websitePublicKey, proxyType, proxyAddress string, proxyPort int) funCaptchaTask {
	return funCaptchaTask{Type: "FunCaptchaTask", WebsiteURL: websiteURL, WebsitePublicKey: websitePublicKey,
		ProxyAddress: proxyAddress, ProxyType: proxyType, ProxyPort: proxyPort}
}

func (t funCaptchaTask) WithFuncaptchaApiJSSubdomain(funcaptchaApiJSSubdomain string) funCaptchaTask {
	t.FuncaptchaApiJSSubdomain = &funcaptchaApiJSSubdomain
	return t
}

func (t funCaptchaTask) WithData(data string) funCaptchaTask {
	t.Data = &data
	return t
}

func (t funCaptchaTask) WithProxyLogin(proxyLogin string) funCaptchaTask {
	t.ProxyLogin = &proxyLogin
	return t
}

func (t funCaptchaTask) WithProxyPassword(proxyPassword string) funCaptchaTask {
	t.ProxyPassword = &proxyPassword
	return t
}

func (t funCaptchaTask) WithUserAgent(userAgent string) funCaptchaTask {
	t.UserAgent = &userAgent
	return t
}

func (t funCaptchaTask) WithCookies(cookies string) funCaptchaTask {
	t.Cookies = &cookies
	return t
}
