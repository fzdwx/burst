package proxy

import (
	"github.com/fzdwx/burst"
	"github.com/fzdwx/burst/pkg"
	"github.com/fzdwx/burst/pkg/wsx"
	"io"
)

type (
	Container struct {
		*wsx.Wsx
		Token  string
		Closer []io.Closer
	}
)

func NewContainer(ws *wsx.Wsx, token string) *Container {
	return &Container{Wsx: ws, Token: token, Closer: []io.Closer{}}
}

// Lunch Start the local service and then generate the format of the proxy information required by the client
//
func (c Container) Lunch(infos []*pkg.ServerProxyInfo) error {
	var clientInfos []*pkg.ClientProxyInfo
	for _, info := range infos {
		var clientInfo *pkg.ClientProxyInfo
		var err error
		switch info.ChannelType {
		case pkg.TCP:
			err, clientInfo = c.handleTCP(info)
		case pkg.HTTP:
			err, clientInfo = c.handlerHttp(info)
		case pkg.UDP:
			err, clientInfo = c.handleUdp(info)
		}

		if err != nil {
			return err
		}

		if clientInfo == nil {
			return burst.NewError("unSupport channelType %s", pkg.UDP)
		}
		clientInfos = append(clientInfos, clientInfo)
	}
	return nil
}

// Close the local service
func (c Container) Close() {
	for _, c := range c.Closer {
		c.Close()
	}
}
