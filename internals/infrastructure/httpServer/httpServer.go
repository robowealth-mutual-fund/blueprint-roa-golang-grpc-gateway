package httpServer

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"strconv"
)

func (s *Server) Start() error {
	srv := &http.Server{
		Addr:    ":" + strconv.Itoa(s.Config.HttpPort),
		Handler: s.HttpMux,
	}
	s.HttpMux.Handle("/", s.Server)
	s.swager()
	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	log.Info("Listening and serving HTTP on " + strconv.Itoa(s.Config.HttpPort))
	return srv.ListenAndServe()
}
