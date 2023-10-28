package proxy

import (
	"fmt"
	"github.com/fzdwx/burst/internal"
	"io"
	"net"
)

// handleUdp handler udp todo
func (c *Container) handleUdp(info *internal.ServerProxyInfo) (error, *internal.ClientProxyInfo, io.Closer) {
	udpConn, err := net.ListenUDP(info.ChannelType, nil)
	if err != nil {
		return err, nil, nil
	}

	fmt.Println("udp port:", udpConn.LocalAddr().(*net.UDPAddr).Port)
	go func() {
		for {
			buf := make([]byte, 1024)
			udp, addr, err := udpConn.ReadFromUDP(buf[:])
			if err != nil {
				fmt.Println("error:", err)
			}
			fmt.Println(string(buf[:udp]))
			fmt.Println(addr.Port)
		}
	}()
	return nil, nil, udpConn
}
