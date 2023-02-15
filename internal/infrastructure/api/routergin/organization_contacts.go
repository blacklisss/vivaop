package routergin

import (
	"context"
	"fmt"
	"net/http"
	"vivaop/internal/infrastructure/token"
	"vivaop/internal/usecases/app/repos/organization_contact_repo"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreateOrganizationContactParams struct {
	OrganizationID string `json:"organization_id" binding:"required"`
	Name           string `json:"name" binding:"required"`
	Phone          string `json:"phone" binding:"required"`
}

func (router *RouterGin) CreateOrganizationContact(ctx *gin.Context) {
	var req CreateOrganizationContactParams
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	organizationID := uuid.MustParse(req.OrganizationID)

	id, err := uuid.NewUUID()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ok, err := router.checkRight(ctx, organizationID, authPayload.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	if !ok {
		ctx.JSON(http.StatusForbidden, errorResponse(fmt.Errorf("%s", "forbidden")))
		return
	}

	organizationContact, err := router.hs.CreateOrganizationContact(ctx, &organization_contact_repo.CreateOrganizationContactParams{
		ID:             id,
		OrganizationID: organizationID,
		Name:           req.Name,
		Phone:          req.Phone,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, organizationContact)
}

type getOrganizationContactParams struct {
	ID string `uri:"id" binding:"required,uuid"`
}

func (router *RouterGin) GetOrganizationContact(ctx *gin.Context) {
	var req getOrganizationContactParams
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	id := uuid.MustParse(req.ID)
	organizationContact, err := router.hs.GetOrganizationContact(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, organizationContact)
}

type deleteOrganizationContactParams struct {
	ID string `uri:"id" binding:"required,uuid"`
}

func (router *RouterGin) DeleteOrganizationContact(ctx *gin.Context) {
	var req deleteOrganizationContactParams
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	id := uuid.MustParse(req.ID)

	organizationContact, err := router.hs.GetOrganizationContact(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	ok, err := router.checkRight(ctx, organizationContact.OrganizationID, authPayload.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	if !ok {
		ctx.JSON(http.StatusForbidden, errorResponse(fmt.Errorf("%s", "forbidden")))
		return
	}

	organizationContact, err = router.hs.DeleteOrganizationContact(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, organizationContact)
}

type listOrganizationContactParams struct {
	OrganizationID string `uri:"id" binding:"required,uuid"`
}

func (router *RouterGin) ListOrganizationContact(ctx *gin.Context) {
	var req listOrganizationContactParams
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	OrganizationID := uuid.MustParse(req.OrganizationID)
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	ok, err := router.checkRight(ctx, OrganizationID, authPayload.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	if !ok {
		ctx.JSON(http.StatusForbidden, errorResponse(fmt.Errorf("%s", "forbidden")))
		return
	}

	organizationContacts, err := router.hs.ListOrganizationContacts(ctx, OrganizationID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, organizationContacts)
}

func (router *RouterGin) checkRight(ctx context.Context, orgId, ownerID uuid.UUID) (bool, error) {
	organization, err := router.hs.GetOrganization(ctx, orgId)
	if err != nil {
		return false, err
	}

	if organization.OwnerID != ownerID {
		return false, nil
	}

	return true, nil
}
