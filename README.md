## Usage
```go
package main

import (
	"fmt"

	"github.com/ZennoLab/capmonstercloud-client-go/pkg/client"
	"github.com/ZennoLab/capmonstercloud-client-go/pkg/tasks"
)

func main() {
	client := client.New("YourClientKey")

	//get balance
	if balance, err := client.GetBalance(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(balance)
	}

	// solve RecaptchaV2 (without proxy)
	recaptureV2Task := tasks.NewRecaptchaV2TaskProxyless(
		"https://lessons.zennolab.com/captchas/recaptcha/v2_simple.php?level=high",
		"6Lcg7CMUAAAAANphynKgn9YAgA4tQ2KI_iqRyTwd",
	)
	noCache := false
	if result, err := client.ResolveRecaptchaV2Proxyless(recaptureV2Task, noCache, nil); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	// solve ImageToTextTask with module
	imgToTextTask := tasks.NewImageToTextTask("BODY").WithCapMonsterModule(tasks.CapMonsterModuleGoogle)
	if result, err := client.ResolveImageToText(imgToTextTask, nil); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
```
