package httprouter

import "net/http"

type Router interface {
	GET(uri string, f func(rw http.ResponseWriter, r *http.Request))
	POST(uri string, f func(rw http.ResponseWriter, r *http.Request))
	SERVE(port string) error
}
