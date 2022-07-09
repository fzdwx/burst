package userConnect

import (
	"github.com/fzdwx/burst/client"
	"github.com/fzdwx/burst/pkg/logx"
	"github.com/fzdwx/burst/pkg/protocal"
	"net"
)

// Handle userConnect
//
// dail internet service
func Handle(c *client.Client, userConnect protocal.UserConnect) {
	proxy, ok := c.GetProxy(userConnect.Key)
	if !ok {
		logx.Error().Str("key", userConnect.Key).Msg("handle user request, proxy not found")
		// todo write err to server,close user client
		return
	}

	conn, err := net.Dial(proxy.ChannelType, proxy.IntranetAddr)
	if err != nil {
		logx.Error().Err(err).Msg("handle user request, dial internet service")
		// todo write err to server,close user client
		return
	}

	interNet := client.NewInternetService(conn, userConnect.ConnId, proxy)
	c.AddInterNetService(interNet)

	clean := func() {
		logx.Debug().Msg("clean user connect")
		c.RemoveInterNetService(interNet)
	}

	go interNet.StartRead(c, clean)
	go interNet.StartWrite(clean)
}
