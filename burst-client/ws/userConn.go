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
			}).Error("forward to local: error")
		}
		log.Debugf("forward to local: write [%d],  %s", write, userConnectId)
	}
}

func (f *Forwarder) ToLocal(message *protocol.BurstMessage) {
	userConnectId, err := protocol.GetUserConnectId(message)
	if err != nil {
		log.Error("forward to local: parse user connect id error ", err)
		return
	}

	f.write(userConnectId, message.Data)
}

func (u UserConnForward) StartForwardToServer(client *Client) {
	userConnectId := u.userConnectId
	conn := u.conn
	defer conn.Close()
	defer log.Debug("forward to server: closed ", userConnectId)
	defer Fw.remove(userConnectId)

	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			log.Error("forward to server: ", "read error [", err, "] ", userConnectId)
			return
		}

		// forward to server
		err = client.Write(userConnectId, buf[:n])
		if err != nil {
			log.Error("forward to server: ", "write error [", err, "] ", userConnectId)
			return
		}
		log.Debug("forward to server: ", "write [", n, "]  ", userConnectId)
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
