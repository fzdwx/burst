package common

import (
	"github.com/fatih/color"
	log "github.com/sirupsen/logrus"
)

// IsDebug check current log level is log.DebugLevel ?
func IsDebug() bool {
	return log.GetLevel() == log.DebugLevel
}

func WrapGreen(format string, a ...interface{}) string {
	return wrap(color.GreenString, format, a)
}

func WrapBlue(format string, a ...interface{}) string {
	return wrap(color.BlueString, format, a)
}

func WrapCyan(format string, a ...interface{}) string {
	return wrap(color.CyanString, format, a)
}

func WrapRed(format string, a ...interface{}) string {
	return wrap(color.RedString, format, a)
}

func wrap(f func(format string, a ...interface{}) string, format string, a ...interface{}) string {
	if len(a) > 1 {
		return "[" + f(format, a) + "]"
	}
	return "[ " + f(format) + " ]"
}
