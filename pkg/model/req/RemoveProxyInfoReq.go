package req

import (
	"fmt"
	"github.com/fzdwx/burst"
	"github.com/fzdwx/burst/pkg"
	"github.com/fzdwx/burst/pkg/model"
	"net"
)

type (
	RemoveProxyInfoReq struct {
		Proxy []RemoveProxyInfo `json:"proxy"`
	}

	RemoveProxyInfo struct {
		Ip          string `json:"ip"`
		Port        int    `json:"port"`
		ChannelType string `json:"channelType,default=tcp"`
	}
)

func (i RemoveProxyInfo) Addr() string {
	return burst.FormatAddr(i.Ip, i.Port)
}

func (i RemoveProxyInfo) String() string {
	return fmt.Sprintf("%s:%s", i.ChannelType, i.Addr())
}

func (r RemoveProxyInfoReq) Check() error {
	if len(r.Proxy) == 0 {
		return model.ProxyIsRequired
	}

	for _, info := range r.Proxy {
		if info.Ip == burst.EmptyStr {
			return model.IpIsNotValid
		}

		if info.Ip != "localhost" {
			ip := net.ParseIP(info.Ip)
			if ip == nil {
				return model.IpIsBlank
			}
		}

		if info.Port <= 0 || info.Port > 65535 {
			return model.PortIsNotValid
		}

		if !pkg.CheckChannelType(info.ChannelType) {
			return pkg.ErrChannelTypeNotValid
		}
	}
	return nil
}
