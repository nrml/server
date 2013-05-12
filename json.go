package server

import (
	"encoding/json"
	"net/http"
)

type Status struct {
	Code    int8
	Message string
}

func Json(w http.ResponseWriter, obj interface{}) {
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	enc.Encode(obj)
}
func JsonStatus(w http.ResponseWriter, code int8, msg string) {
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	s := Status{code, msg}
	enc.Encode(s)
}
func JsonErr(w http.ResponseWriter, err error) {
	JsonStatus(w, -1, err.Error())
}
