package zendesk

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/senorprogrammer/wtf/wtf"
)

type Resource struct {
	Response interface{}
	Raw      string
}

var a = os.Getenv("ZENDESK_API")
var subdomain = os.Getenv("ZENDESK_SUBDOMAIN")
var baseURL = fmt.Sprintf("https://%v.zendesk.com/api/v2", subdomain)

func errHandler(err error) {
	if err != nil {
		log.Print(err)
	}
}

func buildUrl(baseURL string) string {
	ticketURL := baseURL + "/tickets.json?sort_by=status"
	return ticketURL
}

func api(key string, meth string, path string, params string) (*Resource, error) {

	trn := &http.Transport{}

	client := &http.Client{
		Transport: trn,
	}

	var URL = buildUrl(baseURL)

	req, err := http.NewRequest(meth, URL, bytes.NewBufferString(params))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")

	username := wtf.Config.UString("wtf.mods.zendesk.username")
	apiUser := fmt.Sprintf("%v/token", username)
	req.SetBasicAuth(apiUser, key)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return &Resource{Response: &resp, Raw: string(data)}, nil

}
