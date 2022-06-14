package burst

import (
	"github.com/fzdwx/burst/burst-client/common"
	"github.com/fzdwx/burst/burst-client/protocol"
	log "github.com/sirupsen/logrus"
)

type (
	// Forwarder store connections between real users and mapped ports.
	Forwarder struct {
		userConnectContainer map[string]*UserConnect
		// serverPortConnectContainer key : serverPort value: userConnectIds
		serverPortConnectContainer map[int32][]string
	}
)

var (
	// Fw default Forwarder
	Fw = &Forwarder{
		userConnectContainer:       make(map[string]*UserConnect),
		serverPortConnectContainer: map[int32][]string{},
	}
)

// ToLocal forward real user data to the mapped port on the intranet.
func (f *Forwarder) ToLocal(message *protocol.BurstMessage) {
	userConnectId, err := protocol.GetUserConnectId(message)
	if err != nil {
		log.Error("forward to local: parse user connect id error ", err)
		return
	}

	f.write(userConnectId, message.Data)
}

// Add  UserConnect to userConnectContainer
func (f *Forwarder) Add(userConnect *UserConnect) {
	f.userConnectContainer[userConnect.userConnectId] = userConnect

	userConnectIds := f.serverPortConnectContainer[userConnect.serverPort]
	if userConnectIds == nil {
		userConnectIds = []string{userConnect.userConnectId}
	} else {
		userConnectIds = append(userConnectIds, userConnect.userConnectId)
	}
	f.serverPortConnectContainer[userConnect.serverPort] = userConnectIds
}

// RemoveProxyPort remove serverPort from serverPortConnectContainer and close conn
func (f *Forwarder) RemoveProxyPort(serverPort int32) {
	userConnectIds := f.serverPortConnectContainer[serverPort]
	if userConnectIds == nil {
		return
	}

	defer delete(f.serverPortConnectContainer, serverPort)

	for _, userConnectId := range userConnectIds {
		f.remove(userConnectId)
	}
}

// remove user connect from userConnectContainer and close conn.
func (f *Forwarder) remove(key string) {
	userConnect := f.userConnectContainer[key]
	if userConnect == nil {
		return
	}
	defer delete(f.userConnectContainer, key)

	userConnect.Close()
}

func (f *Forwarder) write(userConnectId string, data []byte) {
	if forward, ok := f.userConnectContainer[userConnectId]; ok {
		write, err := forward.conn.Write(data)
		if err != nil {
			log.WithFields(log.Fields{"userConnectId": userConnectId, "status": "error", "len": write, "cause": err}).Errorf("forward to %s:", common.WrapCyan("intranet"))
		}

		if common.IsDebug() {
			log.WithFields(log.Fields{"userConnectId": userConnectId, "len": write, "status": "success"}).Debugf("forward to %s:", common.WrapCyan("intranet"))
		}
	}
}
