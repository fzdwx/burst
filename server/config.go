package server

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	LogLevel string `json:",default=debug"`
}
