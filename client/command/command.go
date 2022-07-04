package command

import (
	"fmt"
	"github.com/fzdwx/burst/client"
	"github.com/zeromicro/go-zero/core/color"
	"os"
	"strings"
)

func Dispatch(line string, client *client.Client) {
	split := strings.Split(line, " ")
	switch strings.TrimSpace(split[0]) {
	case "u", "usage", "h", "help", "?":
		usage()
	case "q":
		client.Close()
		os.Exit(0)
	case "ap":
		addProxy(split[1:], client)
	case "log":
		//Log(client)
	case "version":
		//Version()
	default:
	}
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
