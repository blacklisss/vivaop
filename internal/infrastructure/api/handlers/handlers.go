package handlers

import (
	"vivaop/internal/usecases/app/repos/countryrepo"
	"vivaop/internal/usecases/app/repos/organization_contact_repo"
	"vivaop/internal/usecases/app/repos/organizationrepo"
	"vivaop/internal/usecases/app/repos/sessionrepo"
	"vivaop/internal/usecases/app/repos/userrepo"
)

type Handlers struct {
	us  *userrepo.Users
	cs  *countryrepo.Countries
	ss  *sessionrepo.Session
	os  *organizationrepo.Organizations
	ocs *organization_contact_repo.OrganizationContacts
}

func NewHandlers(
	us *userrepo.Users,
	cs *countryrepo.Countries,
	ss *sessionrepo.Session,
	os *organizationrepo.Organizations,
	ocs *organization_contact_repo.OrganizationContacts,
) *Handlers {
	handlers := &Handlers{
		us,
		cs,
		ss,
		os,
		ocs,
	}
	return handlers
}
