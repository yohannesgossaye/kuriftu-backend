package middleware

import (
	"net/http"

	"github.com/rs/zerolog"
)

func Logger(log *zerolog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Info().Str("method", r.Method).Str("path", r.URL.Path).Msg("Request received")
			next.ServeHTTP(w, r)
		})
	}
}
