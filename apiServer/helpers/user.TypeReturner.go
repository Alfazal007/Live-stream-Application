package helpers

import "github.com/Alfazal007/apiserver/internal/database"

type ReturnUser struct {
	Username string `json:"username"`
	Id       string `json:"id"`
}

func GenerateUserToReturn(user database.User) ReturnUser {
	return ReturnUser{
		Username: user.Username,
		Id:       user.ID.String(),
	}
}
