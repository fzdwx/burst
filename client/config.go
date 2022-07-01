package client

type Config struct {
	Server struct {
		Port int    `json:",default=9999"`
		Host string `json:",required=true"`
	}
	LogLevel string `json:",default=debug"`
}
