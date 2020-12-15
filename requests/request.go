package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	errors2 "github.com/lusantisuper/api-rede-golang/apierr"
	"github.com/lusantisuper/api-rede-golang/login"
	"github.com/lusantisuper/api-rede-golang/structs"
)

// Pay a method to do the payment
func Pay(r *structs.Payment, login *login.Login, isRealTransaction bool) (*structs.Response, error) {
	postParameters, err := r.ToJSON()
	if err != nil {
		errors2.APIErr(err.Error())
	}

	body := ""
	if isRealTransaction {
		_, _, body = DoPostRequest(structs.APIBaseURL(), "POST", postParameters, login)

	} else {
		_, _, body = DoPostRequest(structs.APIBaseURLTest(), "POST", postParameters, login)

	}

	var parseHeader structs.Response
	err = json.Unmarshal([]byte(body), &parseHeader)
	if err != nil {
		errors2.APIErr(err.Error())
	}

	if strings.Compare(parseHeader.ReturnCode, "00") != 0 && strings.Compare(parseHeader.ReturnCode, "174") != 0 {
		err = errors2.APIErr("The payment was not successful!")
	}

	return &parseHeader, err
}

// TestCard is a test function to see if the card is valid
func TestCard(r *structs.Payment, login *login.Login, isRealTransaction bool) (*structs.Response, error) {
	r.Amount = 0

	payment, err := Pay(r, login, isRealTransaction)
	if err != nil {
		err = errors2.APIErr("The card is not valid!")
	}

	return payment, err
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

	body, _ := ioutil.ReadAll(resp.Body)

	return resp.Status, fmt.Sprint(resp.Header), string(body)
}
