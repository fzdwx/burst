package command

import (
	"fmt"
	"github.com/fzdwx/burst/client"
	"github.com/fzdwx/burst/pkg/model/req"
	"github.com/spf13/cast"
	"net/url"
	"strings"
)

func addProxy(s []string, c *client.Client) {
	if len(s) == 0 {
		errorMsg("proxy is empty")
		return
	}

	var infos []req.AddProxyInfo

	for _, line := range s {
		split := strings.Split(line, ":")
		port, err := cast.ToIntE(strings.TrimSuffix(split[2], "\r"))
		if err != nil {
			errorMsg(fmt.Sprintf("port %s is not valid", split[2]))
			return
		}

		proxyInfo := req.AddProxyInfo{
			ChannelType: split[0],
			Ip:          split[1],
			Port:        cast.ToInt(port),
		}
		infos = append(infos, proxyInfo)
	}

	proxyInfoReq := req.AddProxyInfoReq{Proxy: infos}
	err := proxyInfoReq.Check()
	if err != nil {
		errorMsg(err.Error())
		return
	}

	u := url.URL{Path: AddProxy + c.Token(), Scheme: "http", Host: c.ServerAddr()}

	resp, err := PostJson(u, proxyInfoReq)
	if err != nil {
		errorMsg(err.Error())
		return
	}

	f, response := ShowResp(resp)

	f(response)
}
