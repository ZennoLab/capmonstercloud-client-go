package tasks

type MetadataHCaptcha struct {
	Task string `json:"Task"`
}

type HCaptchaComplexImageTask struct {
	ComplexImageTask
	MetadataHCaptcha `json:"metadata"`
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

func (t HCaptchaComplexImageTask) WithMetadataTask(task string) HCaptchaComplexImageTask {
	t.MetadataHCaptcha.Task = task
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
