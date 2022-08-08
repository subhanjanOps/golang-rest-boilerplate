package routes

import (
	"golang-rest/helpers"
	"golang-rest/middlewares"
	"golang-rest/services"
	"net/http"

	"github.com/gorilla/mux"
)

var Routes = []helpers.RoutePrefix{
	{
		Prefix: "/auth",
		SubRoutes: []helpers.Route{
			{
				Name:        "register",
				Method:      "POST",
				Pattern:     "/register",
				Protected:   false,
				HandlerFunc: services.RegisterUser,
			},
			{
				Name:        "Login",
				Method:      "POST",
				Pattern:     "/login",
				Protected:   false,
				HandlerFunc: services.LoginUser,
			},
			{
				Name:        "Logout",
				Method:      "GET",
				Pattern:     "/logout",
				Protected:   true,
				HandlerFunc: services.LogoutUser,
			},
		},
	},
	{
		Prefix: "/users",
		SubRoutes: []helpers.Route{
			{
				Name:        "User List",
				Method:      "GET",
				Pattern:     "/",
				Protected:   true,
				HandlerFunc: services.UserList,
			},
		},
	},
}

// var DocRoutes = helpers.RoutePrefix{
// 	Prefix: "/docs",
// 	SubRoutes: []helpers.Route{
// 		{
// 			Name:        "Documentation",
// 			Method:      "GET",
// 			Pattern:     "",
// 			Protected:   false,
// 			HandlerFunc: httpSwagger.WrapHandler,
// 		},
// 	},
// }

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	var mw middlewares.Middleware
	helpers.AppRoutes = append(helpers.AppRoutes, Routes...) // Appending application routes
	// helpers.AppRoutes = append(helpers.AppRoutes, DocRoutes) // Appending swagger routes
	for _, route := range helpers.AppRoutes {
		routerPrefix := router.PathPrefix(route.Prefix).Subrouter()
		if route.Prefix != "/docs" {
			routerPrefix.Use(mw.BasicMiddleware)                     // Basic header setup  middleware
			routerPrefix.Use(mux.CORSMethodMiddleware(routerPrefix)) // Default CORS middleware
		}
		routerPrefix.Use(mw.Logger)
		for _, r := range route.SubRoutes {
			var handler http.Handler
			handler = r.HandlerFunc

			if r.Protected {
				handler = mw.Authenticate(r.HandlerFunc) // Auth middleware
			}

			routerPrefix.
				Path(r.Pattern).
				Handler(handler).
				Methods(r.Method).
				Name(r.Name)
		}
	}

	return router

}
