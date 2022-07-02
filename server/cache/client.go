package cache

import (
	"github.com/fzdwx/burst/pkg/wsx"
	"github.com/zeromicro/go-zero/core/collection"
)

var (
	ClientContainer clientCache
)

type (
	clientCache struct {
		m *collection.Cache
	}
)

func (c *clientCache) Put(ws *wsx.Wsx) {
	c.m.Set(ws.IdStr, ws)
}

func (c *clientCache) Remove(ws *wsx.Wsx) {
	c.m.Del(ws.IdStr)
}
