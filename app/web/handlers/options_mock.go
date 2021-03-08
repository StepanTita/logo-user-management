package handlers

import (
	"github.com/logo-user-management/app/ctx"
	"net/http"
)

func OptionsMock(w http.ResponseWriter, r *http.Request) {
	log := ctx.Log(r)
	log.Info("Starting options request")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, POST, GET, PATCH, DELETE, PUT")
	w.Header().Set("Access-Control-Max-Age", "2592000")
	w.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Origin, Cache-Control, Pragma, Authorization, Accept, Accept-Encoding")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.WriteHeader(http.StatusNoContent)
}
