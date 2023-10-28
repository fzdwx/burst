package user

import (
	"github.com/fzdwx/burst/internal/cache"
	"github.com/fzdwx/burst/internal/result"
	"github.com/fzdwx/burst/internal/server/svc"
	"github.com/rs/xid"
	"net/http"
)

// Auth The current role is to generate tokens
func Auth(*svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := xid.New().String()

		cache.ProxyInfoContainer.Add(token)
		result.HttpOk(w, token)
	}
}
