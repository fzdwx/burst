package proxy

import (
	"github.com/fzdwx/burst"
	"github.com/fzdwx/burst/pkg"
	"github.com/fzdwx/burst/pkg/logx"
	"github.com/fzdwx/burst/pkg/model"
	"github.com/fzdwx/burst/pkg/model/req"
	"github.com/fzdwx/burst/pkg/protocal"
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

		var clientProxyInfos []pkg.ClientProxyInfo
		var serverProxyInfos []*pkg.ServerProxyInfo
		var cleans []func()
		for _, proxyInfo := range proxyInfoReq.Proxy {
			addr := proxyInfo.Addr()
			if info.Has(addr) {
				serverProxyInfo, _ := info.Get(addr)
				serverProxyInfos = append(serverProxyInfos, serverProxyInfo)
				clientProxyInfos = append(clientProxyInfos, *serverProxyInfo.ClientProxyInfo)

				cleans = append(cleans, func() {
					// remove proxy info and close listener from server
					// 1. close listener
					serverProxyInfo.BindListener.Close()
					// 2. remove proxy from cache
					info.Remove(addr)
					logx.Info().Int("serverPort", serverProxyInfo.ClientProxyInfo.ServerPort).Str("token", token).Str("client intranetAddr", serverProxyInfo.ClientProxyInfo.IntranetAddr).Msg("remove proxy")
				})
			}
		}

		if len(clientProxyInfos) == 0 {
			result.HttpOk(w, "the proxy is not found")
			return
		}

		// 3. notify client remove proxy
		bytes, err := protocal.NewRemoveProxy(clientProxyInfos).Encode()
		if err != nil {
			result.HttpBadRequest(w, err.Error())
			return
		}

		server.WriteBinary(bytes)

		// do clean
		go func() {
			for _, clean := range cleans {
				clean()
			}
		}()

		// update cache
		cache.ProxyInfoContainer.Put(token, serverProxyInfos)
		httpx.OkJson(w, clientProxyInfos)
	}
}

func removeProxyPreCheck(w http.ResponseWriter, r *http.Request) (string, *req.RemoveProxyInfoReq, error) {
	token := burst.GetPars("token", r)
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
