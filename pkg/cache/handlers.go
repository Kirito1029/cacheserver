package cache

import (
	"net/http"
	"strconv"
)

func getFromCache(key string) (value []byte, err error) {
	value, err = ht.Get(key)
	return
}

func setCache(key string, value []byte, expiry uint64) {
	ht.Set(key, value, expiry)
}

func expireCache(key string, ttl uint64) {
	ht.Expire(key, ttl)
}

func cacheGetHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	key := r.Form.Get("key")
	res, err := getFromCache(key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(res)
}

func cacheSetHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	key := r.Form.Get("key")
	value := r.Form.Get("value")
	expiry, _ := strconv.Atoi(r.Form.Get("expiry"))
	setCache(key, []byte(value), uint64(expiry))
}

func cacheExpireHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	key := r.Form.Get("key")
	expiry, _ := strconv.Atoi(r.Form.Get("expiry"))
	expireCache(key, uint64(expiry))
}
