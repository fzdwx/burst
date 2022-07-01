package cache

import (
	"github.com/fzdwx/burst/pkg"
	cmap "github.com/orcaman/concurrent-map"
)

var (
	ProxyInfoContainer = proxyInfoCache{
		m: cmap.New(),
	}
)

type (
	proxyInfoCache struct {
		m cmap.ConcurrentMap
	}

	proxyInfos struct {
		arr []*pkg.ProxyInfo
	}
)

func (p *proxyInfoCache) Put(token string) {
	p.m.Set(token, proxyInfos{arr: []*pkg.ProxyInfo{}})
}

func (p proxyInfoCache) Has(token string) bool {
	return p.m.Has(token)
}

func (p proxyInfoCache) Get(token string) (*proxyInfos, bool) {
	if v, ok := p.m.Get(token); ok {
		infos := v.(proxyInfos)
		return &infos, ok
	}

	return nil, false
}

func (p *proxyInfos) Add(info pkg.ProxyInfo) {
	p.arr = append(p.arr, &info)
}