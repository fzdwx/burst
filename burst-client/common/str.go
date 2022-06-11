package common

import "fmt"

// FormatToAddr format to host:port
func FormatToAddr(host string, port int) string {
	return fmt.Sprintf("%s:%d", host, port)
}
