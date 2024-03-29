package proxy

import (
	"github.com/fzdwx/burst/internal"
	"github.com/fzdwx/burst/internal/wsx"
	"io"
)

type (
	// Container Binding based on the token and the user's websocket connection
	Container struct {
		// the websocket connection to client
		*wsx.Wsx
		// the Token of client
		Token string
		// closers save this client all listeners(tcp/udp/http...) associated and connections from users
		closers []io.Closer
		// UserConnMap save all user connections,
		// key is conn id
		UserConnMap map[string]*UserConn
	}
)

func NewContainer(ws *wsx.Wsx, token string) *Container {
	return &Container{Wsx: ws, Token: token, closers: []io.Closer{}, UserConnMap: make(map[string]*UserConn)}
}

// Lunch Start the local service and then generate the format of the proxy information required by the client
func (c Container) Lunch(infos []*internal.ServerProxyInfo) (error, []internal.ClientProxyInfo, []io.Closer) {
	var (
		// mapping information used to return to the client
		clientInfos []internal.ClientProxyInfo
		// all listeners started by the current request
		listeners []io.Closer
	)

	for _, info := range infos {
		var (
			clientInfo *internal.ClientProxyInfo
			listener   io.Closer
			err        error
		)

		switch info.ChannelType {
		case internal.TCP:
			err, clientInfo, listener = c.handleTCP(info)
		case internal.HTTP:
			err, clientInfo, listener = c.handlerHttp(info)
		case internal.UDP:
			err, clientInfo, listener = c.handleUdp(info)
		}

		if err != nil {
			return err, nil, nil
		}

		if clientInfo == nil {
			return internal.NewError("unSupport channelType %s", internal.UDP), nil, nil
		}

		clientInfos = append(clientInfos, *clientInfo)
		listeners = append(listeners, listener)
	}

	return nil, clientInfos, listeners
}

// Close the local service
func (c Container) Close() {
	c.Wsx.Close()

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

func (c *Container) CleanUserConn(conn *UserConn) func() {
	return func() {
		conn.conn.Close()
		delete(c.UserConnMap, conn.Id)
	}
}
