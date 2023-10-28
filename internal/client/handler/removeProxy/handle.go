package removeProxy

import (
	"github.com/fzdwx/burst/internal/client"
	"github.com/fzdwx/burst/internal/protocal"
)

func Handle(c *client.Client, removeProxy protocal.RemoveProxy) {
	c.RemoveProxy(removeProxy)

	c.CloseInternet(removeProxy.Proxy)
}
