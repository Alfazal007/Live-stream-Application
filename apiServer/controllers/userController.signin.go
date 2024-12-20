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

type Token struct {
	AccessToken string `json:"accessToken"`
	Username    string `json:"username"`
	Id          string `json:"id"`
}

func (apiCfg *ApiConf) Signin(w http.ResponseWriter, r *http.Request) {
	validate := validator.New()
	var signInParams validators.SigninValidators

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&signInParams)
	if err != nil {
		helpers.RespondWithError(w, 400, fmt.Sprintf("Invalid json object provided %v", err.Error()), []string{})
		return
	}

	signInParams.Username = strings.TrimSpace(signInParams.Username)
	signInParams.Password = strings.TrimSpace(signInParams.Password)

	err = validate.Struct(signInParams)
	if err != nil {
		var errMessageArray []string
		for _, fieldError := range err.(validator.ValidationErrors) {
			errMessageArray = append(errMessageArray, fmt.Sprintf("Validation failed on %v at the tag %v", fieldError.Field(), fieldError.Tag()))
		}
		helpers.RespondWithError(w, 400, "Validation errors", errMessageArray)
		return
	}

	existingUser, err := apiCfg.DB.GetUserByName(r.Context(), signInParams.Username)
	if err != nil {
		helpers.RespondWithError(w, 400, "Issue finding the user", []string{})
		return
	}

	isValidPassword := utils.VerifyPassword(signInParams.Password, existingUser.Password)
	if !isValidPassword {
		helpers.RespondWithError(w, 400, "Invalid password", []string{})
		return
	}

	token, err := utils.GenerateJWT(existingUser)
	if err != nil {
		helpers.RespondWithError(w, 400, "Issue generating the jwt tokens", []string{})
		return
	}

	helpers.RespondWithJSON(w, 200, Token{AccessToken: token, Username: existingUser.Username, Id: existingUser.ID.String()})
}
