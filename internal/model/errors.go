package model

import "errors"

var (
	ProxyIsRequired   = errors.New("proxy is required")
	TokenIsRequired   = errors.New("token is required")
	ProxyInfoNotFound = errors.New("the proxy info not found")
	IpIsBlank         = errors.New("ip is blank")
	IpIsNotValid      = errors.New("ip is not valid")
	PortIsNotValid    = errors.New("port is not valid")
	ClientNotFound    = errors.New("client not found")
	ServerClosed      = errors.New("client disconnected")
)
