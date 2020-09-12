package gohttp

import (
	"encoding/json"
	"net/http"
)

type statusHandler struct {
	*goHTTPServer
}

func (s *statusHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	stat, err := json.Marshal(s.GetStat())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, _ = w.Write(stat)
}
