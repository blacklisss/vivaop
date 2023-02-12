package routergin

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
	"vivaop/internal/infrastructure/token"
	"vivaop/internal/usecases/app/repos/organizationrepo"
	"vivaop/internal/util"

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
	fmt.Println(organization)

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
	ID               uuid.UUID `json:"id"  binding:"required"`
	OwnerID          uuid.UUID `json:"owner_id" binding:"required"`
	CountryID        int32     `json:"country_id" binding:"required"`
	Verified         bool      `json:"verified"`
	Name             string    `json:"name" binding:"required"`
	RegistrationCode string    `json:"registration_code" binding:"required"`
	RegistrationDate string    `json:"registration_date" binding:"required"`
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

	organizations, err := router.hs.ListOwnerOrganization(ctx, authPayload.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, organizations)
}

type searchOrganization struct {
	Query string `json:"query" binding:"required"`
}

func (router *RouterGin) SearchOrganizations(ctx *gin.Context) {
	var req searchOrganization
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	organizations, err := router.hs.SearchOrganizations(ctx, req.Query)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, organizations)
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

type uploadOrganizationParams struct {
	ID               string `form:"id"  binding:"required"`
	OwnerID          string `form:"owner_id" binding:"required"`
	CountryID        string `form:"country_id" binding:"required"`
	Name             string `form:"name" binding:"required"`
	RegistrationCode string `form:"registration_code" binding:"required"`
	RegistrationDate string `form:"registration_date" binding:"required"`
}

func (router *RouterGin) UploadRegistration(ctx *gin.Context) {
	var req uploadOrganizationParams
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	file, err := ctx.FormFile("registration_image")
	if err != nil {
		if err.Error() == "http: no such file" {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	if req.OwnerID != authPayload.ID.String() {
		ctx.JSON(http.StatusForbidden, errorResponse(fmt.Errorf("%s", "not your organization")))
		return
	}

	ownerArr := strings.Split(req.OwnerID, "-")
	tmpPath := strings.Replace(req.ID, "-", "", -1)

	uploadPath := router.config.UploadRegPath + string(os.PathSeparator) + ownerArr[0][:2] + string(os.PathSeparator) + ownerArr[0][2:4] + string(os.PathSeparator) + ownerArr[0][4:6] + string(os.PathSeparator) + tmpPath + string(os.PathSeparator)

	ok, err := util.DirExists(uploadPath)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if !ok {
		err = util.MakeDir(uploadPath)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
	}

	uploadPath += file.Filename
	uploadURL := router.config.UploadRegURL + string(os.PathSeparator) + ownerArr[0][:2] + string(os.PathSeparator) + ownerArr[0][2:4] + string(os.PathSeparator) + ownerArr[0][4:6] + string(os.PathSeparator) + tmpPath + string(os.PathSeparator) + file.Filename

	err = ctx.SaveUploadedFile(file, uploadPath)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	id := uuid.MustParse(req.ID)

	organization, err := router.hs.UploadRegistration(ctx, &organizationrepo.UploadOrganizationParams{
		ID:        id,
		UploadURL: uploadURL,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, organization)
}
