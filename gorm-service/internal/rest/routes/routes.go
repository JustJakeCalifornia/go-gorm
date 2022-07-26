package routes

import (
	"net/http"

	"github.com/trite8q1/go-gorm/internal/rest"
)

const (
	APIPrefix    = "/api"
	APIVersionV1 = APIPrefix + "/v1"
	APIVersionV2 = APIPrefix + "/v2"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	APIVersion  string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Health",
		"POST",
		"/user",
		APIVersionV1,
		rest.AddUser,
	},
	Route{
		"Health",
		"GET",
		"/user",
		APIVersionV1,
		rest.GetUser,
	},
}
