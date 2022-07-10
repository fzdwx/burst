package command

import (
	"fmt"
	"github.com/fzdwx/burst/client"
	"github.com/zeromicro/go-zero/core/color"
	"strings"
)

type command interface {
	usage()
	run([]string, *client.Client)
}

var (
	commands = map[string]command{}
	version  = "2.0.0"
)

func init() {
	commands["u"] = &usageCommand{}
	commands["quit"] = &quitCommand{}
	commands["ap"] = &addProxyCommand{}
	commands["rp"] = &removeProxyCommand{}
}

func Dispatch(line string, client *client.Client) {
	split := strings.Split(line, " ")
	switch strings.TrimSpace(split[0]) {
	case "u", "usage", "h", "help", "?":
		commands["u"].run(split[1:], client)
	case "q", "quit", "exit":
		commands["quit"].run(split, client)
	case "ap":
		commands["ap"].run(split[1:], client)
	case "rp":
		commands["rp"].run(split[1:], client)
	case "log":
		//Log(client)
	case "version":
		showVersion()
	default:
		unknownCommand(split)
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
