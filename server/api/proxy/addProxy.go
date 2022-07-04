package proxy

import (
	"fmt"
	"github.com/fzdwx/burst"
	"github.com/fzdwx/burst/pkg"
	"github.com/fzdwx/burst/pkg/model"
	"github.com/fzdwx/burst/pkg/model/req"
	"github.com/fzdwx/burst/pkg/protocal"
	"github.com/fzdwx/burst/pkg/result"
	"github.com/fzdwx/burst/server/cache"
	"github.com/fzdwx/burst/server/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func AddProxy(svcContext *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err, proxyInfoReq, token := addProxyPreCheck(r)
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

		var proxyInfos []*pkg.ServerProxyInfo

		// check if proxy is duplicated
		for _, proxyInfo := range proxyInfoReq.Proxy {
			if info.Has(proxyInfo.Addr()) {
				result.HttpBadRequest(w, fmt.Sprintf("proxy %s already exists", proxyInfo.String()))
				return
			}
			proxyInfos = append(proxyInfos, proxyInfo.ToCache())
		}

		err, clientProxyInfos, closers := server.Lunch(proxyInfos)
		clean := func() {
			for _, c := range closers {
				c.Close()
			}
		}

		if err != nil {
			go clean()
			result.HttpBadRequest(w, err.Error())
			return
		}

		bytes, err := protocal.NewAddProxy(clientProxyInfos).Encode()
		if err != nil {
			go clean()
			result.HttpBadRequest(w, err.Error())
			return
		}

		// send client proxy info to client
		server.WriteBinary(bytes)

		cache.ProxyInfoContainer.Put(token, proxyInfos)

		httpx.OkJson(w, clientProxyInfos)
	}
}

func addProxyPreCheck(r *http.Request) (error, *req.AddProxyInfoReq, string) {
	token := burst.GetPars("token", r)
	if token == burst.EmptyStr {
		return model.TokenIsRequired, nil, ""
	}

	var proxyInfoReq req.AddProxyInfoReq
	err := httpx.ParseJsonBody(r, &proxyInfoReq)
	if err != nil {
		return err, nil, ""
	}

	err = proxyInfoReq.Check()
	if err != nil {
		return err, nil, ""
	}
	return nil, &proxyInfoReq, token
}
