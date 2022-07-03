package proxy

import (
	"github.com/fzdwx/burst/pkg"
	"io"
)

func (c *Container) handlerHttp(info *pkg.ServerProxyInfo) (error, *pkg.ClientProxyInfo, io.Closer) {
	// todo handler http
	return nil, nil, nil
}
