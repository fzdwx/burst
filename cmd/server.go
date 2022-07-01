package main

import (
	"flag"
	"github.com/fzdwx/burst/pkg/logx"
	"github.com/fzdwx/burst/server"
	"github.com/fzdwx/burst/server/api"
	"github.com/fzdwx/burst/server/svc"
	"github.com/rs/zerolog"
	"github.com/zeromicro/go-zero/core/conf"
	zeroLogx "github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	"os"
)

var (
	sc      = flag.String("c", "server.yaml", "the config file path")
	sConfig server.Config
)

func init() {
	flag.Parse()

	conf.MustLoad(*sc, &sConfig)

	level := logx.GetLogLevel(sConfig.LogLevel)
	logx.UseLogLevel(level)
	out := os.Stdout
	logx.InitLogger(zerolog.ConsoleWriter{Out: out, TimeFormat: "2006/01/02 - 15:04:05"})
	zeroLogx.SetWriter(logx.NewZeroLogWrite(logx.GetLog()))
}

func main() {

	server := rest.MustNewServer(sConfig.RestConf)
	defer server.Stop()
	svcContext := svc.NewServiceContext(sConfig)

	api.MountRouters(server, svcContext)

	server.Start()
}
