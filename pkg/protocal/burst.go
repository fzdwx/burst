package protocal

import (
	"encoding/json"
	"github.com/fzdwx/burst/pkg"
)

type (
	Burst struct {
		Type string `json:"type"`

		AddProxy AddProxy `json:"addProxy,omitempty"`

		UserRequest UserRequest `json:"userRequest,omitempty"`
	}

	AddProxy struct {
		Proxy []pkg.ClientProxyInfo `json:"proxy,omitempty"`
	}

	UserRequest struct {
		Data   []byte `json:"data"`
		Key    string `json:"key"`
		ConnId string `json:"connId"`
	}
)

const (
	UserRequestType = "userRequest"
	AddProxyType    = "addProxy"
)

func (b Burst) Encode() ([]byte, error) {
	return json.Marshal(b)
}

func Decode(bytes []byte) (Burst, error) {
	var b Burst
	err := json.Unmarshal(bytes, &b)
	return b, err
}

func NewUserRequest(data []byte, key string, connId string) Burst {
	return Burst{
		Type: UserRequestType,
		UserRequest: UserRequest{
			Data:   data,
			Key:    key,
			ConnId: connId,
		},
	}
}

func NewAddProxy(proxy []pkg.ClientProxyInfo) Burst {
	return Burst{
		Type: AddProxyType,
		AddProxy: AddProxy{
			Proxy: proxy,
		},
	}
}
