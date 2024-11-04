// Package middleware provides functionality for middleware.
package middleware

import (
	"net/http"

	connectcors "connectrpc.com/cors"
	"github.com/rs/cors"
)

// WithCORS function is returning an HTTP Handler with configured CORS settings.
func WithCORS(h http.Handler) http.Handler {
	var token = []string{"accessToken", "refreshToken"}

	middleware := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000", "http://localhost:8080", "http://localhost:8081", "http://localhost:8082"},
		AllowedMethods: connectcors.AllowedMethods(),
		AllowedHeaders: append(connectcors.AllowedHeaders(), token...),
		ExposedHeaders: connectcors.ExposedHeaders(),
	})
	return middleware.Handler(h)
}
