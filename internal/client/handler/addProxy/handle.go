package addProxy

import (
	"github.com/fzdwx/burst/internal/client"
	"github.com/fzdwx/burst/internal/protocal"
)

// Handle handles addProxy burst.
func Handle(c *client.Client, a protocal.AddProxy) {
	c.AddProxy(a)
}
