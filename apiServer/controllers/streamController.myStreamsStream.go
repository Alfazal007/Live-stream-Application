package controllers

import (
	"database/sql"
	"net/http"

	"github.com/Alfazal007/apiserver/helpers"
	"github.com/Alfazal007/apiserver/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *ApiConf) GetMyStream(w http.ResponseWriter, r *http.Request) {
	existingUser, ok := r.Context().Value("user").(database.User)
	if !ok {
		helpers.RespondWithError(w, 400, "Issue with finding the user from the database", []string{})
		return
	}
	streamsFromDB, err := apiCfg.DB.GetMyStreams(r.Context(), uuid.NullUUID{Valid: true, UUID: existingUser.ID})
	if err == sql.ErrNoRows {
		helpers.RespondWithJSON(w, 200, []database.Get10LatestStreamRow{})
	}
	if err != nil {
		helpers.RespondWithError(w, 400, "Issue finding the streams", []string{})
	}
	helpers.RespondWithJSON(w, 200, updateResponseDataMyStream(streamsFromDB))
}

func updateResponseDataMyStream(dataIn []database.GetMyStreamsRow) []UpdatedTypes {
	response := make([]UpdatedTypes, 0)
	for i := 0; i < len(dataIn); i++ {
		response = append(response, UpdatedTypes{
			CreatorName: dataIn[i].Username,
			Id:          dataIn[i].ID,
		})
	}
	return response
}
