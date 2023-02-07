package routergin

import (
	"database/sql"
	"net/http"
	"strconv"
	"time"
	"vivaop/internal/infrastructure/token"
	"vivaop/internal/usecases/app/repos/organizationrepo"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type createOrganizationParams struct {
	Name             string `json:"name" binding:"required"`
	CountryID        string `json:"country_id" binding:"required"`
	RegistrationCode string `json:"registration_code" binding:"required"`
	RegistrationDate string `json:"registration_date" binding:"required"`
}

func (router *RouterGin) CreateOrganization(ctx *gin.Context) {
	var req createOrganizationParams
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	id, err := uuid.NewUUID()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c_id, err := strconv.Atoi(req.CountryID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	registartionDate, err := time.Parse(time.RFC3339, req.RegistrationDate)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	organization, err := router.hs.CreateOrganization(ctx, &organizationrepo.CreateOrganizationParams{
		ID:        id,
		Name:      req.Name,
		CountryID: int32(c_id),
		OwnerID:   authPayload.ID,
		Verified: sql.NullBool{
			Bool:  false,
			Valid: true,
		},
		RegistrationCode: req.RegistrationCode,
		RegistrationDate: registartionDate,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, organization)

}

type getOrganizationParams struct {
	ID string `uri:"id" binding:"required,uuid"`
}

func (router *RouterGin) GetOrganization(ctx *gin.Context) {
	var req getOrganizationParams
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	id := uuid.MustParse(req.ID)

	organization, err := router.hs.GetOrganization(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, organization)
}

func (router *RouterGin) GetOrganizationByOwner(ctx *gin.Context) {
	var req getOrganizationParams
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// Organization ID
	id := uuid.MustParse(req.ID)

	// Getting User ID
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	organization, err := router.hs.GetOrganizationByOwner(ctx, &organizationrepo.GetOrganizationByOwnerParams{
		ID:      id,
		OwnerID: authPayload.ID,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, organization)
}

type updateOrganizationIDParams struct {
	ID string `uri:"id" binding:"required,uuid"`
}

type updateOrganizationParams struct {
	Name      string `json:"name" binding:"required"`
	CountryID int32  `json:"country_id" binding:"required"`
	Verified  bool   `json:"verified"`
}

func (router *RouterGin) UpdateOrganization(ctx *gin.Context) {
	var req updateOrganizationIDParams
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	id := uuid.MustParse(req.ID)

	var req2 updateOrganizationParams
	if err := ctx.ShouldBindJSON(&req2); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	verified := sql.NullBool{
		Bool:  req2.Verified,
		Valid: true,
	}

	organization, err := router.hs.UpdateOrganization(ctx, &organizationrepo.UpdateOrganizationParams{
		ID:        id,
		Name:      req2.Name,
		CountryID: req2.CountryID,
		OwnerID:   authPayload.ID,
		Verified:  verified,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, organization)
}

type deleteOrganizationIDParams struct {
	ID string `uri:"id" binding:"required,uuid"`
}

func (router *RouterGin) DeleteOrganization(ctx *gin.Context) {
	var req deleteOrganizationIDParams
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	id := uuid.MustParse(req.ID)

	organization, err := router.hs.DeleteOrganization(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, organization)
}

func (router *RouterGin) ListMyOrganization(ctx *gin.Context) {

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	organization, err := router.hs.ListOwnerOrganization(ctx, authPayload.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, organization)
}

type verifyOrganizationIDParams struct {
	ID string `uri:"id" binding:"required,uuid"`
}

func (router *RouterGin) VerifyOrganization(ctx *gin.Context) {
	var req verifyOrganizationIDParams
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	id := uuid.MustParse(req.ID)

	organization, err := router.hs.VerifyOrganization(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, organization)
}
