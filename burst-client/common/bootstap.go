package common

import (
	"context"
	log "github.com/sirupsen/logrus"
)

func Run(action func(cancelFunc context.CancelFunc)) {
	ctx, cancel := context.WithCancel(context.Background())

	go action(cancel)
	select {
	case <-ctx.Done():
		log.Info("run action stop")
	}
}
