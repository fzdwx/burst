package main

import (
	"flag"
	"github.com/fzdwx/burst/burst-client/protocol"
	ws "github.com/fzdwx/burst/burst-client/ws"
	log "github.com/sirupsen/logrus"
	"net/url"
	"os"
	"strings"
)

var addr = flag.String("addr", "localhost:8080", "http service address")
var token = flag.String("t", "fda14ac64938420b873226127c5578b1", "connect token")
var debug = flag.Bool("d", true, "log level use debug")

func init() {
	flag.Parse()
	if strings.Compare(*token, "null") == 0 {
		log.Fatal("token is null")
		os.Exit(1)
	}

	log.SetFormatter(&log.JSONFormatter{
		FieldMap: log.FieldMap{
			log.FieldKeyTime:  "time",
			log.FieldKeyLevel: "level",
			log.FieldKeyMsg:   "message",
		},
		TimestampFormat: "2006-01-02 15:04:05 111",
		//PrettyPrint:     false,
	})

	if *debug {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}
}

func main() {
	u := url.URL{Scheme: "ws", Host: *addr, Path: "/connect", RawQuery: "token=" + *token}
	client, err := ws.Connect(u)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer client.Close()

	client.MountBinaryHandler(func(data []byte, ws ws.Client) {
		burstMessage, err := protocol.Decode(data)
		if err != nil {
			log.Error(err)
			return
		}

		switch burstMessage.Type {
		case protocol.BurstType_INIT:
			handlerInit(burstMessage, ws)
		}
	})

	down := make(chan struct{})
	go func() {
		defer close(down)
		client.StartReadMessage()
	}()

	for {
		select {
		case <-down:
			return
		}
	}
}

func handlerInit(message *protocol.BurstMessage, client ws.Client) {
	err := protocol.GetError(message)
	if err != nil {
		log.Error("init error ", err)
		client.Close()
		os.Exit(1)
	}

	ports, err := protocol.GetPorts(message)
	if err != nil {
		log.Error("init get ports error ", err)
		client.Close()
		os.Exit(1)
	}

	client.SetPorts(ports.GetPorts())
	log.Info("init success ", client.Ports())
}
