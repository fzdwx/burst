package main

import (
	"flag"
	"github.com/fzdwx/burst"
	"github.com/fzdwx/burst/client"
	"github.com/fzdwx/burst/client/command"
	"github.com/fzdwx/burst/client/handler"
	"github.com/fzdwx/burst/pkg/logx"
	"github.com/fzdwx/burst/pkg/wsx"
	"github.com/rs/zerolog"
	"github.com/zeromicro/go-zero/core/conf"
	zlog "github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/zero-contrib/logx/zerologx"
	"io/ioutil"
	"net/http"
	"os"
)

var (
	cc        = flag.String("c", "client.yaml", "the config file path")
	tokenFlag = flag.String("t", "", "the access token")
	logFile   = flag.String("l", "", "the log file path, e.g: ./client.log")
	cConfig   client.Config
	token     string
)

func init() {
	flag.Parse()

	conf.MustLoad(*cc, &cConfig)

	if *logFile != "" {
		out, err := os.OpenFile(*logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			logx.Fatal().Err(err).Msg("open log file fail")
		}
		logx.InitLogger(out)
	} else {
		logx.InitLogger(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "2006/01/02 - 15:04:05"})
	}

	logx.UseLogLevel(logx.GetLogLevel(cConfig.LogLevel))
	zlog.SetWriter(zerologx.NewZeroLogWriter(logx.GetLog()))

	serverAddr := burst.FormatAddr(cConfig.Server.Host, cConfig.Server.Port)
	logx.Info().Msgf("server addr %s", serverAddr)

	if *tokenFlag == burst.EmptyStr {
		token = generateToken(serverAddr)
	} else {
		token = *tokenFlag
	}

	logx.Info().Msgf("token: %s", token)
}

func main() {
	c := client.NewClient(token, cConfig)

	c.Connect(func(wsx *wsx.Wsx) {
		wsx.MountBinaryFunc(handler.Dispatch(c))
	})

	c.ReaderCommand(command.Dispatch)
}

func generateToken(serverAddr string) string {
	logx.Info().Msg("token is empty,call server generate")
	response, err := http.Get("http://" + serverAddr + "/user/auth")
	if err != nil {
		logx.Fatal().Err(err).Msg("call server generate token")
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logx.Fatal().Err(err).Msg("read server response fail")
	}

	logx.Info().Msg("generate token success")
	return string(body)
}
