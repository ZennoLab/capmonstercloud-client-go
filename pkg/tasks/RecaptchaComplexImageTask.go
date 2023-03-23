package tasks

type MetadataRecaptcha struct {
	Task           *string `json:"Task"`
	Grid           string  `json:"Grid"`
	TaskDefinition *string `json:"TaskDefinition"`
}

type RecaptchaComplexImageTask struct {
	ComplexImageTask
	MetadataRecaptcha `json:"metadata"`
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
	if len(*t.MetadataRecaptcha.Task) < 1 && len(*t.MetadataRecaptcha.TaskDefinition) < 1 {
		return ErrMetaData
	}
	return nil
}
