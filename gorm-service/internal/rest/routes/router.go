package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

// NewRouter
func NewRouter() *chi.Mux {
	router := chi.NewRouter()

	router.Route(APIVersionV1, func(r chi.Router) {
		getRoutes(router)
	})

	return router
}

func getRoutes(router *chi.Mux) []Route {
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc

		router.Method(route.Method, route.APIVersion+route.Pattern, handler)
	}
	return routes
}
