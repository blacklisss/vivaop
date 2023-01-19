package ginserver

import (
	"net/http"
	"vivaop/internal/usecases/app/repos/countryrepo"

	"github.com/gin-gonic/gin"
)

type createCountryRequest struct {
	Name   string `json:"name" binding:"required,alphanum"`
	NameEn string `json:"name_en" binding:"required,alphanum"`
	Code   string `json:"full_name" binding:"required,alphanum"`
}

func (server *Server) CreateCountry(ctx *gin.Context) {
	var req createCountryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	country, err := server.handlers.CreateCountry(ctx, countryrepo.CreateCountryParams{
		Name:   req.Name,
		NameEn: req.NameEn,
		Code:   req.Code,
	})
	if err != nil {
		/*if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}*/
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, country)
}
