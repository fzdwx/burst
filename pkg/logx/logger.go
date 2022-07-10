package logx

import (
	"context"
	"fmt"
	"github.com/rs/zerolog"
	"io"
	"os"
)

var log = zerolog.New(os.Stderr).With().Timestamp().Logger()

func GetLogLevel(str string) zerolog.Level {
	switch str {
	case "debug":
		return zerolog.DebugLevel
	case "info":
		return zerolog.InfoLevel
	case "warn":
		return zerolog.WarnLevel
	case "error":
		return zerolog.WarnLevel
	case "fatal":
		return zerolog.FatalLevel
	case "panic":
		return zerolog.PanicLevel
	}
	return zerolog.DebugLevel
}

func UseLogLevel(l zerolog.Level) {
	zerolog.SetGlobalLevel(l)
}

func InitLogger(writer ...io.Writer) {
	zerolog.TimeFieldFormat = "2006/01/02 - 15:04:05"
	multi := zerolog.MultiLevelWriter(writer...)
	log = zerolog.New(multi).With().Timestamp().Logger()
}

func init() {
	//consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "2006/01/02 - 15:04:05"}
	//consoleWriter.FormatLevel = func(i interface{}) string {
	//	return fmt.Sprintf("|%s %-6s%s|", getLevel(i), strings.ToUpper(i.(string)), colorx.ResetRaw)
	//}
	//consoleWriter.FormatMessage = func(i interface{}) string {
	//	return fmt.Sprintf("%s", i)
	//}
	//consoleWriter.FormatFieldName = func(i interface{}) string {
	//	return fmt.Sprintf("%s=", colorx.Colorize(i, colorx.ColorDarkGray))
	//}
	//consoleWriter.FormatFieldValue = func(i interface{}) string {
	//	return fmt.Sprintf("%s", i)
	//}
	InitLogger(os.Stdout)
}

func GetLog() zerolog.Logger {
	return log
}

func EnableDebug() bool {
	return zerolog.GlobalLevel() == zerolog.DebugLevel
}

// Output duplicates the global logger and sets w as its output.
func Output(w io.Writer) zerolog.Logger {
	return log.Output(w)
}

// With creates a child logger with the field added to its context.
func With() zerolog.Context {
	return log.With()
}

// Level creates a child logger with the minimum accepted level set to level.
func Level(level zerolog.Level) zerolog.Logger {
	return log.Level(level)
}

// Sample returns a logger with the s sampler.
func Sample(s zerolog.Sampler) zerolog.Logger {
	return log.Sample(s)
}

// Hook returns a logger with the h Hook.
func Hook(h zerolog.Hook) zerolog.Logger {
	return log.Hook(h)
}

// Err starts a new message with error level with err as a field if not nil or
// with info level if err is nil.
//
// You must call Msg on the returned event in order to send the event.
func Err(err error) *zerolog.Event {
	return log.Err(err)
}

// Trace starts a new message with trace level.
//
// You must call Msg on the returned event in order to send the event.
func Trace() *zerolog.Event {
	return log.Trace()
}

// Debug starts a new message with debug level.
//
// You must call Msg on the returned event in order to send the event.
func Debug() *zerolog.Event {
	return log.Debug()
}

// Info starts a new message with info level.
//
// You must call Msg on the returned event in order to send the event.
func Info() *zerolog.Event {
	return log.Info()
}

// Warn starts a new message with warn level.
//
// You must call Msg on the returned event in order to send the event.
func Warn() *zerolog.Event {
	return log.Warn()
}

// Error starts a new message with error level.
//
// You must call Msg on the returned event in order to send the event.
func Error() *zerolog.Event {
	return log.Error()
}

// Fatal starts a new message with fatal level. The os.Exit(1) function
// is called by the Msg method.
//
// You must call Msg on the returned event in order to send the event.
func Fatal() *zerolog.Event {
	return log.Fatal()
}

// Panic starts a new message with panic level. The message is also sent
// to the panic function.
//
// You must call Msg on the returned event in order to send the event.
func Panic() *zerolog.Event {
	return log.Panic()
}

// WithLevel starts a new message with level.
//
// You must call Msg on the returned event in order to send the event.
func WithLevel(level zerolog.Level) *zerolog.Event {
	return log.WithLevel(level)
}

// Log starts a new message with no level. Setting zerolog.GlobalLevel to
// zerolog.Disabled will still disable events produced by this method.
//
// You must call Msg on the returned event in order to send the event.
func Log() *zerolog.Event {
	return log.Log()
}

// Print sends a log event using debug level and no extra field.
// Arguments are handled in the manner of fmt.Print.
func Print(v ...interface{}) {
	log.Debug().CallerSkipFrame(1).Msg(fmt.Sprint(v...))
}

// Printf sends a log event using debug level and no extra field.
// Arguments are handled in the manner of fmt.Printf.
func Printf(format string, v ...interface{}) {
	log.Debug().CallerSkipFrame(1).Msgf(format, v...)
}

// Ctx returns the Logger associated with the ctx. If no logger
// is associated, a disabled logger is returned.
func Ctx(ctx context.Context) *zerolog.Logger {
	return zerolog.Ctx(ctx)
}
