package userRequest

import (
	"github.com/fzdwx/burst/client"
	"github.com/fzdwx/burst/pkg/protocal"
)

func Handle(c *client.Client, burst protocal.UserRequest) {
	c.WriteBinary([]byte("hello world"))
	c.WriteText("qqqqqqqqqqqqqqqqqqqqqqqq")
}
