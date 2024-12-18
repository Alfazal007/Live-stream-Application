package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/Alfazal007/apiserver/helpers"
	"github.com/Alfazal007/apiserver/utils"
	"github.com/Alfazal007/apiserver/validators"
	"github.com/go-playground/validator/v10"
)

func (apiCfg *ApiConf) IsValidStream(w http.ResponseWriter, r *http.Request) {
	validate := validator.New()
	var isValidStreamParams validators.IsValidStream

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&isValidStreamParams)
	if err != nil {
		helpers.RespondWithError(w, 400, fmt.Sprintf("Invalid json object provided %v", err.Error()), []string{})
		return
	}

	isValidStreamParams.AdminId = strings.TrimSpace(isValidStreamParams.AdminId)
	isValidStreamParams.StreamId = strings.TrimSpace(isValidStreamParams.StreamId)
	isValidStreamParams.Secret = strings.TrimSpace(isValidStreamParams.Secret)

	err = validate.Struct(isValidStreamParams)
	if err != nil {
		var errMessageArray []string
		for _, fieldError := range err.(validator.ValidationErrors) {
			errMessageArray = append(errMessageArray, fmt.Sprintf("Validation failed on %v at the tag %v", fieldError.Field(), fieldError.Tag()))
		}
		helpers.RespondWithError(w, 400, "Validation errors", errMessageArray)
		return
	}
	if isValidStreamParams.Secret != utils.LoadEnvFiles().Secret {
		helpers.RespondWithError(w, 400, "Invalid secret", []string{})
		return
	}
	requiredStream, err := apiCfg.DB.GetStreamFromIdForWS(r.Context(), isValidStreamParams.StreamId)
	if err != nil {
		helpers.RespondWithError(w, 400, "Invalid stream", []string{})
		return
	}
	adminIdFromWS := requiredStream.AdminID.UUID.String()
	if adminIdFromWS != isValidStreamParams.AdminId {
		helpers.RespondWithError(w, 400, "You are not the admin", []string{})
		return
	}
	if !requiredStream.Started.Valid {
		helpers.RespondWithError(w, 400, "Invalid stream", []string{})
		return
	}
	if !requiredStream.Started.Bool {
		helpers.RespondWithError(w, 400, "Not started yet", []string{})
		return
	}
	if !requiredStream.Ended.Valid {
		helpers.RespondWithError(w, 400, "Invalid stream", []string{})
		return
	}
	if requiredStream.Ended.Bool {
		helpers.RespondWithError(w, 400, "Stream ended", []string{})
		return
	}
	helpers.RespondWithJSON(w, 200, "")
}
