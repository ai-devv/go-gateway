package main

import (
	"gateway/internal/http/controllers"
	"gateway/internal/http/routes"
	"gateway/internal/repositories/state"
	"gateway/internal/repositories/tokens"
	"github.com/go-chi/chi/v5"
	"github.com/redis/go-redis/v9"
	"net/http"
)

func main() {
	r := chi.NewRouter()

	rdb := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		DB:   0,
	})

	sr := state.NewRedis(rdb)
	c := controllers.New(sr)
	tr := tokens.NewInMemory("very-secret")

	routes.Register(r, c, tr)

	if err := http.ListenAndServe("127.0.0.1:3000", r); err != nil {
		panic(err)
	}
}
