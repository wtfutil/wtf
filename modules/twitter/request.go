package twitter

import (
	"bytes"
	"net/http"
)

func Request(httpClient *http.Client, apiURL string) ([]byte, error) {
	resp, err := httpClient.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	data, err := ParseBody(resp)
	if err != nil {
		return nil, err
	}

	return data, err
}

func ParseBody(resp *http.Response) ([]byte, error) {
	var buffer bytes.Buffer
	_, err := buffer.ReadFrom(resp.Body)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
