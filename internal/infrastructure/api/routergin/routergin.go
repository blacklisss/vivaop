package routergin

import (
	"vivaop/internal/infrastructure/api/handlers"
	"vivaop/internal/infrastructure/token"
	"vivaop/internal/util"

	"github.com/gin-contrib/cors"
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
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}

	r.Use(cors.New(config))

	r.POST("/countries", router.CreateCountry)
	r.GET("/countries", router.ShowCountries)

	r.POST("/users", router.CreateUser)
	r.POST("/users/login", router.loginUser)
	r.POST("/tokens/renew_access", router.renewAccessToken)

	authRoutes := r.Group("/").Use(authMiddleware(router.tokenMaker))
	authRoutes.POST("/organizations/create", router.CreateOrganization)
	authRoutes.GET("/organizations/:id", router.GetOrganization)
	authRoutes.POST("/organizations/:id", router.UpdateOrganization)
	authRoutes.DELETE("/organizations/:id", router.DeleteOrganization)
	authRoutes.GET("/organizations/my", router.ListMyOrganization)
	authRoutes.GET("/organizations/verify/:id", router.VerifyOrganization)

}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
