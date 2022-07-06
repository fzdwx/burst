package proxy

import (
	"github.com/fzdwx/burst"
	"github.com/fzdwx/burst/pkg/model"
	"github.com/fzdwx/burst/pkg/result"
	"github.com/fzdwx/burst/server/svc"
	"net/http"
)

func RemoveProxy(svcContext *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := burst.GetQuery("token", r)
		if token == burst.EmptyStr {
			result.HttpBadRequest(w, model.TokenIsRequired.Error())
			return
		}

		// todo remove proxy
		// 	1. close listener
		//  2. remove proxy from cache
		//  3. notify client remove proxy

	}
}
