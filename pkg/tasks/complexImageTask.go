package tasks

type ComplexImageTask struct {
	Type         string    `json:"type"`
	Class        string    `json:"class"`
	ImageUrls    *[]string `json:"imageUrls"`
	ImagesBase64 *[]string `json:"imagesBase64"`
	WebsiteURL   *string   `json:"websiteURL"`
	userAgentAndCookies
}

type MetadataHCaptcha struct {
	Task string `json:"Task"`
}

type MetadataRecaptcha struct {
	Task           *string `json:"Task"`
	Grid           string  `json:"Grid"`
	TaskDefinition *string `json:"TaskDefinition"`
}

type HCaptchaComplexImageTask struct {
	ComplexImageTask
	MetadataHCaptcha `json:"metadata"`
}

type RecaptchaComplexImageTask struct {
	ComplexImageTask
	MetadataRecaptcha `json:"metadata"`
}

func NewHCaptchaComplexImageTask(metadata MetadataHCaptcha) HCaptchaComplexImageTask {
	return HCaptchaComplexImageTask{
		ComplexImageTask: ComplexImageTask{
			Type:  "ComplexImageTask",
			Class: "hcaptcha",
		},
		MetadataHCaptcha: metadata,
	}
}

func (t HCaptchaComplexImageTask) WithImagesUrls(imageUrls []string) HCaptchaComplexImageTask {
	t.ComplexImageTask.ImageUrls = &imageUrls
	return t
}

func (t HCaptchaComplexImageTask) WithImagesBase64(imagesBase64 []string) HCaptchaComplexImageTask {
	t.ComplexImageTask.ImagesBase64 = &imagesBase64
	return t
}

func (t HCaptchaComplexImageTask) WithWebsiteURL(websiteURL string) HCaptchaComplexImageTask {
	t.ComplexImageTask.WebsiteURL = &websiteURL
	return t
}

func (t HCaptchaComplexImageTask) WithUserAgent(userAgent string) HCaptchaComplexImageTask {
	t.ComplexImageTask.UserAgent = &userAgent
	return t
}

func (t HCaptchaComplexImageTask) Validate() error {
	if len(t.MetadataHCaptcha.Task) < 1 {
		return ErrMetaData
	}
	return nil
}

func NewRecaptchaComplexImageTask(metadata MetadataRecaptcha) RecaptchaComplexImageTask {
	return RecaptchaComplexImageTask{
		ComplexImageTask: ComplexImageTask{
			Type:  "ComplexImageTask",
			Class: "recaptcha",
		},
		MetadataRecaptcha: metadata,
	}
}

func (t RecaptchaComplexImageTask) WithImagesUrls(imageUrls []string) RecaptchaComplexImageTask {
	t.ComplexImageTask.ImageUrls = &imageUrls
	return t
}

func (t RecaptchaComplexImageTask) WithImagesBase64(imagesBase64 []string) RecaptchaComplexImageTask {
	t.ComplexImageTask.ImagesBase64 = &imagesBase64
	return t
}

func (t RecaptchaComplexImageTask) WithWebsiteURL(websiteURL string) RecaptchaComplexImageTask {
	t.ComplexImageTask.WebsiteURL = &websiteURL
	return t
}

func (t RecaptchaComplexImageTask) WithMetadataTask(task string) RecaptchaComplexImageTask {
	t.MetadataRecaptcha.Task = &task
	return t
}

func (t RecaptchaComplexImageTask) WithMetadataTaskDefinition(taskDefinition string) RecaptchaComplexImageTask {
	t.MetadataRecaptcha.TaskDefinition = &taskDefinition
	return t
}

func (t RecaptchaComplexImageTask) WithUserAgent(userAgent string) RecaptchaComplexImageTask {
	t.ComplexImageTask.UserAgent = &userAgent
	return t
}

func (t RecaptchaComplexImageTask) Validate() error {
	if len(t.MetadataRecaptcha.Grid) < 1 {
		return ErrMetaData
	}
	return nil
}

type ComplexImageTaskSolution struct {
	Answer []bool `json:"answer"`
}
