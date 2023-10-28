package internal

import (
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/netx"
	"github.com/zeromicro/go-zero/rest/pathvar"
	"net"
	"net/http"
)

const (
	EmptyStr = ""
)

var (
	CurrentIp string
)

func GetCurrentIp() string {
	if CurrentIp == "" {
		CurrentIp = netx.InternalIp()
	}

	return CurrentIp
}

func Over(errorMessage string) {
	panic(errorMessage)
}

func FormatAddr(host string, port int) string {
	return fmt.Sprintf("%s:%d", host, port)
}

func GetQuery(key string, r *http.Request) string {
	return r.URL.Query().Get(key)
}

func GetPars(key string, r *http.Request) string {
	vars := pathvar.Vars(r)
	return vars[key]
}

func NewError(format string, a ...any) error {
	return errors.New(fmt.Sprintf(format, a))
}

func GetFreePort() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:0")
	if err != nil {
		return 0, err
	}

	l, err := net.Listen("tcp", addr.String())
	if err != nil {
		return 0, err
	}
	defer l.Close()
	return l.Addr().(*net.TCPAddr).Port, nil
}
