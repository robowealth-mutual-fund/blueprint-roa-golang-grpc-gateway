package httpServer

import (
	"log"
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

	log.Println("starting HTTP/REST gateway..." + strconv.Itoa(s.Config.HttpPort))
	return srv.ListenAndServe()
}
