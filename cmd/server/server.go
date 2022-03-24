package main

import (
	"github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/container"
	log "github.com/sirupsen/logrus"
)

func main() {
	server, err := container.NewContainer()
	if err != nil {
		log.Panic(err)
	}

	if err := server.MigrateDB(); err != nil {
		log.Panic(err)
	}

	if err := server.Start(); err != nil {
		log.Panic(err)
	}
}
