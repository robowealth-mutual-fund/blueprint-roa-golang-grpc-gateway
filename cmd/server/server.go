package main

import (
	"os"
	"runtime/pprof"
	"time"

	"github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/container"
	log "github.com/sirupsen/logrus"
)

func main() {
	go func() {
		time.Sleep(120 * time.Second)
		memprof, err := os.Create("mem.pprof")
		if err != nil {
			log.Fatal(err)
		}
		pprof.WriteHeapProfile(memprof)
		memprof.Close()
	}()
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
