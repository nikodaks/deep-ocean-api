package server

import (
	"context"
	"deep-ocean-api/internal/config"
	"net/http"
	"strconv"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(cfg *config.Config, handler http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:           ":" + strconv.Itoa(cfg.Port),
			Handler:        handler,
			ReadTimeout:    cfg.HTTPReadTimeout,
			WriteTimeout:   cfg.HTTPWriteTimeout,
			MaxHeaderBytes: cfg.HTTPMaxHeaderBytes << 20,
		},
	}
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
