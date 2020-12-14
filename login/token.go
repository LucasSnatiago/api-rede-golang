package login

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Login all needed information to login the Rede
type Login struct {
	// PV filiation number
	PV string `json:"PV"`

	// IntegrationKey your integration's key
	IntegrationKey string `json:"IntegrationKey"`
}

// ReadLoginFromJSON Reading all login information about your Rede API
func ReadLoginFromJSON() string {
	loginFile, err := os.Open("./login.json")
	if err != nil {
		fmt.Println("The login.json file doesn't exist, please copy from the github again!")
	}

	apiLogin, err := ioutil.ReadAll(loginFile)
	if err != nil {
		fmt.Println("You put something wrong on the login's file!")
	}

	var login Login
	json.Unmarshal(apiLogin, &login)

	return login.fromLoginToBase64()
}

func (l Login) fromLoginToBase64() string {
	keys := fmt.Sprintf("%s:%s", l.PV, l.IntegrationKey)

	return base64.StdEncoding.EncodeToString([]byte(keys))
}

// String a ToString function for the Login struct
func (l Login) String() string {
	return fmt.Sprintf("PV: %s\nIntegrationKey: %s\n", l.PV, l.IntegrationKey)
}
