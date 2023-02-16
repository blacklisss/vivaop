package routergin

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"
	"vivaop/internal/util"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type VerifyUserEmailRequest struct {
	Token string `uri:"token" binding:"required"`
}

func (router *RouterGin) VerifyUserEmail(ctx *gin.Context) {
	var req VerifyUserEmailRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	verifiedEmail, err := router.hs.CheckEmailToken(ctx, req.Token)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if !util.CompareDate(time.Now(), verifiedEmail.ExpiredAt) {
		ctx.JSON(http.StatusBadRequest, errorResponse(fmt.Errorf("email token expired")))
		return
	}

	user, err := router.hs.VerifyUserEmail(ctx, verifiedEmail.UserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = router.hs.DeleteUserEmailVerification(ctx, verifiedEmail.UserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, user)
}
