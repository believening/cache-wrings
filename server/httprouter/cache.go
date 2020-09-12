package httprouter

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type cacheHandler struct {
	*httpRouterServer
}

func (h *cacheHandler) GetHandle(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	key := ps.ByName("key")
	if key == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
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
}

func (h *cacheHandler) PutHandle(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	key := ps.ByName("key")
	if key == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
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
}

func (h *cacheHandler) DeleteHandle(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	key := ps.ByName("key")
	if key == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err := h.Del(key)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
