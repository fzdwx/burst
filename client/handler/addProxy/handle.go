package addProxy

import (
	"github.com/fzdwx/burst/client"
	"github.com/fzdwx/burst/pkg/protocal"
)

// Handle handles addProxy burst.
func Handle(c *client.Client, a protocal.AddProxy) {
	c.AddProxy(a)
}
