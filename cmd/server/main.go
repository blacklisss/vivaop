package main

import (
	"context"
	"database/sql"
	"os"
	"os/signal"
	"vivaop/internal/infrastructure/api/handlers"
	"vivaop/internal/infrastructure/api/routergin"
	pgstore "vivaop/internal/infrastructure/db/pgstore/sqlc"
	srv "vivaop/internal/infrastructure/server"
	"vivaop/internal/infrastructure/token"
	"vivaop/internal/usecases/app/repos/countryrepo"
	"vivaop/internal/usecases/app/repos/organization_contact_repo"
	"vivaop/internal/usecases/app/repos/organizationrepo"
	"vivaop/internal/usecases/app/repos/sessionrepo"
	"vivaop/internal/usecases/app/repos/userrepo"
	"vivaop/internal/util"

	_ "github.com/lib/pq"

	"github.com/rs/zerolog/log"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config")
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to db")
	}

	store := pgstore.NewStore(conn)

	cs := countryrepo.NewCountries(store)
	us := userrepo.NewUsers(store)
	ss := sessionrepo.NewSession(store)
	osstore := organizationrepo.NewOrganizationStore(store)
	ocstore := organization_contact_repo.NewOrganizationContactsStore(store)

	hs := handlers.NewHandlers(us, cs, ss, osstore, ocstore)

	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey) // config.TokenSymmetricKey
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create token maker")
	}

	router, err := routergin.NewRouterGin(&config, hs, tokenMaker)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create server engine")
	}

	server := runGinServer(&config, store, router)

	<-ctx.Done()

	server.Stop()
	cancel()
	conn.Close()

	log.Print("Exit")
}

func runGinServer(config *util.Config, store pgstore.Store, h *routergin.RouterGin) *srv.Server {
	server, err := srv.NewServer(config, store, h)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create server")
	}

	server.Start()

	return server
}
