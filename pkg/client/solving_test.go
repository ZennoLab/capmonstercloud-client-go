package client

import (
	"errors"
	"os"
	"testing"

	"github.com/ZennoLab/capmonstercloud-client-go/pkg/tasks"
)

func TestSolveRecaptchaV2Proxyless(t *testing.T) {
	client := New(os.Getenv(testingKeyEnvVarName))
	task := tasks.NewRecaptchaV2TaskProxyless(
		"https://lessons.zennolab.com/captchas/recaptcha/v2_simple.php?level=high",
		"6Lcg7CMUAAAAANphynKgn9YAgA4tQ2KI_iqRyTwd",
	)
	result, err := client.SolveRecaptchaV2Proxyless(task, true, nil)

	switch {
	case err != nil:
		t.Error(err)
	case result.GRecaptchaResponse == "":
		t.Error("empty result")
	}
}

func TestSolveHCaptchaProxyless(t *testing.T) {
	client := New(os.Getenv(testingKeyEnvVarName))
	task := tasks.NewHCaptchaTaskProxyless("https://lessons.zennolab.com/captchas/hcaptcha/?level=easy",
		"472fc7af-86a4-4382-9a49-ca9090474471")
	result, err := client.SolveHCaptchaProxyless(task, true, nil)

	switch {
	case err != nil:
		t.Error(err)
	case result.GRecaptchaResponse == "":
		t.Error("empty result")
	case result.RespKey == "":
		t.Error("empty result")
	case result.UserAgent == "":
		t.Error("empty result")
	}
}

func TestIncorrectWebsiteUrl(t *testing.T) {
	wantErr := tasks.ErrInvalidWebsiteUrl

	client := New(os.Getenv(testingKeyEnvVarName))
	task := tasks.NewHCaptchaTaskProxyless("incorrect url",
		"6Lcg7CMUAAAAANphynKgn9YAgA4tQ2KI_iqRyTwd")
	_, gotErr := client.SolveHCaptchaProxyless(task, true, nil)

	if !errors.Is(gotErr, wantErr) {
		t.Errorf("want %q error, got %q error", wantErr, gotErr)
	}
}

func TestIncorrectProxyPort(t *testing.T) {
	ports := []int{65535 + 1, 65535 + 2, 65535 + 100}
	wantErr := tasks.ErrInvalidProxyPort

	client := New(os.Getenv(testingKeyEnvVarName))
	task := tasks.NewRecaptchaV2Task(
		"https://lessons.zennolab.com/captchas/recaptcha/v2_simple.php?level=high",
		"6Lcg7CMUAAAAANphynKgn9YAgA4tQ2KI_iqRyTwd",
		tasks.ProxyTypeHttp,
		"localhost",
		0,
	)
	for _, nextPort := range ports {
		task.ProxyPort = nextPort
		_, gotErr := client.SolveRecaptchaV2(task, true, nil)
		if !errors.Is(gotErr, wantErr) {
			t.Errorf("want %q error, got %q error", wantErr, gotErr)
		}
	}
}

func TestIncorrectMinScore(t *testing.T) {
	minScores := []float64{0.09, 0.901, 1.1}
	wantErr := tasks.ErrInvalidMinScore

	client := New(os.Getenv(testingKeyEnvVarName))
	task := tasks.NewRecaptchaV3TaskProxyless(
		"https://lessons.zennolab.com/captchas/recaptcha/v2_simple.php?level=high",
		"6Lcg7CMUAAAAANphynKgn9YAgA4tQ2KI_iqRyTwd",
	)
	for _, nextMinScore := range minScores {
		task.MinScore = &nextMinScore
		_, gotErr := client.SolveRecaptchaV3Proxyless(task, true, nil)
		if !errors.Is(gotErr, wantErr) {
			t.Errorf("want %q error, got %q error", wantErr, gotErr)
		}
	}
}

func TestIncorrectRecognizingThreshold(t *testing.T) {
	recognizingThresholds := []int{101, 102}
	wantErr := tasks.ErrInvalidRecognizingThreshold

	client := New(os.Getenv(testingKeyEnvVarName))
	task := tasks.NewImageToTextTask("body")

	for _, nextRecognizingThreshold := range recognizingThresholds {
		task.RecognizingThreshold = &nextRecognizingThreshold
		_, gotErr := client.SolveImageToText(task, nil)
		if !errors.Is(gotErr, wantErr) {
			t.Errorf("want %q error, got %q error", wantErr, gotErr)
		}
	}
}

func TestIncorrectWebsiteKey(t *testing.T) {
	wantErr := tasks.ErrInvalidWebSiteKey

	client := New(os.Getenv(testingKeyEnvVarName))
	task := tasks.NewRecaptchaV2TaskProxyless(
		"https://lessons.zennolab.com/captchas/recaptcha/v2_simple.php?level=high",
		"",
	)
	_, gotErr := client.SolveRecaptchaV2Proxyless(task, true, nil)
	if !errors.Is(gotErr, wantErr) {
		t.Errorf("want %q error, got %q error", wantErr, gotErr)
	}
}

func TestIncorrectGt(t *testing.T) {
	wantErr := tasks.ErrInvalidGt

	client := New(os.Getenv(testingKeyEnvVarName))
	task := tasks.NewGeeTestTaskProxyless(
		"https://lessons.zennolab.com/captchas/recaptcha/v2_simple.php?level=high",
		"",
	)
	//task = task.WithChallenge("d93591bdf7860e1e4ee2fca799911215")
	_, gotErr := client.SolveGeeTestProxyless(task, nil)
	if !errors.Is(gotErr, wantErr) {
		t.Errorf("want %q error, got %q error", wantErr, gotErr)
	}
}

func TestIncorrectGeeTestV3(t *testing.T) {
	wantErr := tasks.ErrChallenge

	client := New(os.Getenv(testingKeyEnvVarName))
	task := tasks.NewGeeTestTaskProxyless(
		"https://example.com/geetest.php",
		"81dc9bdb52d04dc20036dbd8313ed055",
	)
	_, gotErr := client.SolveGeeTestProxyless(task, nil)
	if !errors.Is(gotErr, wantErr) {
		t.Errorf("want %q error, got %q error", wantErr, gotErr)
	}
}

func TestGeeTestV4(t *testing.T) {
	client := New(os.Getenv(testingKeyEnvVarName))
	task := tasks.NewGeeTestTaskProxyless(
		"https://faucetpay.io/account/login",
		"eb8b0c2b27f3365b9244d9da81638c6",
	)
	task = task.WithVersion(4)

	task = task.WithInitParametres(struct {
		RiskType string `json:"riskType"`
	}{
		RiskType: "slide",
	})

	_, gotErr := client.SolveGeeTestProxyless(task, nil)
	if gotErr != nil {
		t.Errorf("got %q error", gotErr)
	}
}

func TestTurnstileProxless(t *testing.T) {
	client := New(os.Getenv(testingKeyEnvVarName))
	task := tasks.NewTurnstileTaskProxyless(
		"https://tsinvisble.zlsupport.com",
		"0x4AAAAAAABUY0VLtOUMAHxE",
	)

	_, gotErr := client.SolveTurnstileProxyless(task, nil)
	if gotErr != nil {
		t.Errorf("got %q error", gotErr)
	}
}

func TestIncorectTurnstileProxlessCloudflare(t *testing.T) {
	wantErr := tasks.ErrCloudflareTaskType
	client := New(os.Getenv(testingKeyEnvVarName))
	task := tasks.NewTurnstileTaskProxyless(
		"https://tsinvisble.zlsupport.com",
		"0x4AAAAAAABUY0VLtOUMAHxE",
	)

	task = task.WithCloudflareTaskType("cloud")

	_, gotErr := client.SolveTurnstileProxyless(task, nil)
	if !errors.Is(gotErr, wantErr) {
		t.Errorf("want %q error, got %q error", wantErr, gotErr)
	}
}

func TestIncorectTurnstileProxlessCloudflareUserAgent(t *testing.T) {
	wantErr := tasks.ErrUserAgentRequired
	client := New(os.Getenv(testingKeyEnvVarName))
	task := tasks.NewTurnstileTaskProxyless(
		"https://tsinvisble.zlsupport.com",
		"0x4AAAAAAABUY0VLtOUMAHxE",
	)

	task = task.WithCloudflareTaskType("cf_clearance")

	_, gotErr := client.SolveTurnstileProxyless(task, nil)
	if !errors.Is(gotErr, wantErr) {
		t.Errorf("want %q error, got %q error", wantErr, gotErr)
	}
}

func TestRecaptchaComplexImage(t *testing.T) {
	client := New(os.Getenv(testingKeyEnvVarName))
	task := tasks.NewRecaptchaComplexImageTask(tasks.MetadataRecaptcha{
		Grid: "3x3",
	})
	task = task.WithImagesUrls([]string{"https://i.postimg.cc/yYjg75Kv/payloadtraffic.jpg"})
	task = task.WithMetadataTask("Click on traffic lights")
	task = task.WithUserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36.")

	_, gotErr := client.SolveRecaptchaComplexImage(task, nil)
	if gotErr != nil {
		t.Errorf("got %q error", gotErr)
	}
}

func TestHCaptchaComplexImage(t *testing.T) {
	client := New(os.Getenv(testingKeyEnvVarName))
	task := tasks.NewHCaptchaComplexImageTask(tasks.MetadataHCaptcha{
		Task: "Please click each image containing a mountain",
	})
	task = task.WithImagesUrls([]string{"https://i.postimg.cc/kg71cbRt/image-1.jpg"})
	task = task.WithUserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36.")

	_, gotErr := client.SolveHCaptchaComplexImage(task, nil)
	if gotErr != nil {
		t.Errorf("got %q error", gotErr)
	}
}

func TestFuncaptchaComplexImage(t *testing.T) {
	client := New("f9dbccccad15946867313c3789d8b4d7")
	task := tasks.NewFuncaptchaComplexImageTask(tasks.MetadataFuncaptcha{
		Task: "Pick the image that is the correct way up",
	})
	task = task.WithImagesUrls([]string{"https://i.postimg.cc/s2ZDrHXy/fc1.jpg"})
	task = task.WithUserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36.")

	_, gotErr := client.SolveFuncaptchaComplexImage(task, nil)
	if gotErr != nil {
		t.Errorf("got %q error", gotErr)
	}
}
