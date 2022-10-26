package client

import "errors"

var (
	errMap = map[string]error{
		"ERROR_KEY_DOES_NOT_EXIST":        errors.New("INVALID KEY"),
		"ERROR_ZERO_BALANCE":              errors.New("NO FUNDS"),
		"ERROR_TOO_BIG_CAPTCHA_FILESIZE":  errors.New("BIG IMAGE SIZE"),
		"ERROR_ZERO_CAPTCHA_FILESIZE":     errors.New("ZERO IMAGE SIZE"),
		"ERROR_NO_SUCH_CAPCHA_ID":         errors.New("CAPTCHA ID IS NOT FOUND"),
		"WRONG_CAPTCHA_ID":                errors.New("CAPTCHA ID IS NOT FOUND"),
		"ERROR_CAPTCHA_UNSOLVABLE":        errors.New("UNSUPPORTED CAPTCHA TYPE"),
		"CAPTCHA_NOT_READY":               errors.New("CAPTCHA IS NOT READY"),
		"ERROR_IP_NOT_ALLOWED":            errors.New("REQUEST IS NOT ALLOWED FROM YOUR IP"),
		"ERROR_IP_BANNED":                 errors.New("IP BANNED"),
		"ERROR_NO_SUCH_METHOD":            errors.New("INCORRECT METHOD"),
		"ERROR_TOO_MUCH_REQUESTS":         errors.New("REQUEST LIMIT EXCEEDED"),
		"ERROR_DOMAIN_NOT_ALLOWED":        errors.New("THE DOMAIN IS NOT ALLOWED"),
		"ERROR_TOKEN_EXPIRED":             errors.New("THE TOKEN IS EXPIRED"),
		"ERROR_NO_SLOT_AVAILABLE":         errors.New("NO FREE SERVERS"),
		"ERROR_RECAPTCHA_INVALID_SITEKEY": errors.New("INVALID RECAPTCHA SITEKEY"),
		"ERROR_RECAPTCHA_INVALID_DOMAIN":  errors.New("INVALID RECAPTCHA DOMAIN"),
		"ERROR_RECAPTCHA_TIMEOUT":         errors.New("RECAPTCHA TIMEOUT"),
		"ERROR_IP_BLOCKED":                errors.New("YOUR IP IS BLOCKED"),
		"ERROR_PROXY_CONNECT_REFUSED":     errors.New("FAILED TO CONNECT PROXY"),
		"ERROR_PROXY_BANNED":              errors.New("THE PROXY IP IS BANNED"),
		"ERROR_TASK_NOT_SUPPORTED":        errors.New("INCORRECT TASK TYPE"),
	}

	errUnknown                   = errors.New("unknown error")
	errTimeout                   = errors.New("waiting timeout exceeded")
	errServiceServiceUnavailable = errors.New("service unavailable")
)
