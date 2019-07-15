package http

import (
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
)

type Server struct {
	ln      net.Listener
	Handler http.Handler
	Addr    string
}

const DefaultAddr = ":8080"

func NewServer() *Server {
	return &Server{
		Addr: DefaultAddr,
	}
}

func (s *Server) Open() error {
	ln, err := net.Listen("tcp", s.Addr)
	if err != nil {
		return err
	}
	s.ln = ln

	// Start HTTP server.
	go func() { http.Serve(s.ln, s.Handler) }()

	return nil
}

func (s *Server) Close() error {
	if s.ln != nil {
		s.ln.Close()
	}
	return nil
}

func (s *Server) ListenAndServe() {
	log.Printf("Starting server on %s...\n", s.Addr)
	s.Open()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}
