package command

import (
	"encoding/json"
	"fmt"
	"github.com/fzdwx/burst"
	"github.com/fzdwx/burst/client"
	"github.com/fzdwx/burst/pkg"
	"github.com/fzdwx/burst/pkg/model/req"
	"github.com/spf13/cast"
	"net/url"
	"strings"
)

type (
	addProxyCommand struct{}
)

func (a addProxyCommand) usage() {
	fmt.Println("  ap: add proxy ")
	fmt.Println("      format: ap [channelType]:[ip]:[port]")
	fmt.Println("      example: ap tcp::8888 tcp:192.168.1.1:9999 tcp:11.22.33.44:5555 ...")
}

func (a addProxyCommand) callUsage() {
	Dispatch("u ap", nil)
}

func (a addProxyCommand) run(s []string, c *client.Client) {
	if len(s) == 0 {
		errorMsg("proxy is empty")
		return
	}

	var infos []req.AddProxyInfo

	for _, line := range s {
		split := strings.Split(line, ":")
		if len(split) < 3 {
			errorMsg("proxy format error: " + line)
			a.callUsage()
			return
		}

		port, err := cast.ToIntE(strings.TrimSuffix(split[2], "\r"))
		if err != nil {
			errorMsg(fmt.Sprintf("port %s is not valid", split[2]))
			a.callUsage()
			return
		}

		var ip string
		if split[1] == burst.EmptyStr {
			ip = "localhost"
		} else {
			ip = split[1]
		}

		proxyInfo := req.AddProxyInfo{
			ChannelType: split[0],
			Ip:          ip,
			Port:        cast.ToInt(port),
		}
		infos = append(infos, proxyInfo)
	}

	proxyInfoReq := req.AddProxyInfoReq{Proxy: infos}
	err := proxyInfoReq.Check()
	if err != nil {
		errorMsg(err.Error())
		a.callUsage()
		return
	}

	u := url.URL{Path: AddProxy + c.Token(), Scheme: "http", Host: c.ServerAddr()}

	resp, err := PostJson(u, proxyInfoReq)
	if err != nil {
		errorMsg(err.Error())
		return
	}

	f, response := ShowResp(resp)

	var proxyInfos []pkg.ClientProxyInfo
	err = json.Unmarshal(response, &proxyInfos)
	if err != nil {
		f(string(response))
		return
	}

	for _, proxyInfo := range proxyInfos {
		f("add proxy:")
		f("    " + proxyInfo.String())
	}
}
