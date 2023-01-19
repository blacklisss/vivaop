package routergin

import (
	"net/http"
	"vivaop/internal/usecases/app/repos/countryrepo"

	"github.com/gin-gonic/gin"
)

type createCountryRequest struct {
	Name   string `json:"name" binding:"required"`
	NameEn string `json:"name_en" binding:"required"`
	Code   string `json:"code" binding:"required"`
}

func (router *RouterGin) CreateCountry(ctx *gin.Context) {
	var req createCountryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	country, err := router.hs.CreateCountry(ctx, countryrepo.CreateCountryParams{
		Name:   req.Name,
		NameEn: req.NameEn,
		Code:   req.Code,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, country)
}

func (router *RouterGin) ShowCountries(ctx *gin.Context) {
	countries, err := router.hs.ListCountries(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, countries)
}
