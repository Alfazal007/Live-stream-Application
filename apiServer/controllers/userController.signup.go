package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/Alfazal007/apiserver/helpers"
	"github.com/Alfazal007/apiserver/internal/database"
	"github.com/Alfazal007/apiserver/utils"
	"github.com/Alfazal007/apiserver/validators"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

func (apiCfg *ApiConf) Signup(w http.ResponseWriter, r *http.Request) {
	validate := validator.New()
	var signUpParams validators.SignupValidators

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&signUpParams)
	if err != nil {
		helpers.RespondWithError(w, 400, fmt.Sprintf("Invalid json object provided %v", err.Error()), []string{})
		return
	}

	signUpParams.Username = strings.TrimSpace(signUpParams.Username)
	signUpParams.Password = strings.TrimSpace(signUpParams.Password)

	err = validate.Struct(signUpParams)
	if err != nil {
		var errMessageArray []string
		for _, fieldError := range err.(validator.ValidationErrors) {
			errMessageArray = append(errMessageArray, fmt.Sprintf("Validation failed on %v at the tag %v", fieldError.Field(), fieldError.Tag()))
		}
		helpers.RespondWithError(w, 400, "Validation errors", errMessageArray)
		return
	}

	_, err = apiCfg.DB.GetUserByName(r.Context(), signUpParams.Username)
	if err != nil && err != sql.ErrNoRows {
		helpers.RespondWithError(w, 400, "Issue talking to the database", []string{})
		return
	}

	if err != sql.ErrNoRows {
		helpers.RespondWithError(w, 400, "Use different username", []string{})
		return
	}
	hashedPassword, err := utils.HashPassword(signUpParams.Password)
	if err != nil {
		helpers.RespondWithError(w, 400, "Error hashing the password", []string{})
		return
	}

	newUser, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:       uuid.New(),
		Username: signUpParams.Username,
		Password: hashedPassword,
	})
	if err != nil {
		helpers.RespondWithError(w, 400, "Issue creating the user", []string{})
		return
	}

	helpers.RespondWithJSON(w, 200, helpers.GenerateUserToReturn(newUser))
}
