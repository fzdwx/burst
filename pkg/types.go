package pkg

import "errors"

type (
	ProxyInfo struct {
		Ip          string
		Port        int
		ChannelType string
		Addr        string
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
