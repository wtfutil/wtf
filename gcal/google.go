package gcal

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/senorprogrammer/wtf/logger"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// OAuthClientConfig describes Google oauth2 client configuration
type OAuthClientConfig struct {
	SecretFile string
	TokenFile  string
	Scope      string
}

// OAuthUIHandler returns oauth2 authorization code, implements user interaction
type OAuthUIHandler = func(string, chan string)

func BuildOAuthClient(config OAuthClientConfig, handler OAuthUIHandler) (*http.Client, error) {
	ctx := context.Background()

	b, err := ioutil.ReadFile(config.SecretFile)
	if err != nil {
		logger.Log(fmt.Sprintf("Invalid secret file provided (%s): %v", config.SecretFile, err))
		return nil, err
	}

	oauthConfig, err := google.ConfigFromJSON(b, config.Scope)
	if err != nil {
		logger.Log(fmt.Sprintf("Secret file is not readable as JSON (%s): %v", config.SecretFile, err))
		return nil, err
	}

	tok, err := tokenFromFile(config.TokenFile)
	if err != nil {
		result := make(chan string)
		go handler(oauthConfig.AuthCodeURL("state-token", oauth2.AccessTypeOffline), result)
		code, ok := <-result
		if !ok {
			// Cancelled
			return nil, err
		}
		tok, err = oauthConfig.Exchange(oauth2.NoContext, code)
		if err != nil {
			logger.Log(fmt.Sprintf("Invalid exchange code provided (%s): %v", code, err))
			return nil, err
		}
		err = saveToken(config.TokenFile, tok)
		if err != nil {
			logger.Log(fmt.Sprintf("Unable to store oauth token (%s): %v", config.TokenFile, err))
		}

	}
	return oauthConfig.Client(ctx, tok), nil
}

// CreateCodeInputDialog shows oauth2 code input dialog
func CreateCodeInputDialog(title string, widget *Widget) OAuthUIHandler {
	return func(url string, result chan string) {
		readInput := func() {
			fmt.Printf("\r\nOAuth authorization [%s]\r\n", title)
			fmt.Printf("Please open the following URL in your browser and paste authorization code below:\r\n")
			fmt.Printf("%s\r\nAuthorization code: ", url)
			var code string
			if _, err := fmt.Scan(&code); err != nil {
				logger.Log(fmt.Sprintf("Unable to read authorization code: %v", err))
				close(result)
			} else {
				result <- code
			}
		}
		widget.app.Suspend(readInput)
	}
}

// tokenFromFile retrieves a Token from a given file path.
// It returns the retrieved Token and any read error encountered.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	t := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(t)
	defer f.Close()
	return t, err
}

// saveToken uses a file path to create a file and store the
// token in it.
func saveToken(file string, token *oauth2.Token) error {
	// fmt.Printf("Saving credential file to: %s\n", file)
	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
	return nil
}
