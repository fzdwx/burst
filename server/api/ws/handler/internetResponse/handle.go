package internetResponse

import (
	"github.com/fzdwx/burst/pkg/logx"
	"github.com/fzdwx/burst/pkg/protocal"
	"github.com/fzdwx/burst/server/cache"
)

func Handle(internetResponse protocal.InternetResponse) {
	container, b := cache.ServerContainer.Get(internetResponse.Token)
	if !b {
		logx.Error().Msg("server container not found on write to user")
		return
	}

	conn, b := container.GetUserConn(internetResponse.ConnId)
	if !b {
		logx.Error().Msg("user conn not found on write to user")
		return
	}

	conn.Write(internetResponse.Data)
}
