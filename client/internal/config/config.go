package config

import (
	"fmt"
	"github.com/zeromicro/go-zero/rest"
	"net/url"
)

type Config struct {
	rest.RestConf

	Burst BurstConfig `json:",optional"`
}

type (
	// BurstConfig about the Burst config
	BurstConfig struct {
		Host     string `json:""`               // Host the burst client address
		Port     int    `json:",default=10086"` // Port the burst client  port
		HttpPort int    `json:",default=39399"` // HttpPort the burst client fixed http port
		LogLevel string `json:",default=info"`  // LogLevel the burst client log level
	}
)

// address return the burst client address
func (b BurstConfig) address() string {
	if b.Host == "" {
		return fmt.Sprintf("localhost:%d", b.Port)
	}
	return fmt.Sprintf("%s:%d", b.Host, b.Port)
}

func (b BurstConfig) BuildWsUrl(token string) url.URL {
	return url.URL{Scheme: "ws", Host: b.address(), Path: "/connect", RawQuery: "token=" + token}
}
