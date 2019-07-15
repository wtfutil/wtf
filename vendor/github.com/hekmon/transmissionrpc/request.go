package transmissionrpc

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"sync"
)

const csrfHeader = "X-Transmission-Session-Id"

type requestPayload struct {
	Method    string      `json:"method"`
	Arguments interface{} `json:"arguments,omitempty"`
	Tag       int         `json:"tag,omitempty"`
}

type answerPayload struct {
	Arguments interface{} `json:"arguments"`
	Result    string      `json:"result"`
	Tag       *int        `json:"tag"`
}

func (c *Client) rpcCall(method string, arguments interface{}, result interface{}) (err error) {
	return c.request(method, arguments, result, true)
}

func (c *Client) request(method string, arguments interface{}, result interface{}, retry bool) (err error) {
	// Let's avoid crashing
	if c.httpC == nil {
		err = errors.New("this controller is not initialized, please use the New() function")
		return
	}
	// Prepare the pipeline between payload generation and request
	pOut, pIn := io.Pipe()
	// Prepare the request
	var req *http.Request
	if req, err = http.NewRequest("POST", c.url, pOut); err != nil {
		err = fmt.Errorf("can't prepare request for '%s' method: %v", method, err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", c.userAgent)
	req.Header.Set(csrfHeader, c.getSessionID())
	req.SetBasicAuth(c.user, c.password)
	// Prepare the marshalling goroutine
	var tag int
	var encErr error
	var mg sync.WaitGroup
	mg.Add(1)
	go func() {
		tag = c.rnd.Int()
		encErr = json.NewEncoder(pIn).Encode(&requestPayload{
			Method:    method,
			Arguments: arguments,
			Tag:       tag,
		})
		pIn.Close()
		mg.Done()
	}()
	// Execute request
	var resp *http.Response
	if resp, err = c.httpC.Do(req); err != nil {
		mg.Wait()
		if encErr != nil {
			err = fmt.Errorf("request error: %v | json payload marshall error: %v", err, encErr)
		} else {
			err = fmt.Errorf("request error: %v", err)
		}
		return
	}
	defer resp.Body.Close()
	// Let's test the enc result, just in case
	mg.Wait()
	if encErr != nil {
		err = fmt.Errorf("request payload JSON marshalling failed: %v", encErr)
		return
	}
	// Is the CRSF token invalid ?
	if resp.StatusCode == http.StatusConflict {
		// Recover new token and save it
		c.updateSessionID(resp.Header.Get(csrfHeader))
		// Retry request if first try
		if retry {
			return c.request(method, arguments, result, false)
		}
		err = errors.New("CSRF token invalid 2 times in a row: stopping to avoid infinite loop")
		return
	}
	// Is request successful ?
	if resp.StatusCode != 200 {
		err = fmt.Errorf("HTTP error %d: %s", resp.StatusCode, http.StatusText(resp.StatusCode))
		return
	}
	// // Debug
	// {
	// 	var data []byte
	// 	data, err = ioutil.ReadAll(resp.Body)
	// 	fmt.Println(string(data))
	// 	return
	// }
	// Decode body
	answer := answerPayload{
		Arguments: result,
	}
	if err = json.NewDecoder(resp.Body).Decode(&answer); err != nil {
		err = fmt.Errorf("can't unmarshall request answer body: %v", err)
		return
	}
	// fmt.Println("DEBUG >", answer.Result)
	// Final checks
	if answer.Tag == nil {
		err = errors.New("http answer does not have a tag within it's payload")
		return
	}
	if *answer.Tag != tag {
		err = errors.New("http request tag and answer payload tag do not match")
		return
	}
	if answer.Result != "success" {
		err = fmt.Errorf("http request ok but payload does not indicate success: %s", answer.Result)
		return
	}
	// All good
	return
}

func (c *Client) getSessionID() string {
	defer c.sessionIDAccess.RUnlock()
	c.sessionIDAccess.RLock()
	return c.sessionID
}

func (c *Client) updateSessionID(newID string) {
	defer c.sessionIDAccess.Unlock()
	c.sessionIDAccess.Lock()
	c.sessionID = newID
}
