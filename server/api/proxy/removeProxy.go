package proxy

import (
	"github.com/fzdwx/burst/server/svc"
	"net/http"
)

func RemoveProxy(svcContext *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// todo remove proxy
		// 	1. close listener
		//  2. remove proxy from cache
		//  3. notify client remove proxy

	}
}
