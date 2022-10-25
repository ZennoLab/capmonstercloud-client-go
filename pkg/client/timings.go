package client

import "time"

type resolveCapTimings struct {
	firstRequestDelay time.Duration
	requestsInterval  time.Duration
	timeout           time.Duration
}

var (
	imageToTextTimings = resolveCapTimings{
		firstRequestDelay: 150 * time.Millisecond,
		requestsInterval:  200 * time.Millisecond,
		timeout:           10 * time.Second,
	}

	noCaptchaTaskTimings = resolveCapTimings{
		firstRequestDelay: 7 * time.Second,
		requestsInterval:  3 * time.Second,
		timeout:           180 * time.Second,
	}

	recaptchaV3Timings = resolveCapTimings{
		firstRequestDelay: 7 * time.Second,
		requestsInterval:  3 * time.Second,
		timeout:           180 * time.Second,
	}

	recaptchaV2EnterpriseTimings = resolveCapTimings{
		firstRequestDelay: 7 * time.Second,
		requestsInterval:  3 * time.Second,
		timeout:           180 * time.Second,
	}

	funCaptchaTimings = resolveCapTimings{
		firstRequestDelay: 9 * time.Second,
		requestsInterval:  1 * time.Second,
		timeout:           80 * time.Second,
	}

	hCaptchaTimings = resolveCapTimings{
		firstRequestDelay: 7 * time.Second,
		requestsInterval:  3 * time.Second,
		timeout:           180 * time.Second,
	}

	GeeTestTimings = resolveCapTimings{
		firstRequestDelay: 1 * time.Second,
		requestsInterval:  1 * time.Second,
		timeout:           80 * time.Second,
	}
)
