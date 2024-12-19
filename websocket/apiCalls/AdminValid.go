package apicalls

import (
	"bytes"
	"fmt"
	"net/http"
)

type AuthenticateAdmin struct {
	Secret   string
	StreamId string
	AdminId  string
}

func AuthenticateAdminFunction(authenticateVars AuthenticateAdmin) bool {
	url := "http://localhost:8000/api/v1/stream/admin-validate"
	jsonPayload := []byte("{\"secret\":\"" + authenticateVars.Secret + "\",\"streamId\":\"" + authenticateVars.StreamId + "\",\"adminId\":\"" + authenticateVars.AdminId + "\"}")
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
