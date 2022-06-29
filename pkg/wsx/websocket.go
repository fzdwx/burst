package wsx

import (
	"github.com/fzdwx/burst"
	"github.com/gorilla/websocket"
	"strings"
	"time"
)

type (
	// Wsx websocket
	Wsx struct {
		conn *websocket.Conn

		onBinary func(bytes []byte)

		onText func(text string)

		onClose func(err error)

		writeTextChan   chan []byte
		writeBinaryChan chan []byte
	}
)

func NewWsx(conn *websocket.Conn, onPeerClose func(code int, text string) error) *Wsx {
	conn.SetCloseHandler(onPeerClose)
	return &Wsx{
		conn:            conn,
		writeTextChan:   make(chan []byte),
		writeBinaryChan: make(chan []byte),
	}
}

func NewClassicWsx(conn *websocket.Conn) *Wsx {
	return NewWsx(conn, nil)
}

func (w *Wsx) MountCloseFunc(onClose func(err error)) *Wsx {
	if onClose == nil {
		burst.Over("onClose is nil")
	}

	w.onClose = onClose
	return w
}

func (w *Wsx) MountBinaryFunc(onBinary func(bytes []byte)) *Wsx {
	if onBinary == nil {
		burst.Over("onBinary is nil")
	}

	w.onBinary = onBinary
	return w
}

func (w *Wsx) MountTextFunc(onText func(text string)) *Wsx {
	if onText == nil {
		burst.Over("onText is nil")
	}

	w.onText = onText
	return w
}

func (w *Wsx) WriteText(text string) {
	w.writeTextChan <- []byte(text)
}

func (w *Wsx) WriteBinary(bytes []byte) {
	w.writeBinaryChan <- bytes
}

func (w *Wsx) StartReading(pongWait time.Duration) {
	defer w.Close()
	_ = w.conn.SetReadDeadline(time.Now().Add(pongWait))
	w.conn.SetPongHandler(func(appData string) error {
		_ = w.conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {
		t, m, err := w.conn.ReadMessage()
		if err != nil {
			w.printWebSocketError("read", err)
			return
		}
		if t == websocket.BinaryMessage {
			w.onBinary(m)
		}

		if t == websocket.TextMessage {
			w.onText(string(m))
		}
		// todo log
		//w.debug().Interface("event", fmt.Sprintf("%T", incoming)).Msg("WebSocket Receive")
		//w.read <- ClientMessage{Info: w.info, Incoming: incoming}
	}
}

// startWriteHandler starts the write loop. The method has the following tasks:
// * ping the client in the interval provided as parameter
// * write messages send by the channel to the client
// * on errors exit the loop
func (w *Wsx) startWriteHandler(pingPeriod time.Duration) {
	pingTicker := time.NewTicker(pingPeriod)

	dead := false
	conClosed := func() {
		dead = true
		w.Close()
		pingTicker.Stop()
	}
	defer conClosed()
	defer func() {
		w.debug().Msg("WebSocket Done")
	}()
	for {
		select {
		case reason := <-w.info.Close:
			if reason == CloseDone {
				return
			} else {
				_ = w.conn.CloseHandler()(websocket.CloseNormalClosure, reason)
				conClosed()
			}
		case message := <-w.info.Write:
			if dead {
				w.debug().Msg("WebSocket write on dead connection")
				continue
			}

			_ = w.conn.SetWriteDeadline(time.Now().Add(writeWait))
			typed, err := ToTypedOutgoing(message)
			w.debug().Interface("event", typed.Type).Msg("WebSocket Send")
			if err != nil {
				w.debug().Err(err).Msg("could not get typed message, exiting connection.")
				conClosed()
				continue
			}

			if room, ok := message.(outgoing.Room); ok {
				w.info.RoomID = room.ID
			}

			if err := writeJSON(w.conn, typed); err != nil {
				conClosed()
				w.printWebSocketError("write", err)
			}
		case <-pingTicker.C:
			_ = w.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := ping(w.conn); err != nil {
				conClosed()
				w.printWebSocketError("ping", err)
			}
		}
	}
}

func (w *Wsx) Close() {
	if w.onClose != nil {
		w.onClose(w.conn.Close())
	}
}

func (w *Wsx) printWebSocketError(op string, err error) {
	if strings.Contains(err.Error(), "use of closed network connection") {
		return
	}
	closeError, ok := err.(*websocket.CloseError)

	if ok && closeError != nil && (closeError.Code == 1000 || closeError.Code == 1001) {
		// normal closure
		return
	}

	// todo log
	//c.debug().Str("type", op).Err(err).Msg("WebSocket Error")
}