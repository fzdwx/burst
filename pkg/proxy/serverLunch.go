package proxy

import (
	"github.com/fzdwx/burst"
	"github.com/fzdwx/burst/pkg"
	"github.com/fzdwx/burst/pkg/wsx"
	"io"
)

type (
	Container struct {
		*wsx.Wsx
		Token string
		// closers save this client all listeners(tcp/udp/http...) associated and connections from users
		closers     []io.Closer
		UserConnMap map[string]*UserConn
	}
)

func NewContainer(ws *wsx.Wsx, token string) *Container {
	return &Container{Wsx: ws, Token: token, closers: []io.Closer{}, UserConnMap: make(map[string]*UserConn)}
}

// Lunch Start the local service and then generate the format of the proxy information required by the client
//
func (c Container) Lunch(infos []*pkg.ServerProxyInfo) (error, []pkg.ClientProxyInfo, []io.Closer) {
	var (
		clientInfos []pkg.ClientProxyInfo
		closers     []io.Closer
	)

	for _, info := range infos {
		var (
			clientInfo *pkg.ClientProxyInfo
			listener   io.Closer
			err        error
		)

		switch info.ChannelType {
		case pkg.TCP:
			err, clientInfo, listener = c.handleTCP(info)
		case pkg.HTTP:
			err, clientInfo, listener = c.handlerHttp(info)
		case pkg.UDP:
			err, clientInfo, listener = c.handleUdp(info)
		}

		// todo clean closers
		if err != nil {
			return err, nil, nil
		}

		if clientInfo == nil {
			return burst.NewError("unSupport channelType %s", pkg.UDP), nil, nil
		}

		clientInfos = append(clientInfos, *clientInfo)
		closers = append(closers, listener)
	}

	return nil, clientInfos, closers
}

// Close the local service
func (c Container) Close() {
	for _, c := range c.closers {
		c.Close()
	}
}

// AddCloser add closer
func (c *Container) AddCloser(closer io.Closer) {
	c.closers = append(c.closers, closer)
}

func (c *Container) AddUserConn(conn *UserConn) {
	c.UserConnMap[conn.Id] = conn
}

func (c *Container) GetUserConn(connId string) (*UserConn, bool) {
	userConn, ok := c.UserConnMap[connId]
	return userConn, ok
}