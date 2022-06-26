package main

import (
	"flag"
	"github.com/fzdwx/burst/client/internal/client"
	"github.com/fzdwx/burst/client/internal/config"
	"github.com/fzdwx/burst/common/errx"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "C:\\Users\\98065\\IdeaProjects\\fzdwx\\burst\\client\\etc\\client.yaml", "the config file")
var token = flag.String("t", "dev", "the access token")

func main() {
	flag.Parse()
	if errx.CheckToken(*token) {
		logx.Must(errx.ErrTokenIsRequired)
	}

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	client.Connect(c.Burst.BuildWsUrl(*token), *token)
	server.Start()
}
