package oauth

import (
	"gateway/internal/repositories/state"
)

type Controller struct {
	sr state.Repository
}

func NewController(sr state.Repository) *Controller {
	return &Controller{
		sr,
	}
}
