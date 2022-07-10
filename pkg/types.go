package pkg

import (
	"errors"
	"fmt"
	"github.com/fzdwx/burst"
	"io"
)

type (
	ServerProxyInfo struct {
		// Ip The intranet IP address of the service that the user submits to be exposed
		Ip string
		// Port The port of the service submitted by the user to be exposed
		Port int
		// ChannelType channel type
		ChannelType string
		// Addr format Ip : Port
		Addr string

		// use on remove proxy
		ClientProxyInfo *ClientProxyInfo
		BindListener    io.Closer
		BindUserConn    []io.Closer
	}

	ClientProxyInfo struct {
		// ServerPort The port exposed by the server
		ServerPort int
		// IntranetAddr The address of the intranet service
		IntranetAddr string
		// ChannelType channel type
		ChannelType string
	}
)

const (
	TCP  = "tcp"
	HTTP = "http"

	UDP = "udp"
)

var (
	ErrChannelTypeNotValid = errors.New("channel type is not valid")
)

func CheckChannelType(str string) bool {
	switch str {
	case TCP, HTTP, UDP:
		return true
	}
	return false
}
func (cp ClientProxyInfo) Key() string {
	switch cp.ChannelType {
	case HTTP:
		return "todo"
	default:
		return fmt.Sprint(cp.ServerPort)
	}
}

func (cp ClientProxyInfo) Address(serverAddr string) string {
	switch cp.ChannelType {
	case HTTP:
		return "todo"
	default:
		return burst.FormatAddr(serverAddr, cp.ServerPort)
	}
}
