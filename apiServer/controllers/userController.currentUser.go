package controllers

import (
	"net/http"

	"github.com/Alfazal007/apiserver/helpers"
	"github.com/Alfazal007/apiserver/internal/database"
)

func (apiCfg *ApiConf) GetCurrentUser(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value("user").(database.User)
	if !ok {
		helpers.RespondWithError(w, 400, "Issue with finding the user from the database", []string{})
		return
	}
	helpers.RespondWithJSON(w, 200, helpers.GenerateUserToReturn(user))
}
