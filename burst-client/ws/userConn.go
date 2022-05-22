package cs

import (
	"github.com/fzdwx/burst/burst-client/protocol"
	log "github.com/sirupsen/logrus"
	"net"
	"strconv"
)

type (
	UserConnForward struct {
		userConnectId string
		conn          net.Conn
	}
	Forwarder struct {
		container map[string]*UserConnForward
	}
)

var (
	Fw = &Forwarder{
		container: make(map[string]*UserConnForward),
	}
)

func (f *Forwarder) add(forward *UserConnForward) {
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
			}).Error("Forwarder.Write")
		}
		log.Debugf("Forwarder.Write: %s, %d", userConnectId, write)
	}
}

func (f *Forwarder) Forward(message *protocol.BurstMessage) {
	userConnectId, err := protocol.GetUserConnectId(message)
	if err != nil {
		log.Error("parse user connect id error ", err)
		return
	}

	f.write(userConnectId, message.Data)
}

func (u UserConnForward) StartForwardToServer(client *Client) {
	userConnectId := u.userConnectId
	conn := u.conn
	defer conn.Close()
	defer log.Debug("forward to server: ", userConnectId, "closed")
	defer Fw.remove(userConnectId)

	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			log.Error("forward to server: ", "read error [", err, "]", userConnectId)
			return
		}

		// forward to server
		err = client.Write(userConnectId, buf[:n])
		if err != nil {
			log.Error("UserConnForward", "userConnectId", userConnectId, "write error", err)
			return
		}
		log.Debug("forward to server: ", "write [", n, "]", userConnectId)
	}
}

func NewUserConn(localPort int32, userConnectId string) (*UserConnForward, error) {
	conn, err := net.Dial("tcp", ":"+strconv.Itoa(int(localPort)))
	if err != nil {
		return nil, err
	}

	u := &UserConnForward{
		userConnectId,
		conn,
	}

	Fw.add(u)

	return u, nil
}
