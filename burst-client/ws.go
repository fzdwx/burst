package main

import (
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"net/url"
)

type (
	Websocket struct {
		c        *websocket.Conn
		token    string
		onText   OnText
		onBinary OnBinary
	}

	OnText func(string, Websocket)

	OnBinary func([]byte, Websocket)
)

func Connect(url url.URL) (*Websocket, error) {
	log.Printf("start connecting to %s", url.String())
	c, _, err := websocket.DefaultDialer.Dial(url.String(), nil)

	if err != nil {
		return nil, err
	}

	return &Websocket{
		c:     c,
		token: url.Query().Get("token"),
		onText: func(s string, w Websocket) {
			log.Debugf("onText: %s\n", s)
		},
		onBinary: func(bytes []byte, w Websocket) {
			log.Debugf("onBinary: %s\n", string(bytes))
		},
	}, nil
}

func (w Websocket) Close() {
	w.c.Close()
}

func (w Websocket) StartReadMessage() {
	for {
		msgType, message, err := w.c.ReadMessage()
		if err != nil {
			log.Error("read message error: ", err)
			continue
		}

		switch msgType {
		case websocket.TextMessage:
			w.onText(string(message), w)
		case websocket.BinaryMessage:
			w.onBinary(message, w)
		}
	}
}
