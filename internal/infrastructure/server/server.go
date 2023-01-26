package srv

import (
	"context"
	"net/http"
	"time"
	pgstore "vivaop/internal/infrastructure/db/pgstore/sqlc"
	"vivaop/internal/util"

	"github.com/rs/zerolog/log"
)

type Server struct {
	store  pgstore.Store
	config *util.Config
	Srv    http.Server
}

func NewServer(config *util.Config, store pgstore.Store, h http.Handler) (*Server, error) {
	s := &Server{}

	s.Srv = http.Server{
		Addr:              config.HTTPServerAddress,
		Handler:           h,
		ReadTimeout:       30 * time.Second,
		WriteTimeout:      30 * time.Second,
		ReadHeaderTimeout: 30 * time.Second,
	}

	s.store = store
	s.config = config

	return s, nil
}

func (s *Server) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	_ = s.Srv.Shutdown(ctx)
	cancel()
}

func (s *Server) Start() {
	// TODO: migrations
	go func() {
		err := s.Srv.ListenAndServe()
		if err != nil {
			log.Fatal().
				Err(err).
				Msg("Cannot start server")
		}
	}()
}
