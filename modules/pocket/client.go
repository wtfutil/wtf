package pocket

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Client pocket client Documention at https://getpocket.com/developer/docs/overview
type Client struct {
	consumerKey string
	accessToken *string
	baseURL     string
	redirectURL string
}

//NewClient returns a new PocketClient
func NewClient(consumerKey, redirectURL string) *Client {
	return &Client{
		consumerKey: consumerKey,
		redirectURL: redirectURL,
		baseURL:     "https://getpocket.com/v3",
	}

}

//Item represents link in pocket api
type Item struct {
	ItemID                 string `json:"item_id"`
	ResolvedID             string `json:"resolved_id"`
	GivenURL               string `json:"given_url"`
	GivenTitle             string `json:"given_title"`
	Favorite               string `json:"favorite"`
	Status                 string `json:"status"`
	TimeAdded              string `json:"time_added"`
	TimeUpdated            string `json:"time_updated"`
	TimeRead               string `json:"time_read"`
	TimeFavorited          string `json:"time_favorited"`
	SortID                 int    `json:"sort_id"`
	ResolvedTitle          string `json:"resolved_title"`
	ResolvedURL            string `json:"resolved_url"`
	Excerpt                string `json:"excerpt"`
	IsArticle              string `json:"is_article"`
	IsIndex                string `json:"is_index"`
	HasVideo               string `json:"has_video"`
	HasImage               string `json:"has_image"`
	WordCount              string `json:"word_count"`
	Lang                   string `json:"lang"`
	TimeToRead             int    `json:"time_to_read"`
	TopImageURL            string `json:"top_image_url"`
	ListenDurationEstimate int    `json:"listen_duration_estimate"`
}

//ItemLists represent list of links
type ItemLists struct {
	Status   int             `json:"status"`
	Complete int             `json:"complete"`
	List     map[string]Item `json:"list"`
	Since    int             `json:"since"`
}

type request struct {
	requestBody interface{}
	method      string
	result      interface{}
	headers     map[string]string
	url         string
}

func (client *Client) request(req request, result interface{}) error {
	jsonValues, err := json.Marshal(req.requestBody)
	if err != nil {
		return err
	}
	request, err := http.NewRequest(req.method, req.url, bytes.NewBuffer(jsonValues))
	if err != nil {
		return err
	}

	for key, value := range req.headers {
		request.Header.Add(key, value)
	}

	resp, err := http.DefaultClient.Do(request)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return err
	}
	if resp.StatusCode >= 400 {
		return fmt.Errorf(`server responded with [%d]:%s,url:%s`, resp.StatusCode, responseBody, req.url)
	}

	if err := json.Unmarshal(responseBody, &result); err != nil {
		return fmt.Errorf("Could not unmarshal url [%s] \n\t\tresponse [%s] request[%s] error:%w",
	req.url, responseBody, jsonValues, err)
	}

	return nil

}

type obtainRequestTokenRequest struct {
	ConsumerKey string `json:"consumer_key"`
	RedirectURI string `json:"redirect_uri"`
}

//ObtainRequestToken get request token to be used in the auth workflow
func (client *Client) ObtainRequestToken() (code string, err error) {
	url := fmt.Sprintf("%s/oauth/request", client.baseURL)
	requestData := obtainRequestTokenRequest{ConsumerKey: client.consumerKey, RedirectURI: client.redirectURL}

	var responseData map[string]string
	req := request{
		method:      "POST",
		url:         url,
		requestBody: requestData,
	}
	req.headers = map[string]string{
		"X-Accept":     "application/json",
		"Content-Type": "application/json",
	}
	err = client.request(req, &responseData)

	if err != nil {
		return code, err
	}

	return responseData["code"], nil

}

//CreateAuthLink create authorization link to redirect the user to
func (client *Client) CreateAuthLink(requestToken string) string {
	return fmt.Sprintf("https://getpocket.com/auth/authorize?request_token=%s&redirect_uri=%s", requestToken, client.redirectURL)
}

type accessTokenRequest struct {
	ConsumerKey string `json:"consumer_key"`
	RequestCode string `json:"code"`
}

// accessTokenResponse represents
type accessTokenResponse struct {
	AccessToken string `json:"access_token"`
}

//GetAccessToken exchange request token for accesstoken
func (client *Client) GetAccessToken(requestToken string) (accessToken string, err error) {
	url := fmt.Sprintf("%s/oauth/authorize", client.baseURL)
	requestData := accessTokenRequest{
		ConsumerKey: client.consumerKey,
		RequestCode: requestToken,
	}
	req := request{
		method:      "POST",
		url:         url,
		requestBody: requestData,
	}
	req.headers = map[string]string{
		"X-Accept":     "application/json",
		"Content-Type": "application/json",
	}

	var response accessTokenResponse
	err = client.request(req, &response)
	if err != nil {
		return "", err
	}
	return response.AccessToken, nil

}

/*LinkState  represents links states to be retrived
According to the api https://getpocket.com/developer/docs/v3/retrieve
there are 3 states:
	1-archive
	2-unread
	3-all
however archive does not really well work and returns links that are in the
unread list
buy inspecting getpocket I found out that there is an undocumanted read state
*/
type LinkState string

const (
	//Read links that has been read (undocumanted)
	Read LinkState = "read"
	//Unread links has not been read
	Unread LinkState = "unread"
)

// GetLinks retrive links of a given states https://getpocket.com/developer/docs/v3/retrieve
func (client *Client) GetLinks(state LinkState) (response ItemLists, err error) {
	url := fmt.Sprintf("%s/get?sort=newest&state=%s&consumer_key=%s&access_token=%s", client.baseURL, state, client.consumerKey, *client.accessToken)
	req := request{
		method: "GET",
		url:    url,
	}
	req.headers = map[string]string{
		"X-Accept":     "application/json",
		"Content-Type": "application/json",
	}

	err = client.request(req, &response)
	return response, err
}

//Action represents a mutation to link
type Action string

const (
	//Archive to put the link in the archived list (read list)
	Archive Action = "archive"
	//ReAdd to put the link back in the to reed list
	ReAdd Action = "readd"
)

type actionParams struct {
	Action Action `json:"action"`
	ItemID string `json:"item_id"`
}

//ModifyLink change the state of the link
func (client *Client) ModifyLink(action Action, itemID string) (ok bool, err error) {

	actions := []actionParams{
		{
			Action: action,
			ItemID: itemID,
		},
	}

	urlActionParm, err := json.Marshal(actions)
	if err != nil {
		return false, err
	}
	url := fmt.Sprintf("%s/send?consumer_key=%s&access_token=%s&actions=%s", client.baseURL, client.consumerKey, *client.accessToken, urlActionParm)

	req := request{
		method: "GET",
		url:    url,
	}

	err = client.request(req, nil)

	if err != nil {
		return false, err
	}

	return true, nil

}
