package srv

import (
	"context"
	"fmt"
	"net/http"
	"time"
	pgstore "vivaop/internal/infrastructure/db/pgstore/sqlc"
	"vivaop/internal/infrastructure/token"
	"vivaop/internal/util"

	"github.com/rs/zerolog/log"
)

type Server struct {
	store      pgstore.Store
	tokenMaker token.Maker
	config     *util.Config
	srv        http.Server
}

func NewServer(config *util.Config, store pgstore.Store, h http.Handler) (*Server, error) {
	s := &Server{}

	s.srv = http.Server{
		Addr:              config.HTTPServerAddress,
		Handler:           h,
		ReadTimeout:       30 * time.Second,
		WriteTimeout:      30 * time.Second,
		ReadHeaderTimeout: 30 * time.Second,
	}

	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey) // config.TokenSymmetricKey
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	s.tokenMaker = tokenMaker
	s.store = store
	s.config = config

	return s, nil
}

func (s *Server) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	_ = s.srv.Shutdown(ctx)
	cancel()
}

func (s *Server) Start() {
	// TODO: migrations
	go func() {
		err := s.srv.ListenAndServe()
		if err != nil {
			log.Fatal().
				Err(err).
				Msg("Cannot start server")
		}
	}()
}
