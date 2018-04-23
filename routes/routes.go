package routes

import (
	"net/http"
	"sebatibot/controllers"
)

// Route is used to define http routes for the app
type Route struct {
	Name        string
	Pattern     string
	Method      string
	HandlerFunc http.HandlerFunc
}

// Routes is a collection of multiple http Routes
type Routes []Route

var UsersRoutes = Routes{
	Route{
		Name:        "Index",
		Pattern:     "/",
		Method:      "GET",
		HandlerFunc: controllers.Index,
	},
}