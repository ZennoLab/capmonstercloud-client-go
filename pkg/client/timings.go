package client

import "time"

type resolveCapTimings struct {
	firstRequestDelay        time.Duration
	firstRequestNoCacheDelay time.Duration
	requestsInterval         time.Duration
	timeout                  time.Duration
}

var (
	imageToTextTimings = resolveCapTimings{
		firstRequestDelay:        350 * time.Millisecond,
		firstRequestNoCacheDelay: 350 * time.Millisecond,
		requestsInterval:         200 * time.Millisecond,
		timeout:                  10 * time.Second,
	}

	recaptchaV2TaskTimings = resolveCapTimings{
		firstRequestDelay:        1 * time.Second,
		firstRequestNoCacheDelay: 10 * time.Second,
		requestsInterval:         3 * time.Second,
		timeout:                  180 * time.Second,
	}

	recaptchaV3Timings = resolveCapTimings{
		firstRequestDelay:        1 * time.Second,
		firstRequestNoCacheDelay: 10 * time.Second,
		requestsInterval:         3 * time.Second,
		timeout:                  180 * time.Second,
	}

	recaptchaV2EnterpriseTimings = resolveCapTimings{
		firstRequestDelay:        1 * time.Second,
		firstRequestNoCacheDelay: 10 * time.Second,
		requestsInterval:         3 * time.Second,
		timeout:                  180 * time.Second,
	}

	funCaptchaTimings = resolveCapTimings{
		firstRequestDelay:        1 * time.Second,
		firstRequestNoCacheDelay: 10 * time.Second,
		requestsInterval:         1 * time.Second,
		timeout:                  80 * time.Second,
	}

	hCaptchaTimings = resolveCapTimings{
		firstRequestDelay:        1 * time.Second,
		firstRequestNoCacheDelay: 10 * time.Second,
		requestsInterval:         3 * time.Second,
		timeout:                  180 * time.Second,
	}

	geeTestTimings = resolveCapTimings{
		firstRequestDelay:        1 * time.Second,
		firstRequestNoCacheDelay: 1 * time.Second,
		requestsInterval:         1 * time.Second,
		timeout:                  80 * time.Second,
	}

	turnstileTimings = resolveCapTimings{
		firstRequestDelay:        1 * time.Second,
		firstRequestNoCacheDelay: 10 * time.Second,
		requestsInterval:         1 * time.Second,
		timeout:                  80 * time.Second,
	}
)
