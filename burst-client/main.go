package main

import (
	"errors"
	"flag"
	"github.com/fzdwx/burst/burst-client/protocol"
	ws "github.com/fzdwx/burst/burst-client/ws"
	log "github.com/sirupsen/logrus"
	"net/url"
	"os"
	"strconv"
	"strings"
)

var (
	serverIp   = flag.String("addr", "localhost", "serverIp")
	serverPort = flag.Int("serverPort", 8080, "server serverPort")
	token      = flag.String("t", "de73df98abad4117a53fa2dfa27da7ac", "your key, you can get it from server")
	usage      = flag.Bool("h", false, "help")
	debug      = flag.Bool("d", true, "log level use debug")
	host       string
)

func init() {
	flag.Parse()
	if *usage {
		flag.Usage()
		os.Exit(0)
	}

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

	log.Info("log level is ", log.GetLevel())
	log.Info("server ip:", *serverIp)
	log.Info("server port:", *serverPort)
	host = *serverIp + ":" + strconv.Itoa(*serverPort)
}

func main() {
	u := url.URL{Scheme: "ws", Host: host, Path: "/connect", RawQuery: "token=" + *token}
	client, err := ws.Connect(u)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer client.Close()

	client.MountBinaryHandler(func(data []byte, ws *ws.Client) {
		burstMessage, err := protocol.Decode(data)
		if err != nil {
			log.Error(err)
			return
		}

		switch burstMessage.Type {
		case protocol.BurstType_INIT:
			handlerInit(burstMessage, ws)
		case protocol.BurstType_USER_CONNECT:
			handlerUserConnect(burstMessage, ws)
		case protocol.BurstType_FORWARD_DATA:
			handlerForwardData(burstMessage, ws)
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

func handlerInit(message *protocol.BurstMessage, client *ws.Client) {
	err := protocol.GetError(message)
	if err != nil {
		client.Over(errors.New("init error " + err.Error()))
	}

	ports, err := protocol.GetPorts(message)
	if err != nil {
		client.Over(errors.New("init get ports error " + err.Error()))
	}

	client.SetPorts(ports.GetPorts())
	log.Info("init success ", client.Ports())
}

func handlerUserConnect(message *protocol.BurstMessage, client *ws.Client) {
	serverExportPort, err := protocol.GetServerExportPort(message)
	if err != nil {
		log.Error("parse server export port error ", err)
		return
	}

	localPort, ok := client.LocalPort(serverExportPort)
	if !ok {
		log.Error("local port not found ", serverExportPort)
		return
	}

	userConnectId, err := protocol.GetUserConnectId(message)
	if err != nil {
		log.Error("parse user connect id error ", err)
		return
	}

	userConnForward, err := ws.NewUserConn(localPort, userConnectId)
	if err != nil {
		log.Error("local port connect error ", err)
		return
	}

	// step 5 [forwarded to the server], and then forwarded to a specific user
	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Error("forward to server: recover ", err)
			}
		}()
		userConnForward.StartForwardToServer(client)
	}()
}

func handlerForwardData(message *protocol.BurstMessage, client *ws.Client) {
	// step 4 [forward to local port]
	ws.Fw.ToLocal(message)
}
