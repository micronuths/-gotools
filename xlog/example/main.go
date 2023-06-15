package main

import (
	"fmt"

	log "github.com/micronuths/gotools/xlog"
	"github.com/micronuths/gotools/xlog/config"
)

func main() {
	for i := 0; i < 1; i++ {
		log.Infof("Hi %s, system is starting up ...", "paas-bot")
		log.Info("check-info", config.Data{
			"info": "something",
		})

		log.Debug("check-info", config.Data{
			"info": "something",
		})

		log.Warn("failed-to-do-somthing", config.Data{
			"info": "something",
		})

		err := fmt.Errorf("This is an error")
		log.Error("failed-to-do-somthing", err)

		log.Info("shutting-down")
	}
}
