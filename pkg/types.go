package pkg

import "errors"

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
