package cache

import (
	"cache-server/pkg/hashtable"
	"cache-server/pkg/server"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	ht *hashtable.HashTable
)

func init() {
	ht = hashtable.New()
}

func Start() {
	server.Start(&cacheServer{})
}

type (
	cacheServer struct {
		// test int
	}
)

func (c cacheServer) Handler() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/Get", cacheGetHandler)
	r.HandleFunc("/Set", cacheSetHandler)
	r.HandleFunc("/Expire", cacheExpireHandler)
	return r
}
