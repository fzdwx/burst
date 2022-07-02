package cache

import (
	"github.com/fzdwx/burst/pkg/logx"
	"github.com/zeromicro/go-zero/core/collection"
	"time"
)

func init() {
	ProxyInfoContainer = proxyInfoCache{m: newCache("proxyInfo")}
	ServerContainer = ServerCache{m: newCache("client")}
}

func newCache(name string) *collection.Cache {
	proxyInfo, err := collection.NewCache(24*time.Hour, collection.WithName(name))
	if err != nil {
		logx.Error().Err(err).Msgf("init %s cache fail", name)
	}
	return proxyInfo
}
