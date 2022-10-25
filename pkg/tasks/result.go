package tasks

type result struct {
	ErrorId   int    `json:"errorId"`
	ErrorCode string `json:"errorCode"`
	Status    string `json:"status"`
}
