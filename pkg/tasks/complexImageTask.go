package tasks

type ComplexImageTask struct {
	Type         string    `json:"type"`
	Class        string    `json:"class"`
	ImageUrls    *[]string `json:"imageUrls"`
	ImagesBase64 *[]string `json:"imagesBase64"`
	WebsiteURL   *string   `json:"websiteURL"`
	userAgentAndCookies
}

type ComplexImageTaskSolution struct {
	Answer []bool `json:"answer"`
}
