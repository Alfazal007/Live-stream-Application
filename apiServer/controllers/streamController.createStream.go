package controllers

import (
	"math/rand"
	"net/http"
	"strings"

	"github.com/Alfazal007/apiserver/helpers"
	"github.com/Alfazal007/apiserver/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *ApiConf) CreateStream(w http.ResponseWriter, r *http.Request) {
	adminStream, ok := r.Context().Value("user").(database.User)
	if !ok {
		helpers.RespondWithError(w, 400, "Issue with finding the user from the database", []string{})
		return
	}
	streamId := generateStreamId()
	newStream, err := apiCfg.DB.CreateStream(r.Context(), database.CreateStreamParams{
		ID: streamId,
		AdminID: uuid.NullUUID{
			Valid: true,
			UUID:  adminStream.ID,
		},
	})
	if err != nil {
		helpers.RespondWithError(w, 400, "Issue creating the stream", []string{})
		return
	}
	helpers.RespondWithJSON(w, 200, newStream)
}

func generateStreamId() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ01234567899990000"
	var builder strings.Builder
	for i := 0; i < 9; i++ {
		builder.WriteByte(charset[indexGiver(len(charset))])
		if i == 2 || i == 5 {
			builder.WriteByte('-')
		}
	}
	return builder.String()
}

func indexGiver(max int) int {
	randomNumber := rand.Intn(max)
	return randomNumber
}
