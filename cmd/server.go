package main

import (
	"flag"
	"github.com/fzdwx/burst/pkg/ginx"
	"github.com/fzdwx/burst/pkg/logx"
	"github.com/fzdwx/burst/server"
	"github.com/fzdwx/burst/server/api"
	"github.com/fzdwx/burst/server/svc"
	"github.com/jinzhu/configor"
)

var (
	c       = flag.String("c", "server.yml", "the config file path")
	sConfig = server.Config{}
)

func init() {
	flag.Parse()

	err := configor.Load(&sConfig, *c)
	if err != nil {
		logx.Fatal().Msg(err.Error())
	}

	level := logx.GetLogLevel(sConfig.LogLevel)
	logx.UseLogLevel(level)
}

func main() {
	svcContext := svc.NewServiceContext(sConfig)

	e := ginx.Classic()

	api.MountRouters(e, svcContext)

	err := e.Run(sConfig.Addr)
	logx.Fatal().Msg(err.Error())
	logx.Info()
}