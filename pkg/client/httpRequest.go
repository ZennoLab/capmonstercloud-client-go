package client

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func (c *capmonsterClient) invokeRequest(body []byte, url string) ([]byte, error) {
	bodyReader := bytes.NewReader(body)
	req, err := http.NewRequest(http.MethodPost, url, bodyReader)
	if err != nil {
		return []byte{}, fmt.Errorf("create http request: %w", err)
	}
	req.Header = reqHeaders

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return []byte{}, fmt.Errorf("http request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusServiceUnavailable {
		return []byte{}, errServiceUnavailable
	}
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return []byte{}, fmt.Errorf("responce status code: %v", resp.StatusCode)
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, fmt.Errorf("read response body: %w", err)
	}
	return respBody, nil
}
