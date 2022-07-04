package req

import (
	"fmt"
	"github.com/fzdwx/burst"
	"github.com/fzdwx/burst/pkg"
	"github.com/fzdwx/burst/pkg/model"
	"net"
)

type (
	AddProxyInfoReq struct {
		Proxy []AddProxyInfo `json:"proxy"`
	}

	AddProxyInfo struct {
		Ip          string `json:"ip"`
		Port        int    `json:"port"`
		ChannelType string `json:"channelType,default=tcp"`
	}
)

func (i AddProxyInfo) Addr() string {
	return burst.FormatAddr(i.Ip, i.Port)
}

func (i AddProxyInfo) String() string {
	return fmt.Sprintf("%s:%s", i.ChannelType, i.Addr())
}

func (i AddProxyInfo) ToCache() *pkg.ServerProxyInfo {
	return &pkg.ServerProxyInfo{
		Ip:          i.Ip,
		Port:        i.Port,
		ChannelType: i.ChannelType,
		Addr:        i.Addr(),
	}
}

func (r AddProxyInfoReq) Check() error {
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
