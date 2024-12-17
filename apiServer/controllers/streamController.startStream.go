package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/Alfazal007/apiserver/helpers"
	"github.com/Alfazal007/apiserver/internal/database"
	"github.com/Alfazal007/apiserver/validators"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

func (apiCfg *ApiConf) StartStream(w http.ResponseWriter, r *http.Request) {
	adminStream, ok := r.Context().Value("user").(database.User)
	if !ok {
		helpers.RespondWithError(w, 400, "Issue with finding the user from the database", []string{})
		return
	}
	validate := validator.New()
	var startStreamParams validators.StartEndStream

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&startStreamParams)
	if err != nil {
		helpers.RespondWithError(w, 400, fmt.Sprintf("Invalid json object provided %v", err.Error()), []string{})
		return
	}

	startStreamParams.StreamId = strings.TrimSpace(startStreamParams.StreamId)

	err = validate.Struct(startStreamParams)
	if err != nil {
		var errMessageArray []string
		for _, fieldError := range err.(validator.ValidationErrors) {
			errMessageArray = append(errMessageArray, fmt.Sprintf("Validation failed on %v at the tag %v", fieldError.Field(), fieldError.Tag()))
		}
		helpers.RespondWithError(w, 400, "Validation errors", errMessageArray)
		return
	}

	requiredStream, err := apiCfg.DB.StartStream(r.Context(), database.StartStreamParams{
		ID: startStreamParams.StreamId,
		AdminID: uuid.NullUUID{
			Valid: true,
			UUID:  adminStream.ID,
		},
	})
	if err != nil {
		helpers.RespondWithError(w, 400, "Issue finding the stream", []string{})
		return
	}
	helpers.RespondWithJSON(w, 200, requiredStream)
}
