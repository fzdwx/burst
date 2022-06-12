package main

import (
	"errors"
	"flag"
	burst "github.com/fzdwx/burst/burst-client/client"
	"github.com/fzdwx/burst/burst-client/common"
	"github.com/fzdwx/burst/burst-client/protocol"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/url"
	"os"
	"strings"
)

var (
	serverIp   = flag.String("sip", "localhost", "server ip")
	serverPort = flag.Int("sp", 10086, "server serverPort")
	token      = flag.String("t", "b92a205269d94d38808c3979615245eb", "your key, you can get it from server")
	usage      = flag.Bool("h", false, "help")
	debug      = flag.Bool("d", true, "log level use debug")
	serverAddr string
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
		TimestampFormat: "2006-01-02 15:04:05",
		//PrettyPrint:     false,
	})

	if *debug {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}

	serverAddr = common.FormatToAddr(*serverIp, *serverPort)

	log.Infof("log level is [ %s ]", log.GetLevel().String())
	log.Infof("server ip: [ %s ]", *serverIp)
	log.Infof("server port: [ %d ]", *serverPort)
}

func main() {
	u := url.URL{Scheme: "ws", Host: serverAddr, Path: "/connect", RawQuery: "token=" + *token}
	client, resp, err := burst.Connect(u)
	if err != nil {
		body := resp.Body
		defer body.Close()
		data, _ := ioutil.ReadAll(body)
		log.Fatal(string(data))
	}
	defer client.Close()

	client.MountBinaryHandler(func(data []byte, client *burst.Client) {
		burstMessage, err := protocol.Decode(data)
		if err != nil {
			log.Error(err)
			return
		}

		switch burstMessage.Type {
		case protocol.BurstType_ADD_PROXY_INFO:
			handlerAddProxyInfo(burstMessage, client)
		case protocol.BurstType_USER_CONNECT:
			handlerUserConnect(burstMessage, client)
		case protocol.BurstType_FORWARD_DATA:
			handlerForwardData(burstMessage, client)
		}
	})

	down := make(chan byte)
	go func() {
		defer close(down)
		client.React()
	}()

	for {
		select {
		case <-down:
			return
		}
	}
}

// handlerAddProxyInfo 处理添加映射信息
func handlerAddProxyInfo(message *protocol.BurstMessage, client *burst.Client) {
	err := protocol.GetError(message)
	if err != nil {
		client.Over(errors.New("init error " + err.Error()))
	}

	ports, err := protocol.GetPorts(message)
	if err != nil {
		client.Over(errors.New("init get ports error " + err.Error()))
	}

	proxyInfo := ports.GetPorts()
	client.AddProxyInfo(proxyInfo)

	log.Infoln("add proxy info success")
	for serverExportPort, proxy := range proxyInfo {
		log.Infof("proxy intranet: [ %s ] to server [ %s ]", proxy.Host(), common.FormatToAddr(*serverIp, int(serverExportPort)))
	}
}

func handlerUserConnect(message *protocol.BurstMessage, client *burst.Client) {
	serverExportPort, err := protocol.GetServerExportPort(message)
	if err != nil {
		log.Error("parse server export port error ", err)
		return
	}

	proxy, ok := client.GetProxy(serverExportPort)
	if !ok {
		log.Error("local port not found ", serverExportPort)
		return
	}

	userConnectId, err := protocol.GetUserConnectId(message)
	if err != nil {
		log.Error("parse user connect id error ", err)
		return
	}

	userConnForward, err := burst.NewUserConn(proxy, userConnectId)
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
		userConnForward.React(client)
	}()
}

func handlerForwardData(message *protocol.BurstMessage, client *burst.Client) {
	// step 4 [forward to local port]
	burst.Fw.ToLocal(message)
}
