package command

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	AddProxy    = "/proxy/add/"
	RemoveProxy = "/proxy/remove/"
)

func PostJson(url url.URL, body interface{}) (*http.Response, error) {
	b, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	return http.Post(url.String(), "application/json", bytes.NewBuffer(b))
}

func ShowResp(resp *http.Response) (func(msg string), []byte) {
	if resp.Body != nil {
		defer resp.Body.Close()
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		errorMsg(err.Error())
	}

	switch resp.StatusCode {
	case 200:
		return infoMsg, b
	default:
		return errorMsg, b
	}
}
