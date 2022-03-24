package grpcserver

import (
	log "github.com/sirupsen/logrus"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

// Server ...

// Start ...
func (s *Server) Start() {
	go func() {
		listen, err := net.Listen("tcp", ":"+strconv.Itoa(s.Config.Port))

		if err != nil {
			log.Panic(err)
		}

		if err := s.Server.Serve(listen); err != nil {
			log.Panic(err)
		}
	}()

	log.Info("Listening and serving GRPC on ", strconv.Itoa(s.Config.Port))
	// Gracefully Shutdown
	// Make channel listen for signals from OS
	gracefulStop := make(chan os.Signal, 1)

	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)

	<-gracefulStop
}

// Stop GracefulStop GRPC
func (s *Server) Stop() {
	s.Server.GracefulStop()
	log.Info("Server gracefully stopped")
}

// NewServer ...
