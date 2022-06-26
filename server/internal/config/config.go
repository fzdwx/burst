package config

import (
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf

	WsConfig struct {
		// HandshakeTimeout specifies the duration for the handshake to complete.
		// unit is second.
		HandshakeTimeout int64 `json:",default=10"`

		// ReadBufferSize and WriteBufferSize specify I/O buffer sizes in bytes. If a buffer
		// size is zero, then buffers allocated by the HTTP client are used. The
		// I/O buffer sizes do not limit the size of the messages that can be sent
		// or received.
		ReadBufferSize  int `json:",default=8192"`
		WriteBufferSize int `json:",default=8192"`

		// EnableCompression specify if the client should attempt to negotiate per
		// message compression (RFC 7692). Setting this value to true does not
		// guarantee that compression will be supported. Currently only "no context
		// takeover" modes are supported.
		EnableCompression bool
	}
}
