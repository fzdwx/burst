package burst

import (
	"fmt"
)

const (
	EmptyStr = ""
)

func Over(errorMessage string) {
	panic(errorMessage)
}

func FormatAddr(host string, port int) string {
	return fmt.Sprintf("%s:%d", host, port)
}
