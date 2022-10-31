package tasks

type taskProxy struct {
	ProxyType     string  `json:"proxyType"`
	ProxyAddress  string  `json:"proxyAddress"`
	ProxyPort     int     `json:"proxyPort"`
	ProxyLogin    *string `json:"proxyLogin,omitempty"`
	ProxyPassword *string `json:"proxyPassword,omitempty"`
}

func (t taskProxy) WithProxyLogin(proxyLogin string) taskProxy {
	t.ProxyLogin = &proxyLogin
	return t
}

func (t taskProxy) WithProxyPassword(proxyPassword string) taskProxy {
	t.ProxyPassword = &proxyPassword
	return t
}

func (t taskProxy) validate() error {
	if t.ProxyPort < 0 || t.ProxyPort > 65535 {
		return ErrInvalidProxyPort
	}
	return nil
}
