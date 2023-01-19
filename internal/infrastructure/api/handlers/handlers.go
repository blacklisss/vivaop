package handlers

import (
	"vivaop/internal/usecases/app/repos/countryrepo"
	"vivaop/internal/usecases/app/repos/userrepo"
)

type Handlers struct {
	us *userrepo.Users
	cs *countryrepo.Countries
}

func NewHandlers(us *userrepo.Users, cs *countryrepo.Countries) *Handlers {
	handlers := &Handlers{
		us,
		cs,
	}
	return handlers
}
