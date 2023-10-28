package command

import (
	"fmt"
	"github.com/fzdwx/burst/internal/client"
	"github.com/knz/bubbline"
	"os"
)

type (
	quitCommand struct {
	}
)

func (q quitCommand) autocomplete() bubbline.AutoCompleteFn {
	//TODO implement me
	panic("implement me")
}

func (q quitCommand) usage() {
	fmt.Println("  q, quit, exit: quit client")
}

func (q quitCommand) run(strings []string, client *client.Client) {
	client.Close()
	infoMsg("Bye!")
	os.Exit(0)
}
