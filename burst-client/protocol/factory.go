package protocol

import (
	"errors"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/wrappers"
)

func Decode(data []byte) (*BurstMessage, error) {
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

func GetError(message *BurstMessage) error {
	a := message.Header["error"]
	if a == nil {
		return nil
	}

	s := wrappers.StringValue{}
	_ = a.UnmarshalTo(&s)
	return errors.New(s.Value)
}
