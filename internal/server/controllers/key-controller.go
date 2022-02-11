package controllers

import (
	"encoding/json"
	"main/internal/server/services"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	c := make(chan string, 1)
	v := services.Default(r)
	services.GivePublicKey(v, c)
	json.NewEncoder(w).Encode(map[string]string{"PublicKey": <-c})
}
