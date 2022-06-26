package dynServer

import "net"

type (
	Container struct {
		addProxy chan AddProxyReq
	}
)

func GetAvailablePort() (port int, err error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return 0, nil
	}
	listen, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, nil
	}
	defer listen.Close()
	return listen.Addr().(*net.TCPAddr).Port, nil
}
