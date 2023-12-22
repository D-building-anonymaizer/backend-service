package backend

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, ip string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ip + ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 8192,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
