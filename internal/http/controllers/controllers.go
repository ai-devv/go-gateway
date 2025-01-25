package controllers

import (
	"gateway/internal/http/controllers/oauth"
	"gateway/internal/repositories/state"
)

type Controllers struct {
	OAuthController *oauth.Controller
}

func New(sr state.Repository) *Controllers {
	return &Controllers{
		OAuthController: oauth.NewOAuthController(sr),
	}
}
