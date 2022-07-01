package pkg

type (
	ProxyInfo struct {
		Ip          string
		Port        int
		ChannelType string
	}
)

const (
	TCP  = "tcp"
	HTTP = "http"

	UDP = "udp"
)