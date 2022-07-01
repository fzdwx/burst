package ping

import (
	"github.com/fzdwx/burst/pkg/logx"
	"github.com/fzdwx/burst/pkg/result"
	"github.com/fzdwx/burst/server/svc"
	"net/http"
)

// Ping test function
func Ping(svcContext *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logx.Info().Msgf("hello world")
		result.HttpOk(w, "pong")
	}
}
