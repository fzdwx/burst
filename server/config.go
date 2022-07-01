package server

type Config struct {
	Addr     string `default:":9999" required:"true"`
	LogLevel string `default:"debug"`
}