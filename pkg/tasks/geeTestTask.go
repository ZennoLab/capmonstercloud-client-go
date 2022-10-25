package tasks

type GeeTestTaskProxyless struct {
	Type                      string  `json:"type"`
	WebsiteURL                string  `json:"websiteURL"`
	Gt                        string  `json:"gt"`
	Challenge                 string  `json:"challenge"`
	GeetestApiServerSubdomain *string `json:"geetestApiServerSubdomain,omitempty"`
	GeetestGetLib             *string `json:"geetestGetLib,omitempty"`
	userAgentAndCookies
}

func NewGeeTestTaskProxyless(websiteURL, gt, challenge string) GeeTestTaskProxyless {
	return GeeTestTaskProxyless{
		Type:      "GeeTestTaskProxyless",
		Gt:        gt,
		Challenge: challenge,
	}
}

func (t GeeTestTaskProxyless) WithGeetestApiServerSubdomain(geetestApiServerSubdomain string) GeeTestTaskProxyless {
	t.GeetestApiServerSubdomain = &geetestApiServerSubdomain
	return t
}

func (t GeeTestTaskProxyless) WithGeetestGetLib(geetestGetLib string) GeeTestTaskProxyless {
	t.GeetestGetLib = &geetestGetLib
	return t
}

type GeeTestTask struct {
	GeeTestTaskProxyless
	taskProxy
}

func NewGeeTestTask(websiteURL, gt, challenge, proxyType, proxyAddress string, proxyPort int) GeeTestTask {
	return GeeTestTask{
		GeeTestTaskProxyless: GeeTestTaskProxyless{
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

type GeeTestTaskSolution struct {
	Challenge string            `json:"challenge"`
	Validate  string            `json:"validate"`
	Seccode   string            `json:"seccode"`
	Cookies   map[string]string `json:"cookies"`
}
type GeeTestTaskResult struct {
	result
	Solution GeeTestTaskSolution `json:"solution"`
}
