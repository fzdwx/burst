package client

import (
	"github.com/fzdwx/burst/common/wsx"
	"github.com/zeromicro/go-zero/core/logx"
	"io/ioutil"
	"net/url"
)

func Connect(url url.URL, token string) {
	hub := wsx.NewHub(nil)
	client, response, err := wsx.Connect(url, token, hub)
	if err != nil {
		if response != nil {
			defer response.Body.Close()
			message, err := ioutil.ReadAll(response.Body)
			logx.Severe("connect error: %v", message)
			logx.Must(err)
		} else {
			logx.Must(err)
		}
	}

	logx.Infof("connect success: %v", client.RemoteAddr())

	go hub.React()

	client.OnBinary(func(data []byte) {
		logx.Infof("binary: %v", data)
	})

	go client.React()
}
