package command

import (
	"fmt"
	"github.com/fzdwx/burst/internal/client"
	"strings"
)

type (
	usageCommand struct {
	}
)

func (u usageCommand) usage() {
	fmt.Println("  u: show usage")
	fmt.Println("      format: u [command]")
	fmt.Println("      example: u ap")
}

func (u usageCommand) run(s []string, client *client.Client) {
	var uf = usageAll

	if len(s) > 0 {
		c := commands[strings.TrimSpace(s[0])]
		if c != nil {
			uf = func() {
				c.usage()
			}
		} else {
			unknownCommand(s)
			return
		}
	}
	showVersion()
	fmt.Println("Usage:")
	uf()
}

var (
	usageAll = func() {
		for _, c := range commands {
			c.usage()
		}
	}
)
