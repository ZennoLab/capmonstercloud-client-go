package tasks

import "errors"

const (
	ProxyTypeHttp   = "http"
	ProxyTypeHttps  = "https"
	ProxyTypeSocks4 = "socks4"
	ProxyTypeSocks5 = "socks5"

	CapMonsterModuleAmazon     = "amazon"
	CapMonsterModuleBotdetect  = "botdetect"
	CapMonsterModuleFacebook   = "facebook"
	CapMonsterModuleGmx        = "gmx"
	CapMonsterModuleGoogle     = "google"
	CapMonsterModuleHotmail    = "hotmail"
	CapMonsterModuleMailru     = "mailru"
	CapMonsterModuleOk         = "ok"
	CapMonsterModuleOknew      = "oknew"
	CapMonsterModuleRamblerrus = "ramblerrus"
	CapMonsterModuleSolvemedia = "solvemedia"
	CapMonsterModuleSteam      = "steam"
	CapMonsterModuleVk         = "vk"
	CapMonsterModuleYandex     = "yandex"
	CapMonsterModuleYandexnew  = "yandexnew"
	CapMonsterModuleYandexwave = "yandexwave"
	CapMonsterModuleUniversal  = "universal"
)

var (
	ErrInvalidWebsiteUrl           = errors.New("invalid WebsiteURL")
	ErrInvalidProxyPort            = errors.New("invalid proxy port")
	ErrInvalidMinScore             = errors.New("minScore is not in [0.1,0.9] range")
	ErrInvalidWebSiteKey           = errors.New("website key len error")
	ErrInvalidGt                   = errors.New("gt len error")
	ErrInvalidRecognizingThreshold = errors.New("recognizingThreshold is not in [0,100] range")
	ErrInvalidEnterprisePayload    = errors.New("enterprisePayload len error")

	ErrUserAgentRequired = errors.New("userAgent required")

	ErrMetaData = errors.New("invalid Metadata")

	ErrChallenge = errors.New("challenge required")
)
