package handler

import (
	"github.com/fzdwx/burst/client"
	"github.com/fzdwx/burst/client/handler/addProxy"
	"github.com/fzdwx/burst/client/handler/removeProxy"
	"github.com/fzdwx/burst/client/handler/userConnect"
	"github.com/fzdwx/burst/client/handler/userRequest"
	"github.com/fzdwx/burst/pkg/logx"
	"github.com/fzdwx/burst/pkg/protocal"
)

func Dispatch(c *client.Client) func(bytes []byte) {
	return func(bytes []byte) {
		burst, err := protocal.Decode(bytes)
		if err != nil {
			logx.Err(err).Msg("decode burst")
			return
		}

		switch burst.Type {
		case protocal.AddProxyType:
			addProxy.Handle(c, burst.AddProxy)
		case protocal.RemoveProxyType:
			removeProxy.Handle(c, burst.RemoveProxy)
		case protocal.UserConnectType:
			userConnect.Handle(c, burst.UserConnect)
		case protocal.UserRequestType:
			userRequest.Handle(c, burst.UserRequest)
		}
	}
}
