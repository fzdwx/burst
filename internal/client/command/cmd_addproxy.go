package command

import (
	"encoding/json"
	"fmt"
	"github.com/fzdwx/burst/internal"
	"github.com/fzdwx/burst/internal/client"
	"github.com/fzdwx/burst/internal/model/req"
	"github.com/knz/bubbline"
	"github.com/knz/bubbline/computil"
	"github.com/knz/bubbline/editline"
	"github.com/spf13/cast"
	"net/url"
	"strings"
)

type (
	addProxyCommand struct{}
)

func (a addProxyCommand) autocomplete() bubbline.AutoCompleteFn {
	return func(v [][]rune, line, col int) (msg string, comp editline.Completions) {
		// Detect the word under the cursor.
		word, wstart, wend := computil.FindWord(v, line, col)

		wordWithoutSpace := strings.TrimSpace(word)

		if len(wordWithoutSpace) == 0 {
			return "", editline.SimpleWordsCompletion(internal.ChannelTypes, "channelType", col, wstart, wend)
		}
		var channelTypeCandidates []string
		for _, name := range internal.ChannelTypes {
			if strings.HasPrefix(name, word) {
				channelTypeCandidates = append(channelTypeCandidates, fmt.Sprintf("%s:", name))
			}
		}
		// todo 返回完成后会默认加一个空格
		if len(channelTypeCandidates) != 0 {
			return "", editline.SimpleWordsCompletion(channelTypeCandidates, "channelType", col, wstart, wend)
		}

		return fmt.Sprintf("your input word is %s", word), nil
	}
}

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
		if split[1] == internal.EmptyStr {
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

	var proxyInfos []internal.ClientProxyInfo
	err = json.Unmarshal(response, &proxyInfos)
	if err != nil {
		f(string(response))
		return
	}

	for _, proxyInfo := range proxyInfos {
		f("add proxy:")
		f("    " + fmt.Sprintf("%s -> %s", proxyInfo.IntranetAddr, proxyInfo.Address(internal.GetCurrentIp())))
	}
}
