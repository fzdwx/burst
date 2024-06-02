package cmd

import (
	"github.com/fzdwx/burst/internal"
	"github.com/fzdwx/burst/internal/client"
	"github.com/fzdwx/burst/internal/client/command"
	"github.com/fzdwx/burst/internal/client/handler"
	"github.com/fzdwx/burst/internal/logx"
	"github.com/fzdwx/burst/internal/wsx"
	"github.com/spf13/cobra"
	"io/ioutil"
	"net/http"
	"strings"
)

var (
	clientCmd = &cobra.Command{
		Use: "client",
		Run: func(cmd *cobra.Command, args []string) {
			loadLog()
			var cConfig = client.Config{
				Server: struct {
					Port int    `json:",default=9999"`
					Host string `json:",required=true"`
				}(struct {
					Port int
					Host string
				}{
					Port: serverPort,
					Host: serverHost,
				}),
			}

			serverAddr := internal.FormatAddr(cConfig.Server.Host, cConfig.Server.Port)
			logx.Info().Msgf("server addr %s", serverAddr)
			if token == internal.EmptyStr {
				token = generateToken(serverAddr)
			}

			c := client.NewClient(token, cConfig)

			c.Connect(func(wsx *wsx.Wsx) {
				wsx.MountBinaryFunc(handler.Dispatch(c))
			})

			if interactive {
				c.ReaderCommand(command.Dispatch, command.Autocomplete)
				return
			}

			userCommand := strings.Join(args, " ")
			command.Dispatch(userCommand, c)
			select {}
		},
	}
	token       = ""
	serverPort  = 9999
	serverHost  = ""
	interactive = false
)

func init() {
	clientCmd.PersistentFlags().StringVarP(&token, "token", "t", "", "the access token, if not set, will generate a new one")
	clientCmd.PersistentFlags().IntVarP(&serverPort, "port", "p", 9999, "the server port")
	clientCmd.PersistentFlags().StringVarP(&serverHost, "host", "s", "0.0.0.0", "the server host")
	clientCmd.PersistentFlags().BoolVarP(&interactive, "interactive", "i", false, "interactive mode")
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
