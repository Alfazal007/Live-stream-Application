package controllers

import (
	"database/sql"
	"net/http"

	"github.com/Alfazal007/apiserver/helpers"
	"github.com/Alfazal007/apiserver/internal/database"
)

func (apiCfg *ApiConf) GetTenStreams(w http.ResponseWriter, r *http.Request) {
	_, ok := r.Context().Value("user").(database.User)
	if !ok {
		helpers.RespondWithError(w, 400, "Issue with finding the user from the database", []string{})
		return
	}
	streamsFromDB, err := apiCfg.DB.Get10LatestStream(r.Context())
	if err == sql.ErrNoRows {
		helpers.RespondWithJSON(w, 200, []database.Get10LatestStreamRow{})
	}
	if err != nil {
		helpers.RespondWithError(w, 400, "Issue finding the streams", []string{})
	}
	helpers.RespondWithJSON(w, 200, streamsFromDB)
}
