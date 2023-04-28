package tasks

type MetadataFuncaptcha struct {
	Task string `json:"Task"`
}

type FuncaptchaComplexImageTask struct {
	ComplexImageTask
	MetadataFuncaptcha `json:"metadata"`
}

func NewFuncaptchaComplexImageTask(metadata MetadataFuncaptcha) FuncaptchaComplexImageTask {
	return FuncaptchaComplexImageTask{
		ComplexImageTask: ComplexImageTask{
			Type:  "ComplexImageTask",
			Class: "funcaptcha",
		},
		MetadataFuncaptcha: metadata,
	}
}

func (t FuncaptchaComplexImageTask) WithImagesUrls(imageUrls []string) FuncaptchaComplexImageTask {
	t.ComplexImageTask.ImageUrls = &imageUrls
	return t
}

func (t FuncaptchaComplexImageTask) WithImagesBase64(imagesBase64 []string) FuncaptchaComplexImageTask {
	t.ComplexImageTask.ImagesBase64 = &imagesBase64
	return t
}

func (t FuncaptchaComplexImageTask) WithWebsiteURL(websiteURL string) FuncaptchaComplexImageTask {
	t.ComplexImageTask.WebsiteURL = &websiteURL
	return t
}

func (t FuncaptchaComplexImageTask) WithMetadataTask(task string) FuncaptchaComplexImageTask {
	t.MetadataFuncaptcha.Task = task
	return t
}

func (t FuncaptchaComplexImageTask) WithUserAgent(userAgent string) FuncaptchaComplexImageTask {
	t.ComplexImageTask.UserAgent = &userAgent
	return t
}

func (t FuncaptchaComplexImageTask) Validate() error {
	if len(t.MetadataFuncaptcha.Task) < 1 {
		return ErrMetaData
	}
	return nil
}
