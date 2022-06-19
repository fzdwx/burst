package main

import (
	"context"
	"flag"
	burst "github.com/fzdwx/burst/burst-client/client"
	"github.com/fzdwx/burst/burst-client/common"
	log "github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
	"io/ioutil"
	"net/url"
	"strings"
)

var (
	serverIp   = flag.String("sip", "localhost", "server ip")
	serverPort = flag.Int("sp", 10086, "server serverPort")
	token      = flag.String("t", "5def0beba0474ac79d2cfc8b6292727c", "your key, you can get it from server")
	usage      = flag.Bool("h", false, "help")
	debug      = flag.Bool("d", true, "log level use debug")
	serverAddr string
)

func init() {
	flag.Parse()
	if *usage {
		flag.Usage()
	}

	formatter := new(prefixed.TextFormatter)
	formatter.FullTimestamp = true
	formatter.TimestampFormat = "2006-01-02 15:04:05.000"
	log.SetFormatter(formatter)

	if *debug {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}

	serverAddr = common.FormatToAddr(*serverIp, *serverPort)

	if strings.Compare(*token, "null") == 0 || strings.Compare(*token, "") == 0 {
		log.Fatal("token is null")
	}

	log.Infoln("log level:", common.WrapGreen(log.GetLevel().String()))
	log.Infoln("server address:", common.WrapGreen(serverAddr))
}

func main() {
	common.Run(func(cancelFunc context.CancelFunc) {
		u := url.URL{Scheme: "ws", Host: serverAddr, Path: "/connect", RawQuery: "token=" + *token}
		client, resp, err := burst.Connect(u)
		if err != nil {
			body := resp.Body
			defer body.Close()
			data, _ := ioutil.ReadAll(body)
			log.Fatal(string(data))
		}

		client.MountBinaryHandler(burst.HandlerBinaryData())

		go func() {
			defer func() {
				cancelFunc()
				client.Close()
			}()

			client.React()
		}()
	})
}
