package burst

import (
	"github.com/fzdwx/burst/burst-client/common"
	"github.com/fzdwx/burst/burst-client/protocol"
	log "github.com/sirupsen/logrus"
	"net"
	"strconv"
)

type (
	// UserConnect user connect.
	UserConnect struct {
		userConnectId string
		conn          net.Conn
	}

	// Forwarder store connections between real users and mapped ports.
	Forwarder struct {
		container map[string]*UserConnect
	}
)

var (
	// Fw default Forwarder
	Fw = &Forwarder{
		container: make(map[string]*UserConnect),
	}
)

// NewUserConn open a connection for the specified user id to listen on the mapped port on the intranet.
func NewUserConn(localPort int32, userConnectId string) (*UserConnect, error) {
	conn, err := net.Dial("tcp", ":"+strconv.Itoa(int(localPort)))
	if err != nil {
		return nil, err
	}

	u := &UserConnect{
		userConnectId,
		conn,
	}

	Fw.add(u)

	return u, nil
}

// React read data from mapped port and forward to server.
func (u UserConnect) React(client *Client) {
	userConnectId := u.userConnectId
	conn := u.conn
	defer conn.Close()
	defer Fw.remove(userConnectId)
	defer log.Debug("forward to server: closed ", userConnectId)

	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			log.Errorf("forward to server: read error:[%s] userConnectId:%s", err, userConnectId)
			return
		}

		// forward to server
		err = client.ToServer(userConnectId, buf[:n])
		if err != nil {
			log.Errorf("forward to server: write error:[%s] userConnectId:%s", err, userConnectId)
			return
		}

		if common.IsDebug() {
			log.Debugf("forward to server: write size:%d,userConnectId:%s", n, userConnectId)
		}
	}
}

// ToLocal forward real user data to the mapped port on the intranet.
func (f *Forwarder) ToLocal(message *protocol.BurstMessage) {
	userConnectId, err := protocol.GetUserConnectId(message)
	if err != nil {
		log.Error("forward to local: parse user connect id error ", err)
		return
	}

	f.write(userConnectId, message.Data)
}

func (f *Forwarder) add(forward *UserConnect) {
	f.container[forward.userConnectId] = forward
}

func (f *Forwarder) remove(key string) {
	delete(f.container, key)
}

func (f *Forwarder) write(userConnectId string, data []byte) {
	if forward, ok := f.container[userConnectId]; ok {
		write, err := forward.conn.Write(data)
		if err != nil {
			log.WithFields(log.Fields{
				"userConnectId": userConnectId,
				"write":         write,
				"err":           err,
			}).Error("forward to local: error")
		}

		if common.IsDebug() {
			log.Debugf("forward to local: write [%d],  %s", write, userConnectId)
		}
	}
}
