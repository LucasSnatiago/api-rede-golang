package requests

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	errors2 "github.com/lusantisuper/api-rede-golang/apierr"
	"github.com/lusantisuper/api-rede-golang/login"
	"github.com/lusantisuper/api-rede-golang/structs"
)

// Pay a method to do the payment
func Pay(r *structs.Request, login *login.Login) {
	postParameters, err := r.ToJSON()
	if err != nil {
		errors2.APIErr(err.Error())
	}

	DoPostRequest(structs.APIBaseURL(), "POST", postParameters, login)
}

// DoPostRequest do the low level needs for the requests
func DoPostRequest(url string, method string, content []byte, login *login.Login) (string, string, string) {

	var jsonStr = content
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonStr))
	req.SetBasicAuth(login.PV, login.IntegrationKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	return resp.Status, fmt.Sprint(resp.Header), string(body)
}
