package protocal

import "encoding/json"

type (
	UserRequest struct {
		Data       []byte `json:"data"`
		ServerPort int    `json:"serverPort"`
		ConnId     string `json:"connId"`
	}
)

func (r UserRequest) Encode() ([]byte, error) {
	return json.Marshal(r)
}
