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
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func (apiCfg *ApiConf) IsValidUser(w http.ResponseWriter, r *http.Request) {
	validate := validator.New()
	var isValidUserParams validators.IsValidUser

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&isValidUserParams)
	if err != nil {
		helpers.RespondWithError(w, 400, fmt.Sprintf("Invalid json object provided %v", err.Error()), []string{})
		return
	}

	isValidUserParams.UserId = strings.TrimSpace(isValidUserParams.UserId)
	isValidUserParams.Token = strings.TrimSpace(isValidUserParams.Token)

	err = validate.Struct(isValidUserParams)
	if err != nil {
		var errMessageArray []string
		for _, fieldError := range err.(validator.ValidationErrors) {
			errMessageArray = append(errMessageArray, fmt.Sprintf("Validation failed on %v at the tag %v", fieldError.Field(), fieldError.Tag()))
		}
		helpers.RespondWithError(w, 400, "Validation errors", errMessageArray)
		return
	}

	jwtToken := isValidUserParams.Token
	jwtSecret := utils.LoadEnvFiles().AccessTokenSecret
	if jwtToken == "" {
		helpers.RespondWithError(w, 400, "Provide access token", []string{})
		return
	}

	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtSecret), nil
	})
	if err != nil {
		helpers.RespondWithError(w, 401, "Invalid token here", []string{})
		return
	}
	if !token.Valid {
		helpers.RespondWithError(w, 401, "Invalid token", []string{})
		return
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		helpers.RespondWithError(w, 400, "Invalid claims login again", []string{})
		return
	}

	username := claims["username"].(string)
	id := claims["user_id"].(string)

	user, err := apiCfg.DB.GetUserByName(r.Context(), username)
	if err != nil {
		helpers.RespondWithError(w, 400, "Some manpulation done with the token", []string{})
		return
	}
	idUUID, err := uuid.Parse(id)
	if err != nil {
		helpers.RespondWithError(w, 400, "Some manpulation done with the token", []string{})
		return
	}
	if idUUID != user.ID {
		helpers.RespondWithError(w, 400, "Some manipulations done with the token try again", []string{})
		return
	}
	if id != isValidUserParams.UserId {
		helpers.RespondWithError(w, 400, "Some manipulations done with the token try again", []string{})
		return
	}
	helpers.RespondWithJSON(w, 200, "")
}
