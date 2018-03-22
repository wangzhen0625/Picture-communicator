package routers

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"time"
)

type Route struct {
	Name        string
	Method      string
	Path        string
	HandlerFunc httprouter.Handle
}

type Routes []Route

// A Logger function which simply wraps the handler function around some log messages
func Logger(fn func(w http.ResponseWriter, r *http.Request, param httprouter.Params)) func(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
		start := time.Now()
		log.Printf("%s %s", r.Method, r.URL.Path)
		fn(w, r, param)
		log.Printf("Done in %v (%s %s)", time.Since(start), r.Method, r.URL.Path)
	}
}

//Reads from the routes slice to translate the values to httprouter.Handle
func NewRouter(routes Routes) *httprouter.Router {
	router := httprouter.New()
	for _, route := range routes {
		var handle httprouter.Handle

		handle = route.HandlerFunc
		handle = Logger(handle)

		router.Handle(route.Method, route.Path, handle)
	}

	return router
}
