package routergin

import (
	"net/http"
	"vivaop/internal/infrastructure/token"

	"github.com/gin-gonic/gin"
)

func (router *RouterGin) CreateOrganization(ctx *gin.Context) {
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	ctx.JSON(http.StatusOK, authPayload)
}
