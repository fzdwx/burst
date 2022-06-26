package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf

	// Burst about the Burst config
	Burst struct {
		Host     string `json:""`               // Host the burst server address
		Port     int    `json:",default=10086"` // Port the burst server  port
		HttpPort int    `json:",default=39399"` // HttpPort the burst server fixed http port
	}
}
