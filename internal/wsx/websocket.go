package wsx

import (
	"github.com/fzdwx/burst/internal"
	"github.com/fzdwx/burst/internal/logx"
	"github.com/gorilla/websocket"
	"github.com/rs/xid"
	"github.com/rs/zerolog"
	"net"
	"strings"
	"sync"
	"time"
)

type (
	// Wsx websocket
	Wsx struct {
		conn    *websocket.Conn
		Id      xid.ID
		IdStr   string
		Addr    net.IP
		AddrStr string

		onBinary func(bytes []byte)

		onText func(text string)

		// onClose actively close the connection to the other party,
		// Usually do some function of resource release.
		onClose func(err error)

		writeTextChan   chan []byte
		writeBinaryChan chan []byte

		cleaned bool
		lock    sync.Mutex
	}
)

var (
	defaultBinary = func(bytes []byte) {

	}

	defaultText = func(text string) {

	}
)

const (
	writeWait = 2 * time.Second
)

func NewWsx(conn *websocket.Conn, onPeerClose func(code int, text string) error) *Wsx {
	conn.SetCloseHandler(onPeerClose)
	id := xid.New()
	ip := conn.RemoteAddr().(*net.TCPAddr).IP

	return &Wsx{
		conn:            conn,
		Id:              id,
		IdStr:           id.String(),
		Addr:            ip,
		AddrStr:         ip.String(),
		writeTextChan:   make(chan []byte),
		writeBinaryChan: make(chan []byte),
		onBinary:        defaultBinary,
		onText:          defaultText,
		cleaned:         false,
		lock:            sync.Mutex{},
	}
}

func NewClassicWsx(conn *websocket.Conn) *Wsx {
	return NewWsx(conn, nil)
}

func (w *Wsx) MountCloseFunc(onClose func(err error)) *Wsx {
	if onClose == nil {
		internal.Over("onClose is nil")
	}

	w.onClose = onClose
	return w
}

func (w *Wsx) MountBinaryFunc(onBinary func(bytes []byte)) *Wsx {
	if onBinary == nil {
		internal.Over("onBinary is nil")
	}

	w.onBinary = onBinary
	return w
}

func (w *Wsx) MountTextFunc(onText func(text string)) *Wsx {
	if onText == nil {
		internal.Over("onText is nil")
	}

	w.onText = onText
	return w
}

// WriteText write websocket.TextMessage to peer.
func (w *Wsx) WriteText(text string) {
	if text == internal.EmptyStr {
		return
	}

	w.writeTextChan <- []byte(text)
}

// WriteBinary write websocket.BinaryMessage to peer.
func (w *Wsx) WriteBinary(bytes []byte) {
	if bytes == nil {
		return
	}

	w.writeBinaryChan <- bytes
}

// StartReading starts listening on the client connection,
// distribute handlers based on message type,
// currently supports websocket.TextMessage and websocket.BinaryMessage.
func (w *Wsx) StartReading(pongWait time.Duration) {
	defer w.Close()

	if pongWait != 0 {
		_ = w.conn.SetReadDeadline(time.Now().Add(pongWait))
		w.conn.SetPongHandler(func(appData string) error {
			_ = w.conn.SetReadDeadline(time.Now().Add(pongWait))
			return nil
		})
	}

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

		w.debug().Int("read", len(m)).Msg("WebSocket Receive")
	}
}

// StartWriteHandler starts the write loop. The method has the following tasks:
//
// * ping the client in the interval provided as parameter
//
// * write messages send by the channel to the client
//
// * on errors exit the loop
func (w *Wsx) StartWriteHandler(pingPeriod time.Duration) {
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
		case message := <-w.writeTextChan:
			if message == nil {
				return
			}
			if dead {
				w.debug().Msg("WebSocket write on dead connection")
				return
			}

			_ = w.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := w.conn.WriteMessage(websocket.TextMessage, message); err != nil {
				conClosed()
				w.printWebSocketError("write text", err)
			} else {
				w.debug().Int("write", len(message)).Msg("WebSocket Send Text")
			}
		case message := <-w.writeBinaryChan:
			if message == nil {
				return
			}

			if dead {
				w.debug().Msg("WebSocket write on dead connection")
				return
			}

			_ = w.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := w.conn.WriteMessage(websocket.BinaryMessage, message); err != nil {
				conClosed()
				w.printWebSocketError("write binary", err)
			} else {
				w.debug().Int("write", len(message)).Msg("WebSocket Send Binary")
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

// Closed the websocket connection is closed ?
func (w *Wsx) Closed() bool {
	return Ping(w.conn) != nil
}

// Close the connection with the peer
func (w *Wsx) Close() {
	w.lock.Lock()
	defer w.lock.Unlock()

	if w.cleaned {
		return
	}

	defer func() {
		w.cleaned = true
	}()

	if w.onClose != nil {
		w.onClose(w.conn.Close())
	} else {
		w.conn.Close()
	}

	w.debug().Msg("WebSocket Close")

	close(w.writeBinaryChan)
	close(w.writeTextChan)
}

// websocket easy debug log
func (w *Wsx) debug() *zerolog.Event {
	return logx.Debug().Str("id", w.IdStr).Str("ip", w.AddrStr)
}

// printWebSocketError check close error
func (w *Wsx) printWebSocketError(op string, err error) {
	if strings.Contains(err.Error(), "use of closed network connection") {
		return
	}
	closeError, ok := err.(*websocket.CloseError)

	if ok && closeError != nil && (closeError.Code == 1000 || closeError.Code == 1001) {
		// normal closure
		return
	}

	w.debug().Str("type", op).Err(err).Msg("WebSocket Error")
}
