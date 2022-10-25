package tasks

type geeTestTaskProxyless struct {
	Type                      string  `json:"type"`
	WebsiteURL                string  `json:"websiteURL"`
	Gt                        string  `json:"gt"`
	Challenge                 string  `json:"challenge"`
	GeetestApiServerSubdomain *string `json:"geetestApiServerSubdomain,omitempty"`
	GeetestGetLib             *string `json:"geetestGetLib,omitempty"`
	userAgentAndCookies
}

func NewGeeTestTaskProxyless(websiteURL, gt, challenge string) geeTestTaskProxyless {
	return geeTestTaskProxyless{
		Type:      "GeeTestTaskProxyless",
		Gt:        gt,
		Challenge: challenge,
	}
}

func (t geeTestTaskProxyless) WithGeetestApiServerSubdomain(geetestApiServerSubdomain string) geeTestTaskProxyless {
	t.GeetestApiServerSubdomain = &geetestApiServerSubdomain
	return t
}

func (t geeTestTaskProxyless) WithGeetestGetLib(geetestGetLib string) geeTestTaskProxyless {
	t.GeetestGetLib = &geetestGetLib
	return t
}

type geeTestTask struct {
	geeTestTaskProxyless
	taskProxy
}

func NewGeeTestTask(websiteURL, gt, challenge, proxyType, proxyAddress string, proxyPort int) geeTestTask {
	return geeTestTask{
		geeTestTaskProxyless: geeTestTaskProxyless{
			Type:      "GeeTestTask",
			Gt:        gt,
			Challenge: challenge,
		},
		taskProxy: taskProxy{
			ProxyType:    proxyType,
			ProxyAddress: proxyAddress,
			ProxyPort:    proxyPort,
		},
	}
}
