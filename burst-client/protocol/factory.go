package protocol

import (
	"errors"
	"github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	"github.com/golang/protobuf/ptypes/wrappers"
	proto2 "google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"strconv"
)

/**
  tool method
*/

// Host format ip:port
func (x *Proxy) Host() string {
	return x.Ip + ":" + strconv.Itoa(int(x.Port))
}

func Decode(data []byte) (*BurstMessage, error) {
	message := BurstMessage{}
	err := proto.Unmarshal(data, &message)
	if err != nil {
		return nil, err
	}
	return &message, nil
}

func Encode(userConnectIdStr string, data []byte, tokenStr string) []byte {
	message := BurstMessage{Header: map[int32]*any.Any{}}
	userConnectId, _ := anypb.New(&wrappers.StringValue{Value: userConnectIdStr})
	token, _ := anypb.New(&wrappers.StringValue{Value: tokenStr})
	message.Header[int32(Headers_USER_CONNECT_ID)] = userConnectId
	message.Header[int32(Headers_TOKEN)] = token
	message.Type = BurstType_FORWARD_DATA
	message.Data = data
	d, _ := proto.Marshal(&message)
	return d
}

func GetPorts(message *BurstMessage) (*Ports, error) {
	p, err := getStruct(message, Headers_PORTS, &Ports{})
	if err != nil {
		return nil, err
	}
	return p.(*Ports), nil
}

func GetError(message *BurstMessage) error {
	a := message.Header[int32(Headers_ERROR)]
	if a == nil {
		return nil
	}

	s := wrappers.StringValue{}
	_ = a.UnmarshalTo(&s)
	return errors.New(s.Value)
}

func GetServerExportPort(message *BurstMessage) (int32, error) {
	return getInt32(message, Headers_SERVER_EXPORT_PORT)
}

func GetUserConnectId(message *BurstMessage) (string, error) {
	return getString(message, Headers_USER_CONNECT_ID)
}

func getInt32(message *BurstMessage, header Headers) (int32, error) {
	p, err := get(message, header, &wrappers.Int32Value{})
	if err != nil {
		return 0, err
	}
	return p.(*wrappers.Int32Value).Value, nil
}

func getStruct(message *BurstMessage, headers Headers, p proto.Message) (proto.Message, error) {
	a := message.Header[int32(headers)]
	err := proto.Unmarshal(a.GetValue(), p)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func getString(message *BurstMessage, header Headers) (string, error) {
	p, err := get(message, header, &wrappers.StringValue{})
	if err != nil {
		return "", err
	}
	return p.(*wrappers.StringValue).Value, nil
}

func get(message *BurstMessage, header Headers, m proto2.Message) (proto2.Message, error) {
	a := message.Header[int32(header)]
	if a == nil {
		return nil, NotFound
	}

	_ = a.UnmarshalTo(m)
	return m, nil
}

var NotFound = errors.New("not found")
