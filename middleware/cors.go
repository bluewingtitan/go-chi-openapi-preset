package middleware

import (
	"github.com/go-chi/cors"
	"net/http"
)

func CorsMiddleware(allowedOrigins []string) func(next http.Handler) http.Handler {

	corsOpts := cors.Options{
		// Allow only specific origins
		AllowedOrigins: allowedOrigins,
		// Consider allowing specific methods such as GET, OPTIONS
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
	}

	return cors.Handler(corsOpts)
}
