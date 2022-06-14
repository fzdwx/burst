package burst

import (
	"github.com/fzdwx/burst/burst-client/common"
	"github.com/fzdwx/burst/burst-client/protocol"
	log "github.com/sirupsen/logrus"
	"net"
	"strings"
)

type (
	// UserConnect user connect.
	UserConnect struct {
		// userConnectId user Connect id (from server)
		userConnectId string
		// conn 与内网端口的连接
		// The connection to the Intranet port
		conn net.Conn
		// serverPort 用于在移除代理时关闭与内网的连接
		// This command is used to close the connection to the Intranet when the agent is removed
		serverPort int32
	}
)

// NewUserConn open a connection for the specified user id to listen on the mapped port on the intranet.
func NewUserConn(serverPort int32, proxy *protocol.Proxy, userConnectId string) (*UserConnect, error) {
	conn, err := net.Dial("tcp", proxy.Host())
	if err != nil {
		return nil, err
	}

	u := &UserConnect{
		userConnectId,
		conn,
		serverPort,
	}

	Fw.Add(u)

	return u, nil
}

// React read data from mapped port and forward to server.
func (u UserConnect) React(client *Client) {
	userConnectId := u.userConnectId
	conn := u.conn
	defer func() {
		log.WithFields(log.Fields{"userConnectId": userConnectId}).Infoln("close user connect")
		Fw.remove(userConnectId)
	}()

	for {
		buf := make([]byte, 1024)
		read, err := conn.Read(buf)
		if err != nil {
			// todo 是否要判断？strings.Contains 还是所有都直接打印日志
			if !strings.Contains(err.Error(), "use of closed network connection") {
				log.WithFields(log.Fields{"status": "read from intranet error", "cause": err, "userConnectId": userConnectId}).Errorf("forward to %s  :", common.WrapRed("server"))
			}
			return
		}

		// forward to server
		err = client.ToServer(userConnectId, buf[:read])
		if err != nil {
			log.WithFields(log.Fields{"status": "error", "cause": err, "userConnectId": userConnectId, "len": read}).Errorf("forward to %s  :", common.WrapRed("server"))
			return
		}

		if common.IsDebug() {
			log.WithFields(log.Fields{"status": "success", "userConnectId": userConnectId, "len": read}).Debugf("forward to %s  :", common.WrapRed("server"))
		}
	}
}

func (u *UserConnect) Close() error {
	return u.conn.Close()
}
