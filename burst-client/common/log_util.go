package common

import log "github.com/sirupsen/logrus"

// IsDebug check current log level is log.DebugLevel ?
func IsDebug() bool {
	return log.GetLevel() == log.DebugLevel
}
