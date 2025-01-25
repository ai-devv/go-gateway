package controllers

import (
	"gateway/internal/http/controllers/oauth"
	"gateway/internal/repositories/state"
)

type Controllers struct {
	OAuth *oauth.Controller
}

func New(sr state.Repository) *Controllers {
	return &Controllers{
		OAuth: oauth.NewController(sr),
	}
}
