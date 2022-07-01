package main

import (
	"flag"
	"fmt"
	"github.com/fzdwx/burst/pkg/ginx"
	"github.com/fzdwx/burst/pkg/logx"
	"github.com/fzdwx/burst/pkg/wsx"
	"github.com/fzdwx/burst/server"
	"github.com/gorilla/websocket"
	"github.com/jinzhu/configor"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)

var c = flag.String("c", "server.yml", "the config file path")
var sConfig = server.Config{}

func init() {
	flag.Parse()

	err := configor.Load(&sConfig, *c)
	if err != nil {
		logx.Fatal().Msg(err.Error())
	}

}

func main() {
	r := setupRouter()

	//fmt.Println(cf.Port)
	// Listen and Server in 0.0.0.0:8080
	r.Run(sConfig.Addr)
}

func setupRouter() *gin.Engine {

	// Disable Console Color
	logx.UseDebugLevel()
	r := ginx.Classic()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		logx.Info().Msg("hello world")
		logx.Warn().Msg("ttttttttttttt")
		logx.Error().Msg("qqqqqqqqqqq")
		c.String(http.StatusOK, "pong")
	})

	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	r.GET("/ws", func(context *gin.Context) {
		conn, err := upgrader.Upgrade(context.Writer, context.Request, nil)
		if err != nil {
			fmt.Println(err)
			return
		}
		ws := wsx.NewClassicWsx(conn)

		go ws.StartReading(time.Second * 20)
		go ws.StartWriteHandler(time.Second * 5)

		ws.WriteText("hello world")
	})

	return r
}