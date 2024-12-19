package apicalls

import (
	"bytes"
	"fmt"
	"net/http"
)

type AuthenticateStruct struct {
	Token  string
	UserId string
}

func AuthenticateUser(authenticateVars AuthenticateStruct) bool {
	url := "http://localhost:8000/api/v1/user/user-validate"
	jsonPayload := []byte("{\"token\":\"" + authenticateVars.Token + "\",\"userId\":\"" + authenticateVars.UserId + "\"}")

	response, err := http.Post(url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		fmt.Println("Error making the POST request:", err)
		return false
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		return true
	} else {
		return false
	}
}
