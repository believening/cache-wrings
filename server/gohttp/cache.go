package gohttp

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type cacheHanler struct {
	*goHTTPServer
}

func (h *cacheHanler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// ip:port/cache/key -> [cache key] -> [key]
	key := strings.Split(r.URL.EscapedPath(), "/")[2]
	if key == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	switch r.Method {
	case http.MethodGet:
		val, err := h.Get(key)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if len(val) == 0 {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		_, _ = w.Write(val)
	case http.MethodPut:
		val, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if len(val) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = h.Set(key, val)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	case http.MethodDelete:
		err := h.Del(key)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
