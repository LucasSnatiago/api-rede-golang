package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/lusantisuper/api-rede-golang/internal"
	"io/ioutil"
	"net/http"
	"strings"

	errors2 "github.com/lusantisuper/api-rede-golang/internal/apierr"
	"github.com/lusantisuper/api-rede-golang/internal/login"
)

// Rede interface for the Rede's API
type Rede interface {
	Pay(r *Payment) (*Response, error)
	TestCard(r *Payment) (*Response, error)
}

type rede struct {
	config *login.Login
}

// Pay a method to do the payment
func (r rede) Pay(req *Payment) (*Response, error) {
	postParameters, err := req.toJSON()
	if err != nil {
		errors2.APIErr(err.Error())
	}

	body := ""
	if r.config.IsProduction {
		_, _, body = doPostRequest(internal.APIBaseURL(), "POST", postParameters, r.config)

	} else {
		_, _, body = doPostRequest(internal.APIBaseURLTest(), "POST", postParameters, r.config)

	}

	var parseHeader Response
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
func (r rede) TestCard(req *Payment) (*Response, error) {
	req.Amount = 0

	payment, err := r.Pay(req)
	if err != nil {
		err = errors2.APIErr("The card is not valid!")
	}

	return payment, err
}

// doPostRequest do the low level needs for the requests
func doPostRequest(url string, method string, content []byte, login *login.Login) (string, string, string) {

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
