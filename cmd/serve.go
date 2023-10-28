package cmd

import (
	"github.com/fzdwx/burst/internal/logx"
	"github.com/fzdwx/burst/internal/proxy"
	"github.com/fzdwx/burst/internal/server"
	"github.com/fzdwx/burst/internal/server/api"
	"github.com/fzdwx/burst/internal/server/svc"
	"github.com/spf13/cobra"
	"github.com/zeromicro/go-zero/rest"
)

var (
	serve = &cobra.Command{
		Use: "serve",
		Run: func(cmd *cobra.Command, args []string) {
			loadLog()

			var sConfig = server.Config{
				RestConf: rest.RestConf{
					Port: port,
				},
				LogLevel: logLevel,
				AloneHttpServer: proxy.AloneHttpServerConfig{
					Ip:        aloneHttpServerIp,
					Port:      aloneHttpServerPort,
					Enable:    aloneHttpServerEnable,
					RouterKey: aloneHttpServerRouterKey,
				},
			}

			s := rest.MustNewServer(sConfig.RestConf)
			defer s.Stop()
			svcContext := svc.NewServiceContext(sConfig)

			api.MountRouters(s, svcContext)

			// startup alone http server
			go proxy.AloneHttpServer.Startup(sConfig.AloneHttpServer)

			logx.Debug().Int("port", sConfig.Port).Msg("burst server start")
			s.Start()
		},
	}

	serverName               = "burstServer"
	port                     = 9999
	aloneHttpServerEnable    = true
	aloneHttpServerIp        = "0.0.0.0"
	aloneHttpServerPort      = 39939
	aloneHttpServerRouterKey = []string{"Host", ":authority"}
)
