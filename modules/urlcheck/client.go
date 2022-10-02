package urlcheck

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/wtfutil/wtf/logger"
)

func GetClient() *http.Client {

	// Client
	client := &http.Client{
		Transport: http.DefaultTransport,
	}

	return client
}

func DoRequest(urlRequest string, timeout time.Duration, client *http.Client) (int, string) {
	// Context
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// Request
	req, err := http.NewRequest(http.MethodHead, urlRequest, nil)
	if err != nil {
		logger.Log(fmt.Sprintf("[urlcheck] ERROR %s: %s", urlRequest, err.Error()))
		return InvalidResultCode, "New Request Error"
	}
	req = req.WithContext(ctx)

	// Send the request
	res, err := client.Do(req)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			status := "Timeout"
			logger.Log(fmt.Sprintf("[urlcheck] %s: %s", urlRequest, status))
			return InvalidResultCode, status
		}
		logger.Log(fmt.Sprintf("[urlcheck] %s: %s", urlRequest, err.Error()))
		return InvalidResultCode, "Error"
	}

	defer res.Body.Close()

	return res.StatusCode, res.Status
}
