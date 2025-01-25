package main

import (
	"gateway/internal/http/controllers"
	"gateway/internal/http/routes"
	"gateway/internal/repositories/state"
	"gateway/internal/repositories/tokens"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {
	r := chi.NewRouter()
	sr := state.NewInMemory()
	c := controllers.New(sr)
	tr := tokens.NewInMemory("very-secret")

	routes.Register(r, c, tr)

	if err := http.ListenAndServe("127.0.0.1:3000", r); err != nil {
		panic(err)
	}
}
