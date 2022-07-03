package userRequest

import (
	"github.com/fzdwx/burst/client"
	"github.com/fzdwx/burst/pkg/logx"
	"github.com/fzdwx/burst/pkg/protocal"
)

// Handle user request
//
// write data to internet service
func Handle(c *client.Client, userRequest protocal.UserRequest) {
	internet, b := c.GetInternetService(userRequest.ConnId)
	if !b {
		logx.Error().Msg("internet not found, discard message")
		return
	}

	// write data to internet
	internet.Write(userRequest.Data)
}
