package handlers

import (
	"context"
	"fmt"
	"time"
	"vivaop/internal/entities/organizationentity"
	"vivaop/internal/usecases/app/repos/organizationrepo"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type Organization struct {
	ID                uuid.UUID `json:"id"`
	OwnerID           uuid.UUID `json:"owner_id"`
	CountryID         int32     `json:"country_id"`
	Verified          bool      `json:"verified"`
	Name              string    `json:"name"`
	RegistrationCode  string    `json:"registration_code"`
	RegistrationImage string    `json:"registration_image"`
	RegistrationDate  time.Time `json:"registration_date"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	DeletedAt         time.Time `json:"deleted_at"`
}

var ErrOrganizationNotFound = errors.New("organization not found")

func (rt *Handlers) CreateOrganization(ctx context.Context, params *organizationrepo.CreateOrganizationParams) (*Organization, error) {
	o, err := rt.os.CreateOrganization(ctx, params)
	if err != nil {
		return &Organization{}, fmt.Errorf("error when creating: %w", err)
	}

	return &Organization{
		ID:               o.ID,
		Name:             o.Name,
		CountryID:        o.CountryID,
		OwnerID:          o.OwnerID,
		Verified:         o.Verified,
		RegistrationCode: o.RegistrationCode,
		RegistrationDate: o.RegistrationDate,
		CreatedAt:        o.CreatedAt,
	}, nil
}

func (rt *Handlers) GetOrganization(ctx context.Context, id uuid.UUID) (*Organization, error) {
	o, err := rt.os.GetOrganization(ctx, id)
	if err != nil {
		return &Organization{}, fmt.Errorf("error when getting: %w", err)
	}

	return &Organization{
		ID:                o.ID,
		Name:              o.Name,
		CountryID:         o.CountryID,
		OwnerID:           o.OwnerID,
		Verified:          o.Verified,
		RegistrationCode:  o.RegistrationCode,
		RegistrationDate:  o.RegistrationDate,
		RegistrationImage: o.RegistrationImage,
		CreatedAt:         o.CreatedAt,
		UpdatedAt:         o.UpdatedAt,
		DeletedAt:         o.DeletedAt,
	}, nil
}

func (rt *Handlers) GetOrganizationByOwner(ctx context.Context, params *organizationrepo.GetOrganizationByOwnerParams) (*Organization, error) {
	o, err := rt.os.GetOrganizationByOwner(ctx, params)
	if err != nil {
		return &Organization{}, fmt.Errorf("error when getting by owner: %w", err)
	}

	return &Organization{
		ID:                o.ID,
		Name:              o.Name,
		CountryID:         o.CountryID,
		OwnerID:           o.OwnerID,
		Verified:          o.Verified,
		RegistrationCode:  o.RegistrationCode,
		RegistrationDate:  o.RegistrationDate,
		RegistrationImage: o.RegistrationImage,
		CreatedAt:         o.CreatedAt,
		UpdatedAt:         o.UpdatedAt,
		DeletedAt:         o.DeletedAt,
	}, nil
}

func (rt *Handlers) ListOwnerOrganization(ctx context.Context, ownerID uuid.UUID) ([]*organizationentity.Organization, error) {
	o, err := rt.os.ListOwnerOrganization(ctx, ownerID)
	if err != nil {
		return nil, fmt.Errorf("error when getting by owner: %w", err)
	}

	return o, nil
}

func (rt *Handlers) SearchOrganizations(ctx context.Context, query string) ([]*organizationentity.Organization, error) {
	o, err := rt.os.SearchOrganizations(ctx, query)
	if err != nil {
		fmt.Printf("%v\n", err)
		return nil, fmt.Errorf("error when search organization: %w", err)
	}

	return o, nil
}

func (rt *Handlers) UpdateOrganization(ctx context.Context, params *organizationrepo.UpdateOrganizationParams) (*Organization, error) {
	o, err := rt.os.UpdateOrganization(ctx, params)
	if err != nil {
		return &Organization{}, fmt.Errorf("error when updating: %w", err)
	}

	return &Organization{
		ID:                o.ID,
		Name:              o.Name,
		CountryID:         o.CountryID,
		OwnerID:           o.OwnerID,
		Verified:          o.Verified,
		RegistrationCode:  o.RegistrationCode,
		RegistrationDate:  o.RegistrationDate,
		RegistrationImage: o.RegistrationImage,
		CreatedAt:         o.CreatedAt,
		UpdatedAt:         o.UpdatedAt,
		DeletedAt:         o.DeletedAt,
	}, nil
}

func (rt *Handlers) DeleteOrganization(ctx context.Context, id uuid.UUID, userID uuid.UUID) (*Organization, error) {
	myOrg, err := rt.os.GetOrganization(ctx, id)
	if err != nil {
		return &Organization{}, fmt.Errorf("error when updating: %w", err)
	}

	if myOrg.OwnerID != userID {
		return &Organization{}, fmt.Errorf("error when deleting organization")
	}

	o, err := rt.os.DeleteOrganization(ctx, id)
	if err != nil {
		return &Organization{}, fmt.Errorf("error when updating: %w", err)
	}

	return &Organization{
		ID:                o.ID,
		Name:              o.Name,
		CountryID:         o.CountryID,
		OwnerID:           o.OwnerID,
		Verified:          o.Verified,
		RegistrationCode:  o.RegistrationCode,
		RegistrationDate:  o.RegistrationDate,
		RegistrationImage: o.RegistrationImage,
		CreatedAt:         o.CreatedAt,
		UpdatedAt:         o.UpdatedAt,
		DeletedAt:         o.DeletedAt,
	}, nil
}

func (rt *Handlers) VerifyOrganization(ctx context.Context, id uuid.UUID) (*Organization, error) {
	o, err := rt.os.VerifyOrganization(ctx, id)
	if err != nil {
		return &Organization{}, fmt.Errorf("error when updating: %w", err)
	}

	return &Organization{
		ID:                o.ID,
		Name:              o.Name,
		CountryID:         o.CountryID,
		OwnerID:           o.OwnerID,
		Verified:          o.Verified,
		RegistrationCode:  o.RegistrationCode,
		RegistrationDate:  o.RegistrationDate,
		RegistrationImage: o.RegistrationImage,
		CreatedAt:         o.CreatedAt,
		UpdatedAt:         o.UpdatedAt,
		DeletedAt:         o.DeletedAt,
	}, nil
}

func (rt *Handlers) UploadRegistration(ctx context.Context, params *organizationrepo.UploadOrganizationParams) (*Organization, error) {
	o, err := rt.os.AddRegistrationImage(ctx, params)
	if err != nil {
		return &Organization{}, fmt.Errorf("error when updating organization registartion: %w", err)
	}

	return &Organization{
		ID:                o.ID,
		Name:              o.Name,
		CountryID:         o.CountryID,
		OwnerID:           o.OwnerID,
		Verified:          o.Verified,
		RegistrationCode:  o.RegistrationCode,
		RegistrationDate:  o.RegistrationDate,
		RegistrationImage: o.RegistrationImage,
		CreatedAt:         o.CreatedAt,
		UpdatedAt:         o.UpdatedAt,
		DeletedAt:         o.DeletedAt,
	}, nil
}
