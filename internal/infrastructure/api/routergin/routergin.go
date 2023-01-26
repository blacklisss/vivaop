package routergin

import (
	"vivaop/internal/infrastructure/api/handlers"
	"vivaop/internal/infrastructure/token"
	"vivaop/internal/util"

	"github.com/gin-gonic/gin"
)

type RouterGin struct {
	*gin.Engine
	hs         *handlers.Handlers
	tokenMaker token.Maker
	config     *util.Config
}

func NewRouterGin(config *util.Config, hs *handlers.Handlers, tokenMaker token.Maker) (*RouterGin, error) {
	r := gin.Default()
	ret := &RouterGin{
		hs: hs,
	}

	ret.tokenMaker = tokenMaker
	ret.config = config
	ret.setupRouter(r)

	ret.Engine = r
	return ret, nil
}

func (router *RouterGin) setupRouter(r *gin.Engine) {
	r.POST("/countries", router.CreateCountry)
	r.GET("/countries", router.ShowCountries)

	r.POST("/users", router.CreateUser)
	r.POST("/users/login", router.loginUser)
	r.POST("/tokens/renew_access", router.renewAccessToken)

	authRoutes := r.Group("/").Use(authMiddleware(router.tokenMaker))
	authRoutes.POST("/organizations/create", router.CreateOrganization)
	/*

		authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))
		authRoutes.POST("/accounts", server.createAccount)
		authRoutes.GET("/accounts/:id", server.getAccount)
		authRoutes.GET("/accounts", server.listAccounts)

		authRoutes.POST("/transfers", server.createTransfer)*/
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
