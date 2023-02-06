package handlers

import (
	"vivaop/internal/usecases/app/repos/countryrepo"
	"vivaop/internal/usecases/app/repos/organizationrepo"
	"vivaop/internal/usecases/app/repos/sessionrepo"
	"vivaop/internal/usecases/app/repos/userrepo"
)

type Handlers struct {
	us *userrepo.Users
	cs *countryrepo.Countries
	ss *sessionrepo.Session
	os *organizationrepo.Organizations
}

func NewHandlers(us *userrepo.Users, cs *countryrepo.Countries, ss *sessionrepo.Session, os *organizationrepo.Organizations) *Handlers {
	handlers := &Handlers{
		us,
		cs,
		ss,
		os,
	}
	return handlers
}
