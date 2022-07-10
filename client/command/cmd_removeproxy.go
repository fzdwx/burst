package command

import (
	"fmt"
	"github.com/fzdwx/burst/client"
)

type (
	removeProxyCommand struct {
	}
)

func (r removeProxyCommand) usage() {
	fmt.Println("  rp: remove proxy ")
	fmt.Println("      format: rp [channelType]:[ip]:[port]")
	fmt.Println("      example: rp tcp::8888 tcp::63342 tcp:192.168.1.1:8080 ...")
}

func (r removeProxyCommand) run(strings []string, c *client.Client) {
}
