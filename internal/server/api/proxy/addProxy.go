package proxy

import (
	"fmt"
	"github.com/fzdwx/burst/internal"
	cache "github.com/fzdwx/burst/internal/cache"
	"github.com/fzdwx/burst/internal/logx"
	"github.com/fzdwx/burst/internal/model"
	"github.com/fzdwx/burst/internal/model/req"
	"github.com/fzdwx/burst/internal/protocal"
	"github.com/fzdwx/burst/internal/result"
	"github.com/fzdwx/burst/internal/server/svc"
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

		// check if proxy is duplicated
		var proxyInfos []*internal.ServerProxyInfo
		for _, proxyInfo := range proxyInfoReq.Proxy {
			if info.Has(proxyInfo.Addr()) {
				result.HttpBadRequest(w, fmt.Sprintf("proxy %s already exists", proxyInfo.String()))
				return
			}
			proxyInfos = append(proxyInfos, proxyInfo.ToCache())
		}

		err, clientProxyInfos, closers := server.Lunch(proxyInfos)
		clean := func() {
			logx.Error().Str("token", token).Interface("proxy", proxyInfoReq).Msg("clean listeners")
			for _, c := range closers {
				c.Close()
			}
		}

		if err != nil {
			go clean()
			result.HttpBadRequest(w, err.Error())
			return
		}

		// notify client save proxy info
		bytes, err := protocal.NewAddProxy(clientProxyInfos).Encode()
		if err != nil {
			go clean()
			result.HttpBadRequest(w, err.Error())
			return
		}

		server.WriteBinary(bytes)
		cache.ProxyInfoContainer.Put(token, proxyInfos)

		httpx.OkJson(w, clientProxyInfos)
	}
}

func addProxyPreCheck(r *http.Request) (error, *req.AddProxyInfoReq, string) {
	token := internal.GetPars("token", r)
	if token == internal.EmptyStr {
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
