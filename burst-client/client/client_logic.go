package burst

import (
	"errors"
	"github.com/fzdwx/burst/burst-client/common"
	"github.com/fzdwx/burst/burst-client/protocol"
	log "github.com/sirupsen/logrus"
)

// HandlerBinaryData default logic
func HandlerBinaryData() OnBinary {
	return func(bytes []byte, client *Client) {
		burstMessage, err := protocol.Decode(bytes)
		if err != nil {
			log.Error(err)
			return
		}

		switch burstMessage.Type {
		case protocol.BurstType_ADD_PROXY_INFO:
			handlerAddProxyInfo(burstMessage, client)
		case protocol.BurstType_USER_CONNECT:
			handlerUserConnect(burstMessage, client)
		case protocol.BurstType_FORWARD_DATA:
			handlerForwardData(burstMessage, client)
		case protocol.BurstType_REMOVE_PROXY_INFO:
			handlerRemoveProxyInfo(burstMessage, client)
		}
	}
}

// handlerAddProxyInfo 处理添加映射信息
func handlerAddProxyInfo(message *protocol.BurstMessage, client *Client) {
	err := protocol.GetError(message)
	if err != nil {
		client.Over(errors.New("init error " + err.Error()))
	}

	ports, err := protocol.GetPorts(message)
	if err != nil {
		client.Over(errors.New("init get ports error " + err.Error()))
	}

	proxyInfo := ports.GetPorts()
	client.AddProxyInfo(proxyInfo)

	for serverExportPort, proxy := range proxyInfo {
		log.Infof("add proxy: %s to server %s", common.WrapRed(proxy.Host()), common.WrapRed(common.FormatToAddr(client.serverHostName, int(serverExportPort))))
	}
}

// handlerRemoveProxyInfo 移除映射信息
func handlerRemoveProxyInfo(message *protocol.BurstMessage, client *Client) {
	err := protocol.GetError(message)
	if err != nil {
		client.Over(errors.New("init error " + err.Error()))
	}

	port := message.GetServerPort()
	if len(port) == 0 {
		return
	}

	client.RemoveProxyPorts(port)
}

func handlerUserConnect(message *protocol.BurstMessage, client *Client) {
	serverExportPort, err := protocol.GetServerExportPort(message)
	if err != nil {
		log.Error("parse server export port error ", err)
		return
	}

	proxy, ok := client.GetProxy(serverExportPort)
	if !ok {
		log.Error("local port not found ", serverExportPort)
		return
	}

	userConnectId, err := protocol.GetUserConnectId(message)
	if err != nil {
		log.Error("parse user connect id error ", err)
		return
	}

	userConnForward, err := NewUserConn(serverExportPort, proxy, userConnectId)
	if err != nil {
		log.Error("local port connect error ", err)
		return
	}

	// step 5 [forwarded to the server], and then forwarded to a specific user
	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Error("forward to server: recover ", err)
			}
		}()
		userConnForward.React(client)
	}()
}

func handlerForwardData(message *protocol.BurstMessage, client *Client) {
	// step 4 [forward to local port]
	Fw.ToLocal(message)
}
