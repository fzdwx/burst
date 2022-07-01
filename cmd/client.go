package main

import (
	"flag"
	"github.com/fzdwx/burst"
	"github.com/fzdwx/burst/client"
	"github.com/fzdwx/burst/pkg/logx"
	"github.com/jinzhu/configor"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	cc        = flag.String("c", "client.yml", "the config file path")
	tokenFlag = flag.String("t", "", "the access token")
	cConfig   = client.Config{}
	token     string
)

func init() {
	flag.Parse()

	err := configor.Load(&cConfig, *cc)
	if err != nil {
		logx.Fatal().Msg(err.Error())
	}

	level := logx.GetLogLevel(cConfig.LogLevel)

	logx.UseLogLevel(level)

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
	client := client.NewClient(token, cConfig)

	client.Connect()

	go client.StartReading(time.Second * 20)
	go client.StartWriteHandler(time.Second * 5)
	client.WriteText("hello world")
	select {}
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