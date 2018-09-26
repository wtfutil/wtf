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

func apiKey() string {
	return wtf.Config.UString(
		"wtf.mods.zendesk.apiKey",
		os.Getenv("ZENDESK_API"),
	)
}

func subdomain() string {
	return wtf.Config.UString(
		"wtf.mods.zendesk.subdomain",
		os.Getenv("ZENDESK_SUBDOMAIN"),
	)
}

func errHandler(err error) {
	if err != nil {
		log.Print(err)
	}
}

func api(key string, meth string, path string, params string) (*Resource, error) {
	trn := &http.Transport{}

	client := &http.Client{
		Transport: trn,
	}

	baseURL := fmt.Sprintf("https://%v.zendesk.com/api/v2", subdomain())
	URL := baseURL + "/tickets.json?sort_by=status"

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
