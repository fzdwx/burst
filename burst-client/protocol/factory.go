package protocol

import (
	"github.com/golang/protobuf/proto"
)

func Unmarshal(data []byte) (*BurstMessage, error) {
	message := BurstMessage{}
	err := proto.Unmarshal(data, &message)
	if err != nil {
		return nil, err
	}
	return &message, nil
}

func GetPorts(message *BurstMessage) (*Ports, error) {
	a := message.Header["ports"]
	ports := &Ports{}
	err := proto.Unmarshal(a.GetValue(), ports)
	if err != nil {
		return nil, err
	}
	return ports, nil
}
