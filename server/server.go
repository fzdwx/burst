package main

import (
	"flag"
	"fmt"
	"github.com/fzdwx/burst/common/wsx"
	"github.com/fzdwx/burst/server/internal/config"
	"github.com/fzdwx/burst/server/internal/handler"
	"github.com/fzdwx/burst/server/internal/svc"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "C:\\Users\\98065\\IdeaProjects\\fzdwx\\burst\\server\\etc\\server.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	upgrader := wsx.NewUpgrader(c.WsConfig)
	hub := wsx.NewHub(&upgrader)
	ctx := svc.NewServiceContext(c, hub)

	handler.RegisterHandlers(server, ctx)
	handler.RegisterCustomHandlers(server, ctx)

	fmt.Printf("Starting client at %s:%d...\n", c.Host, c.Port)
	go hub.React()
	server.Start()
}
