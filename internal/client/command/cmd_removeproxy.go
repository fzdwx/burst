package command

import (
	"encoding/json"
	"fmt"
	"github.com/fzdwx/burst/internal/client"
	"github.com/fzdwx/burst/internal/model/req"
	"github.com/spf13/cast"
	"net/url"
	"strings"
)

type (
	removeProxyCommand struct {
	}
)

func (r removeProxyCommand) usage() {
	fmt.Println("  rp: remove proxy ")
	fmt.Println("      format: rp [channelType]:[ip]:[port]")
	fmt.Println("      example: rp tcp::8888 tcp::63342 tcp:192.168.1.1:8080 ...")
}

func (r removeProxyCommand) callUsage() {
	Dispatch("u rp", nil)
}

func (r removeProxyCommand) run(s []string, c *client.Client) {
	if len(s) == 0 {
		errorMsg("proxy is empty")
		return
	}

	var infos []req.RemoveProxyInfo

	for _, line := range s {
		split := strings.Split(line, ":")
		if len(split) < 3 {
			errorMsg("proxy format error: " + line)
			r.callUsage()
			return
		}

		port, err := cast.ToIntE(strings.TrimSuffix(split[2], "\r"))
		if err != nil {
			errorMsg(fmt.Sprintf("port %s is not valid", split[2]))
			r.callUsage()
			return
		}

		var ip string
		if split[1] == internal.EmptyStr {
			ip = "localhost"
		} else {
			ip = split[1]
		}

		proxyInfo := req.RemoveProxyInfo{
			ChannelType: split[0],
			Ip:          ip,
			Port:        cast.ToInt(port),
		}
		infos = append(infos, proxyInfo)
	}

	proxyInfoReq := req.RemoveProxyInfoReq{Proxy: infos}
	err := proxyInfoReq.Check()
	if err != nil {
		errorMsg(err.Error())
		r.callUsage()
		return
	}

	u := url.URL{Path: RemoveProxy + c.Token(), Scheme: "http", Host: c.ServerAddr()}

	resp, err := PostJson(u, proxyInfoReq)
	if err != nil {
		errorMsg(err.Error())
		return
	}

	f, response := ShowResp(resp)

	var proxyInfos []internal.ClientProxyInfo
	err = json.Unmarshal(response, &proxyInfos)
	if err != nil {
		f(string(response))
		return
	}

	for _, proxyInfo := range proxyInfos {
		f("remove proxy:")
		f("    " + fmt.Sprintf("%s -\\> %s", proxyInfo.IntranetAddr, proxyInfo.Address(internal.GetCurrentIp())))
	}
}
