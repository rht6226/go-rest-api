package httprouter

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

type chiRouter struct{}

var (
	chiDispatcher = chi.NewRouter()
)

// New Chi Router
func NewChiRouter() Router {
	return &chiRouter{}
}

// Register GET uris with this method
func (*chiRouter) GET(uri string, f func(rw http.ResponseWriter, r *http.Request)) {
	chiDispatcher.Get(uri, f)
}

// Register POST uris with this method
func (*chiRouter) POST(uri string, f func(rw http.ResponseWriter, r *http.Request)) {
	chiDispatcher.Post(uri, f)
}

// start the server
func (*chiRouter) SERVE(port string) error {
	log.Println("Chi HTTP server running on port: ", port)
	return http.ListenAndServe(port, chiDispatcher)
}
