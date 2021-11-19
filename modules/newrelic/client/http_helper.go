package newrelic

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func (c *Client) doGet(path string, params fmt.Stringer, out interface{}) error {
	var s string
	if params != nil {
		s = params.String()
	}
	r := strings.NewReader(s)
	req, err := http.NewRequest("GET", c.url.String()+path, r)
	if err != nil {
		return err
	}
	req.Header.Add("X-Api-Key", c.apiKey)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	return c.doRequest(req, out)
}

func (c *Client) doRequest(req *http.Request, out interface{}) error {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer func() { _ = resp.Body.Close() }()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("newrelic http error (%s): %s", resp.Status, b)
	}
	if len(b) == 0 {
		b = []byte{'{', '}'}
	}
	err = json.Unmarshal(b, &out)
	if err != nil {
		return err
	}
	return nil
}

func encodeGetParams(params map[string]interface{}) string {
	s := url.Values{}
	for k, v := range params {
		switch val := v.(type) {
		case string:
			if val != "" {
				s.Add(k, val)
			}
		case int:
			if val != 0 {
				s.Add(k, strconv.Itoa(val))
			}
		case []string:
			if len(val) != 0 {
				s.Add(k, strings.Join(val, ","))
			}
		case []int:
			arr := []string{}
			for _, v := range val {
				arr = append(arr, strconv.Itoa(v))
			}
			if len(arr) != 0 {
				s.Add(k, strings.Join(arr, ","))
			}
		case time.Time:
			if !val.IsZero() {
				s.Add(k, val.String())
			}
		case Array:
			for _, v := range val.arr {
				s.Add(k, v)
			}
		case bool:
			s.Add(k, "true")
		default:
			s.Add(k, fmt.Sprintf("%v", v))
		}
	}
	return s.Encode()
}
