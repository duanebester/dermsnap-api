package api

import (
	"dermsnap/services"
)

type API struct {
	services *services.Services
}

func NewApi(s *services.Services) API {
	return API{
		services: s,
	}
}
