package svc

import (
	"github.com/fzdwx/burst/common/wsx"
	"github.com/fzdwx/burst/server/internal/config"
)

type ServiceContext struct {
	Config config.Config
	Hub    *wsx.Hub
}

func NewServiceContext(c config.Config, hub *wsx.Hub) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Hub:    hub,
	}
}
