package burst

import (
	"fmt"
	"github.com/zeromicro/go-zero/rest/pathvar"
	"net/http"
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

func GetQuery(key string, r *http.Request) string {
	return r.URL.Query().Get(key)
}

func GetPars(key string, r *http.Request) string {
	vars := pathvar.Vars(r)
	return vars[key]
}
