package main

import (
	"net/http"
	"testing"

	"github.com/go-chi/chi/v5"
)

//Test_routes_exist checks if all routes exist
func Test_routes_exist(t *testing.T) {

	testApp := Config{}

	testRoutes := testApp.routes()

	chiRoutes := testRoutes.(chi.Router)

	//define routes to check
	routes := []string{
		"/authenticate",
	}

	//for loop to check if all routes exist and executes routeExists function
	for _, route := range routes {
		routeExists(t, chiRoutes, route)
	}
}

//routeExists uses chi.Walk to check if a route exists
func routeExists(t *testing.T, routes chi.Router, route string) {
	found := false

	_ = chi.Walk(routes, func(method string, foundRoute string, handler http.Handler, middleware ...func(http.Handler) http.Handler) error {
		if route == foundRoute {
			found = true
		}
		return nil

	})

	if !found {
		t.Errorf("Route %s not found", route)
	}

}
