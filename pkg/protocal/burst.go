package protocal

import (
	"encoding/json"
	"github.com/fzdwx/burst/pkg"
)

type (
	Burst struct {
		Type string `json:"type"`

		AddProxy AddProxy `json:"addProxy,omitempty"`

		UserConnect UserConnect `json:"userConnect,omitempty"`

		UserRequest UserRequest `json:"userRequest,omitempty"`

		IntranetResponse IntranetResponse `json:"internetResponse,omitempty"`
	}

	AddProxy struct {
		Proxy []pkg.ClientProxyInfo `json:"proxy,omitempty"`
	}

	UserConnect struct {
		Key    string `json:"key"`
		ConnId string `json:"connId"`
	}

	UserRequest struct {
		Data   []byte `json:"data"`
		Key    string `json:"key"`
		ConnId string `json:"connId"`
	}

	IntranetResponse struct {
		Data   []byte `json:"data"`
		ConnId string `json:"connId"`
		Token  string `json:"token"`
	}
)

/* to client */
const (
	UserRequestType = "userRequest"
	AddProxyType    = "addProxy"
	UserConnectType = "userConnect"
)

/* to server */
const (
	IntranetResponseType = "intranetResponse"
)

func (b Burst) Encode() ([]byte, error) {
	return json.Marshal(b)
}

func Decode(bytes []byte) (Burst, error) {
	var b Burst
	err := json.Unmarshal(bytes, &b)
	return b, err
}

func NewAddProxy(proxy []pkg.ClientProxyInfo) Burst {
	return Burst{
		Type: AddProxyType,
		AddProxy: AddProxy{
			Proxy: proxy,
		},
	}
}

func NewUserConnect(key string, connId string) Burst {
	return Burst{
		Type: UserConnectType,
		UserConnect: UserConnect{
			Key:    key,
			ConnId: connId,
		},
	}
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

func NewIntranetResponse(data []byte, token string, connId string) Burst {
	return Burst{
		Type: IntranetResponseType,
		IntranetResponse: IntranetResponse{
			Data:   data,
			Token:  token,
			ConnId: connId,
		},
	}
}
