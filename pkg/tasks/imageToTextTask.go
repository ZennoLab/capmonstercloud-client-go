package tasks

type ImageToTextTask struct {
	Type                 string  `json:"type"`
	Body                 string  `json:"body"`
	CapMonsterModule     *string `json:"CapMonsterModule,omitempty"`
	RecognizingThreshold *int    `json:"recognizingThreshold,omitempty"`
	Case                 *bool   `json:"Case,omitempty"`
	Numeric              *int    `json:"numeric,omitempty"`
	Math                 *bool   `json:"math,omitempty"`
}

func NewImageToTextTask(body string) ImageToTextTask {
	return ImageToTextTask{
		Type: "ImageToTextTask",
		Body: body,
	}
}

func (t ImageToTextTask) WithCapMonsterModule(capMonsterModule string) ImageToTextTask {
	t.CapMonsterModule = &capMonsterModule
	return t
}

func (t ImageToTextTask) WithRecognizingThreshold(recognizingThreshold int) ImageToTextTask {
	t.RecognizingThreshold = &recognizingThreshold
	return t
}

func (t ImageToTextTask) WithCase(taskCase bool) ImageToTextTask {
	t.Case = &taskCase
	return t
}

func (t ImageToTextTask) WithNumeric(numeric int) ImageToTextTask {
	t.Numeric = &numeric
	return t
}

func (t ImageToTextTask) WithMath(math bool) ImageToTextTask {
	t.Math = &math
	return t
}

type ImageToTextTaskResult struct {
	result
	Solution struct {
		Text    string            `json:"text"`
		Cookies map[string]string `json:"cookies"`
	} `json:"solution"`
}
