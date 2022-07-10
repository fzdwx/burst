package command

import (
	"fmt"
	"github.com/fzdwx/burst/client"
	"os"
)

type (
	quitCommand struct {
	}
)

func (q quitCommand) usage() {
	fmt.Println("  q, quit, exit: quit client")
}

func (q quitCommand) run(strings []string, client *client.Client) {
	client.Close()
	infoMsg("Bye!")
	os.Exit(0)
}
