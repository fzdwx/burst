package cache

import (
	"github.com/fzdwx/burst/pkg/proxy"
	"github.com/fzdwx/burst/pkg/wsx"
	"github.com/zeromicro/go-zero/core/collection"
)

var (
	ServerContainer ServerCache
)

type (
	ServerCache struct {
		m *collection.Cache
	}
)

func (c *ServerCache) Put(token string, ws *wsx.Wsx) {
	c.m.Set(token, proxy.NewContainer(ws, token))
}

func (c *ServerCache) Get(token string) (*proxy.Container, bool) {
	client, b := c.m.Get(token)
	if !b {
		return nil, false
	}
	return client.(*proxy.Container), true
}

func (c *ServerCache) Remove(token string) {
	container, b := c.Get(token)
	if b {
		go container.Close()
		c.m.Del(token)
	}
}
