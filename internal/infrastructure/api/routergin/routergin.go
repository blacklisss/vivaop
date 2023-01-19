package routergin

import (
	"vivaop/internal/infrastructure/api/handlers"

	"github.com/gin-gonic/gin"
)

type RouterGin struct {
	*gin.Engine
	hs *handlers.Handlers
}

func NewRouterGin(hs *handlers.Handlers) *RouterGin {
	r := gin.Default()
	ret := &RouterGin{
		hs: hs,
	}

	ret.setupRouter(r)

	ret.Engine = r
	return ret
}

func (router *RouterGin) setupRouter(r *gin.Engine) {
	r.POST("/countries", router.CreateCountry)
	r.GET("/countries", router.ShowCountries)

	r.POST("/users", router.CreateUser)
	/*router.POST("/users/login", server.loginUser)
	router.POST("/tokens/renew_access", server.renewAccessToken)

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))
	authRoutes.POST("/accounts", server.createAccount)
	authRoutes.GET("/accounts/:id", server.getAccount)
	authRoutes.GET("/accounts", server.listAccounts)

	authRoutes.POST("/transfers", server.createTransfer)*/
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
