package middlewares

import (
	"gateway/internal/repositories/tokens"
	"net/http"
	"regexp"
)

func Authenticate(tr tokens.Repository) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authorizationHeader := r.Header.Get("authorization")
			matches := regexp.MustCompile(`(?i)^Bearer\s(.*)$`).FindStringSubmatch(authorizationHeader)

			if len(matches) < 2 {
				w.WriteHeader(http.StatusUnauthorized)

				return
			}

			if !tr.Check(tokens.Token(matches[1])) {
				w.WriteHeader(http.StatusUnauthorized)

				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
