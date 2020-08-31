package router

import (
	"github.com/chulista/weather_api/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

type Route struct {
	Name 			string
	Method 			string
	Pattern 		string
	HandlerFunc 	http.HandlerFunc
}
type Routes []Route

func NewRouter() *mux.Router{
	router := mux.NewRouter().StrictSlash(true)
	for _,route := range routes{
		router.
			Name(route.Name).
			Methods(route.Method).
			Path(route.Pattern).
			Handler(route.HandlerFunc)
	}
	return router
}
var routes = Routes{
	Route{"HelloWorld", "GET", "/", handlers.HelloWorld},
	Route{"Ping", "GET", "/ping", handlers.Ping},
	Route{"Options", "OPTIONS", "/v1", handlers.Options},
	Route{"Location", "GET", "/v1/location", handlers.GetLocation},
	//Route{"City", "GET", "/current/{city}", handlers.GetCityStatus},
	//Route{"CityHistory", "GET", "/forecast/{city}", handlers.GetCityStatusHistory},
}