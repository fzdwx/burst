package cmd

import (
	"github.com/fzdwx/burst/internal/logx"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	zlog "github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/zero-contrib/logx/zerologx"
	"os"
)

var (
	cmd = &cobra.Command{
		Use: "burst",
	}

	logFile  = ""
	logLevel = "debug"
)

func init() {
	cmd.PersistentFlags().StringVarP(&logFile, "log", "l", "", "the log file path, e.g: ./server.log")
	cmd.PersistentFlags().StringVarP(&logLevel, "level", "v", "debug", "the log level, e.g: debug, info, warn, error, fatal, panic")

	cmd.AddCommand(serve)
	cmd.AddCommand(exportCmd)
}

func loadLog() {
	if logFile != "" {
		out, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			logx.Fatal().Err(err).Msg("open log file fail")
		}
		logx.InitLogger(out)
	} else {
		logx.InitLogger(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "2006/01/02 - 15:04:05"})
		logx.Debug().Msg("log file is empty, use console writer")
	}

	logx.UseLogLevel(logx.GetLogLevel(logLevel))
	zlog.SetWriter(zerologx.NewZeroLogWriter(logx.GetLog()))
}

func Execute() error {
	return cmd.Execute()
}
