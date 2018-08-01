package twitter

import (
	"bytes"
	"fmt"
	"net/http"
)

func Request(bearerToken string, apiURL string) ([]byte, error) {
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return nil, err
	}

	// Expected authorization format for single-application twitter dev accounts
	req.Header.Add("Authorization",
		fmt.Sprintf("Bearer %s", bearerToken))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

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
