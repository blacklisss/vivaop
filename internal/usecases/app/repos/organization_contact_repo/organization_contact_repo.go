package organization_contact_repo

import (
	"context"
	"fmt"
	"vivaop/internal/entities/organization_contact_entity"

	"github.com/google/uuid"
)

type OrganizationContactsStore interface {
	CreateOrganizationContact(ctx context.Context, arg *CreateOrganizationContactParams) (*organization_contact_entity.OrganizationContact, error)
	DeleteOrganizationContact(ctx context.Context, id uuid.UUID) (*organization_contact_entity.OrganizationContact, error)
	GetOrganizationContact(ctx context.Context, id uuid.UUID) (*organization_contact_entity.OrganizationContact, error)
	ListOrganizationContacts(ctx context.Context, organizationID uuid.UUID) ([]*organization_contact_entity.OrganizationContact, error)
}

type OrganizationContacts struct {
	ocstore OrganizationContactsStore
}

func NewOrganizationContactsStore(ocstore OrganizationContactsStore) *OrganizationContacts {
	return &OrganizationContacts{
		ocstore,
	}
}

type CreateOrganizationContactParams struct {
	ID             uuid.UUID `json:"id"`
	OrganizationID uuid.UUID `json:"organization_id"`
	Name           string    `json:"name"`
	Phone          string    `json:"phone"`
}

func (os *OrganizationContacts) CreateOrganizationContact(ctx context.Context, params *CreateOrganizationContactParams) (*organization_contact_entity.OrganizationContact, error) {
	organizationContact, err := os.ocstore.CreateOrganizationContact(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("create organization contact error: %w", err)
	}

	return organizationContact, nil
}

func (os *OrganizationContacts) DeleteOrganizationContact(ctx context.Context, id uuid.UUID) (*organization_contact_entity.OrganizationContact, error) {
	organizationContact, err := os.ocstore.DeleteOrganizationContact(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("delete organization contact error: %w", err)
	}

	return organizationContact, nil
}

func (os *OrganizationContacts) GetOrganizationContact(ctx context.Context, id uuid.UUID) (*organization_contact_entity.OrganizationContact, error) {
	organizationContact, err := os.ocstore.GetOrganizationContact(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("create organization contact error: %w", err)
	}

	return organizationContact, nil
}

func (os *OrganizationContacts) ListOrganizationContacts(ctx context.Context, organizationID uuid.UUID) ([]*organization_contact_entity.OrganizationContact, error) {
	organizationContacts, err := os.ocstore.ListOrganizationContacts(ctx, organizationID)
	if err != nil {
		return nil, fmt.Errorf("list organization error: %w", err)
	}

	return organizationContacts, nil
}
