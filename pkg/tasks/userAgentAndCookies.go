package tasks

type userAgentAndCookies struct {
	UserAgent *string `json:"userAgent,omitempty"`
	Cookies   *string `json:"cookies,omitempty"`
}

func (t userAgentAndCookies) WithCookies(cookies string) userAgentAndCookies {
	t.Cookies = &cookies
	return t
}

func (t userAgentAndCookies) WithUserAgent(userAgent string) userAgentAndCookies {
	t.UserAgent = &userAgent
	return t
}
