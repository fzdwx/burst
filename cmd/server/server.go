package main

import (
	"flag"
	"github.com/fzdwx/burst/pkg/logx"
	"github.com/fzdwx/burst/server"
	"github.com/fzdwx/burst/server/api"
	"github.com/fzdwx/burst/server/svc"
	"github.com/rs/zerolog"
	"github.com/zeromicro/go-zero/core/conf"
	zlog "github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/zero-contrib/logx/zerologx"
	"os"
)

var (
	sc      = flag.String("c", "server.yaml", "the config file path")
	logFile = flag.String("l", "", "the log file path, e.g: ./server.log")
	sConfig server.Config
)

func init() {
	flag.Parse()

	conf.MustLoad(*sc, &sConfig)

	if *logFile != "" {
		out, err := os.OpenFile(*logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			logx.Fatal().Err(err).Msg("open log file fail")
		}
		logx.InitLogger(out)
	} else {
		logx.InitLogger(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "2006/01/02 - 15:04:05"})
	}

	logx.UseLogLevel(logx.GetLogLevel(sConfig.LogLevel))
	zlog.SetWriter(zerologx.NewZeroLogWriter(logx.GetLog()))
}

func main() {

	server := rest.MustNewServer(sConfig.RestConf)
	defer server.Stop()
	svcContext := svc.NewServiceContext(sConfig)

	api.MountRouters(server, svcContext)

	server.Start()
}
