package httprouter

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Implementation of gorilla mux framework
type muxRouter struct{}

var (
	muxDispatcher = mux.NewRouter()
)

// Create a new Mux Router
func NewMuxRouter() Router {
	return &muxRouter{}
}

// Register GET uris with this method
func (*muxRouter) GET(uri string, f func(rw http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("GET")
}

// Register POST uris with this method
func (*muxRouter) POST(uri string, f func(rw http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("POST")
}

// start the server
func (*muxRouter) SERVE(port string) error {
	log.Println("Mux HTTP server running on port: ", port)
	return http.ListenAndServe(port, muxDispatcher)
}
