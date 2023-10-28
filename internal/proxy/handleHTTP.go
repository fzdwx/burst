package proxy

import (
	"fmt"
	"github.com/fzdwx/burst/internal"
	"github.com/fzdwx/burst/internal/logx"
	"io"
	"net"
)

var AloneHttpServer = aloneHttpServer{
	running: false,
}

type (
	AloneHttpServerConfig struct {
		Ip        string   `json:",default=0.0.0.0"`
		Port      int      `json:",default=39939"`
		Enable    bool     `json:",default=true"`
		RouterKey []string `json:",routerKey,default=[Host,:authority]"`
	}

	aloneHttpServer struct {
		running bool
		cfg     AloneHttpServerConfig
	}
)

// Startup alone http server
func (ahs *aloneHttpServer) Startup(config AloneHttpServerConfig) {
	// ensure aloneHttpServer not running.
	if ahs.running {
		panic("repeated calls aloneHttpServer#Startup")
	}

	// update status
	ahs.running = true
	ahs.cfg = config
	if !ahs.cfg.Enable {
		logx.Debug().Msg("aloneHttpServer is disable")
		return
	}

	// bind port
	addr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", config.Ip, config.Port))
	if err != nil {
		panic(fmt.Errorf("aloneHttpServer resolve tcp adder fail:%v", err))
	}

	listen, err := net.ListenTCP("tcp", addr)
	if err != nil {
		panic(fmt.Errorf("aloneHttpServer bind port fail:%v", err))
	}
	logx.Debug().Msgf("aloneHttpServer startup success,listen on %s", addr)

	// accept user request.
	for {
		conn, err := listen.AcceptTCP()
		if err != nil {
			logx.Err(err).Msg("aloneHttpServer accept error")
		}
		go func() {
			for {
				bytes := make([]byte, 1024)
				n, err := conn.Read(bytes)
				if err != nil {
					logx.Err(err).Msg("aloneHttpServer  error")
				}

				if n == 0 {
					continue
				}
				logx.Info().Msgf("from client get message:%s", string(bytes[:n]))
			}
		}()

		/*
			todo
				1.获取 router key 对应的代理信息
				2.根据代理信息中的ip转发到
		*/
	}
}

// handlerHttp.
func (c *Container) handlerHttp(info *internal.ServerProxyInfo) (error, *internal.ClientProxyInfo, io.Closer) {
	return nil, nil, nil
}
