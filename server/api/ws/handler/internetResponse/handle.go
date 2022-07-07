package internetResponse

import (
	"github.com/fzdwx/burst/pkg/logx"
	"github.com/fzdwx/burst/pkg/protocal"
	"github.com/fzdwx/burst/server/cache"
)

func Handle(internetResponse protocal.IntranetResponse) {
	container, b := cache.ServerContainer.Get(internetResponse.Token)
	if !b {
		logx.Error().Msg("server container not found on write to user")
		return
	}

	conn, b := container.GetUserConn(internetResponse.ConnId)
	if !b {
		logx.Debug().
			Bytes("bytes", internetResponse.Data).
			Str("str", string(internetResponse.Data)).
			Msg("user conn not found on write to user")
		return
	}

	conn.Write(internetResponse.Data)
}
