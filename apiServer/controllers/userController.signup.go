package controllers

import (
	"net/http"

	"github.com/Alfazal007/apiserver/helpers"
)

type Message struct {
	Data string
}

func (apiCfg *ApiConf) Signup(w http.ResponseWriter, r *http.Request) {
	message := Message{
		Data: "hello data",
	}
	helpers.RespondWithJSON(w, 200, message)
}
