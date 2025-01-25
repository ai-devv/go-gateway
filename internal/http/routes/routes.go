package routes

import (
	"gateway/internal/http/controllers"
	"gateway/internal/http/middlewares"
	"gateway/internal/repositories/tokens"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Register(
	r *chi.Mux,
	c *controllers.Controllers,
	tr tokens.Repository,
) {
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)

	// ========== OAuth ==========
	// ===== authorize =====
	r.
		With(middlewares.Authenticate(tr)).
		Post("/oauth/authorize", c.OAuth.Authorize)

	// ===== callback =====
	r.Get("/oauth/callback", c.OAuth.Callback)
}
