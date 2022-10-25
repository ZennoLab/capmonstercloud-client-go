package tasks

type imageToTextTask struct {
	Type                 string  `json:"type"`
	Body                 string  `json:"body"`
	CapMonsterModule     *string `json:"CapMonsterModule,omitempty"`
	RecognizingThreshold *int    `json:"recognizingThreshold,omitempty"`
	Case                 *bool   `json:"Case,omitempty"`
	Numeric              *int    `json:"numeric,omitempty"`
	Math                 *bool   `json:"math,omitempty"`
}

func NewImageToTextTask(body string) imageToTextTask {
	return imageToTextTask{Type: "ImageToTextTask", Body: body}
}

func (t imageToTextTask) WithCapMonsterModule(capMonsterModule string) imageToTextTask {
	t.CapMonsterModule = &capMonsterModule
	return t
}

func (t imageToTextTask) WithRecognizingThreshold(recognizingThreshold int) imageToTextTask {
	t.RecognizingThreshold = &recognizingThreshold
	return t
}

func (t imageToTextTask) WithCase(taskCase bool) imageToTextTask {
	t.Case = &taskCase
	return t
}

func (t imageToTextTask) WithNumeric(numeric int) imageToTextTask {
	t.Numeric = &numeric
	return t
}

func (t imageToTextTask) WithMath(math bool) imageToTextTask {
	t.Math = &math
	return t
}
