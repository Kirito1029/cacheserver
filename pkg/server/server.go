package server

import "net/http"

func init() {

}

func Start(s Server) (err error) {
	return http.ListenAndServe(":8080", s.Handler())
}

type Server interface {
	Handler() http.Handler
}
