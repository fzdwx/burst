package server

import (
	"github.com/fzdwx/burst/internal/proxy"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	LogLevel        string `json:",default=debug"`
	AloneHttpServer proxy.AloneHttpServerConfig
}
