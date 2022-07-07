package proxy

import (
	"github.com/fzdwx/burst"
	"github.com/fzdwx/burst/pkg"
	"github.com/fzdwx/burst/pkg/model"
	"github.com/fzdwx/burst/pkg/model/req"
	"github.com/fzdwx/burst/pkg/result"
	"github.com/fzdwx/burst/server/cache"
	"github.com/fzdwx/burst/server/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func RemoveProxy(svcContext *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token, proxyInfoReq, err := removeProxyPreCheck(w, r)
		if err != nil {
			result.HttpBadRequest(w, err.Error())
			return
		}

		info, ok := cache.ProxyInfoContainer.Get(token)
		if !ok {
			result.HttpBadRequest(w, model.ProxyInfoNotFound.Error())
		}

		server, b := cache.ServerContainer.Get(token)
		if !b {
			result.HttpBadRequest(w, model.ClientNotFound.Error())
			return
		}

		if server.Closed() {
			server.Close()
			result.HttpBadRequest(w, model.ServerClosed.Error())
		}

		// check if proxy is existed
		var proxyInfos []*pkg.ServerProxyInfo
		for _, proxyInfo := range proxyInfoReq.Proxy {
			if info.Has(proxyInfo.Addr()) {
				proxyInfos = append(proxyInfos, proxyInfo.ToCache())
			}
		}
		if len(proxyInfos) == 0 {
			result.HttpOk(w, "the proxy is not found")
			return
		}

		// todo remove proxy
		// 	1. close listener
		//  2. remove proxy from cache
		//  3. notify client remove proxy

	}
}

func removeProxyPreCheck(w http.ResponseWriter, r *http.Request) (string, *req.RemoveProxyInfoReq, error) {
	token := burst.GetQuery("token", r)
	if token == burst.EmptyStr {
		result.HttpBadRequest(w, model.TokenIsRequired.Error())
		return "", nil, nil
	}

	var proxyInfoReq req.RemoveProxyInfoReq
	err := httpx.ParseJsonBody(r, &proxyInfoReq)
	if err != nil {
		result.HttpBadRequest(w, err.Error())
		return "", nil, nil
	}
	err = proxyInfoReq.Check()
	if err != nil {
		result.HttpBadRequest(w, err.Error())
		return "", nil, nil
	}

	return token, &proxyInfoReq, nil
}
