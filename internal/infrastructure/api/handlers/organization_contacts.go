package handlers

import (
	"context"
	"fmt"
	"time"
	"vivaop/internal/entities/organization_contact_entity"
	"vivaop/internal/usecases/app/repos/organization_contact_repo"

	"github.com/google/uuid"
)

type OrganizationContact struct {
	ID             uuid.UUID
	OrganizationID uuid.UUID
	Name           string
	Phone          string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      time.Time
}

func (rt *Handlers) CreateOrganizationContact(ctx context.Context, arg *organization_contact_repo.CreateOrganizationContactParams) (*OrganizationContact, error) {
	oc, err := rt.ocs.CreateOrganizationContact(ctx, arg)
	if err != nil {
		return &OrganizationContact{}, fmt.Errorf("error when creating: %w", err)
	}

	return &OrganizationContact{
		ID:             oc.ID,
		OrganizationID: oc.OrganizationID,
		Name:           oc.Name,
		Phone:          oc.Phone,
		CreatedAt:      oc.CreatedAt,
		UpdatedAt:      oc.UpdatedAt,
		DeletedAt:      oc.DeletedAt,
	}, nil
}

func (rt *Handlers) DeleteOrganizationContact(ctx context.Context, id uuid.UUID) (*OrganizationContact, error) {
	oc, err := rt.ocs.DeleteOrganizationContact(ctx, id)
	if err != nil {
		return &OrganizationContact{}, fmt.Errorf("error when deleting: %w", err)
	}

	return &OrganizationContact{
		ID:             oc.ID,
		OrganizationID: oc.OrganizationID,
		Name:           oc.Name,
		Phone:          oc.Phone,
		CreatedAt:      oc.CreatedAt,
		UpdatedAt:      oc.UpdatedAt,
		DeletedAt:      oc.DeletedAt,
	}, nil
}

func (rt *Handlers) GetOrganizationContact(ctx context.Context, id uuid.UUID) (*OrganizationContact, error) {
	oc, err := rt.ocs.GetOrganizationContact(ctx, id)
	if err != nil {
		return &OrganizationContact{}, fmt.Errorf("error when deleting: %w", err)
	}

	return &OrganizationContact{
		ID:             oc.ID,
		OrganizationID: oc.OrganizationID,
		Name:           oc.Name,
		Phone:          oc.Phone,
		CreatedAt:      oc.CreatedAt,
		UpdatedAt:      oc.UpdatedAt,
		DeletedAt:      oc.DeletedAt,
	}, nil
}

func (rt *Handlers) ListOrganizationContacts(ctx context.Context, organizationID uuid.UUID) ([]*organization_contact_entity.OrganizationContact, error) {
	oc, err := rt.ocs.ListOrganizationContacts(ctx, organizationID)
	if err != nil {
		return nil, fmt.Errorf("error when getting organization contacts: %w", err)
	}

	return oc, nil
}
