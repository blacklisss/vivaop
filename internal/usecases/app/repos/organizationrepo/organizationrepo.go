package organizationrepo

import (
	"context"
	"database/sql"
	"fmt"
	"time"
	"vivaop/internal/entities/organizationentity"

	"github.com/google/uuid"
)

type CreateOrganizationParams struct {
	ID               uuid.UUID    `json:"id"`
	CountryID        int32        `json:"country_id"`
	Name             string       `json:"name"`
	RegistrationCode string       `json:"registration_code"`
	RegistrationDate time.Time    `json:"registration_date"`
	OwnerID          uuid.UUID    `json:"owner_id"`
	Verified         sql.NullBool `json:"verified"`
}

type GetOrganizationByOwnerParams struct {
	ID      uuid.UUID `json:"id"`
	OwnerID uuid.UUID `json:"owner_id"`
}

type UpdateOrganizationParams struct {
	ID               uuid.UUID    `json:"id"`
	CountryID        int32        `json:"country_id"`
	Name             string       `json:"name"`
	RegistrationCode string       `json:"registration_code"`
	RegistrationDate time.Time    `json:"registration_date"`
	OwnerID          uuid.UUID    `json:"owner_id"`
	Verified         sql.NullBool `json:"verified"`
}

type UploadOrganizationParams struct {
	ID        uuid.UUID
	UploadURL string
}

type OrganizationStore interface {
	AddRegistrationImage(ctx context.Context, params *UploadOrganizationParams) (*organizationentity.Organization, error)
	CreateOrganization(ctx context.Context, params *CreateOrganizationParams) (*organizationentity.Organization, error)
	GetOrganization(ctx context.Context, id uuid.UUID) (*organizationentity.Organization, error)
	GetOrganizationByOwner(ctx context.Context, params *GetOrganizationByOwnerParams) (*organizationentity.Organization, error)
	ListOwnerOrganization(ctx context.Context, ownerID uuid.UUID) ([]*organizationentity.Organization, error)
	UpdateOrganization(ctx context.Context, params *UpdateOrganizationParams) (*organizationentity.Organization, error)
	DeleteOrganization(ctx context.Context, id uuid.UUID) (*organizationentity.Organization, error)
	SearchOrganizations(ctx context.Context, query string) ([]*organizationentity.Organization, error)
	VerifyOrganization(ctx context.Context, id uuid.UUID) (*organizationentity.Organization, error)
}

type Organizations struct {
	ostore OrganizationStore
}

func NewOrganizationStore(ostore OrganizationStore) *Organizations {
	return &Organizations{
		ostore,
	}
}

func (os *Organizations) CreateOrganization(ctx context.Context, params *CreateOrganizationParams) (*organizationentity.Organization, error) {
	organization, err := os.ostore.CreateOrganization(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("create organization error: %w", err)
	}

	return organization, nil
}

func (os *Organizations) GetOrganization(ctx context.Context, id uuid.UUID) (*organizationentity.Organization, error) {
	organization, err := os.ostore.GetOrganization(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("get organization error: %w", err)
	}

	return organization, nil
}

func (os *Organizations) GetOrganizationByOwner(ctx context.Context, params *GetOrganizationByOwnerParams) (*organizationentity.Organization, error) {
	organization, err := os.ostore.GetOrganizationByOwner(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("get by owner organization error: %w", err)
	}

	return organization, nil
}

func (os *Organizations) ListOwnerOrganization(ctx context.Context, ownerID uuid.UUID) ([]*organizationentity.Organization, error) {
	organizations, err := os.ostore.ListOwnerOrganization(ctx, ownerID)
	if err != nil {
		return nil, fmt.Errorf("list organization error: %w", err)
	}

	return organizations, nil
}

func (os *Organizations) UpdateOrganization(ctx context.Context, params *UpdateOrganizationParams) (*organizationentity.Organization, error) {
	organization, err := os.ostore.UpdateOrganization(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("update organization error: %w", err)
	}

	return organization, nil
}

func (os *Organizations) DeleteOrganization(ctx context.Context, id uuid.UUID) (*organizationentity.Organization, error) {
	organization, err := os.ostore.DeleteOrganization(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("delete organization error: %w", err)
	}

	return organization, nil
}

func (os *Organizations) VerifyOrganization(ctx context.Context, id uuid.UUID) (*organizationentity.Organization, error) {
	organization, err := os.ostore.VerifyOrganization(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("verify organization error: %w", err)
	}

	return organization, nil
}

func (os *Organizations) AddRegistrationImage(ctx context.Context, params *UploadOrganizationParams) (*organizationentity.Organization, error) {
	organization, err := os.ostore.AddRegistrationImage(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("upload organization registration error: %w", err)
	}

	return organization, nil
}

func (os *Organizations) SearchOrganizations(ctx context.Context, query string) ([]*organizationentity.Organization, error) {
	organizations, err := os.ostore.SearchOrganizations(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("list organization error: %w", err)
	}

	return organizations, nil
}
