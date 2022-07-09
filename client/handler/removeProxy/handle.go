package removeProxy

import (
	"github.com/fzdwx/burst/client"
	"github.com/fzdwx/burst/pkg/protocal"
)

func Handle(c *client.Client, removeProxy protocal.RemoveProxy) {
	c.RemoveProxy(removeProxy)

	c.CloseInternet(removeProxy.Proxy)
}
