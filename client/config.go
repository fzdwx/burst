package client

type Config struct {
	Server struct {
		Port int    `default:"9999" required:"true"`
		Host string `required:"true"`
	}
	LogLevel string `default:"debug"`
}