package command

import (
	"fmt"
	"github.com/fzdwx/burst/internal/client"
	"github.com/knz/bubbline"
	"github.com/zeromicro/go-zero/core/color"
	"strings"
)

type command interface {
	usage()
	autocomplete() bubbline.AutoCompleteFn
	run([]string, *client.Client)
}

var (
	commands = map[string]command{}
	version  = "2.1.0"

	commandNames = []string{
		"usage", "quit", "addProxy", "removeProxy",
	}
)

func init() {
	commands["usage"] = &usageCommand{}
	commands["quit"] = &quitCommand{}
	commands["addProxy"] = &addProxyCommand{}
	commands["removeProxy"] = &removeProxyCommand{}
}

func Dispatch(line string, client *client.Client) {
	pairs := strings.Split(line, " ")

	switch strings.TrimSpace(pairs[0]) {
	case "u", "usage", "h", "help", "?":
		commands["usage"].run(pairs[1:], client)
	case "q", "quit", "exit":
		commands["quit"].run(pairs, client)
	case "ap", "addProxy":
		commands["addProxy"].run(pairs[1:], client)
	case "rp", "removeProxy":
		commands["removeProxy"].run(pairs[1:], client)
	case "log":
		//Log(client)
	case "version":
		showVersion()
	default:
		unknownCommand(pairs)
	}
}

func showVersion() {
	fmt.Printf("burst-cli %s\n", version)
}

func unknownCommand(cmd []string) {
	errorMsg("unknown command: " + cmd[0])
}

func errorMsg(msg string) {
	print(msg, color.BgRed)
}

func infoMsg(msg string) {
	print(msg, color.FgGreen)
}

func print(msg string, colour color.Color) {
	fmt.Println(color.WithColor("[cli]", color.FgBlue)+":", color.WithColor(msg, colour))
}
