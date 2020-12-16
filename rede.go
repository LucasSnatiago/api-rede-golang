package rede

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	errors2 "github.com/lusantisuper/api-rede-golang/apierrs"
	"github.com/lusantisuper/api-rede-golang/utils"

	"github.com/lusantisuper/api-rede-golang/login"
	"github.com/lusantisuper/api-rede-golang/models"
)

// Rede interface for the Rede's API
type Rede interface {
	Pay(p *models.Payment) (*models.Response, error)
	TestCard(p *models.Payment) (*models.Response, error)
}

type rede struct {
	config *login.Login
}

// NewRede instantiate a new Rede API object
func (r rede) NewRede(pv string, ik string, isProduction bool) *rede {
	return &rede{
		config: &login.Login{
			PV:             pv,
			IntegrationKey: ik,
			IsProduction:   isProduction,
		},
	}
}

// Pay a method to do the payment
func (r rede) Pay(req *models.Payment) (*models.Response, error) {
	postParameters, err := req.ToJSON()
	if err != nil {
		errors2.APIErr(err.Error())
	}

	body := ""
	if r.config.IsProduction {
		_, _, body = doPostRequest(utils.APIBaseURL(), "POST", postParameters, r.config)

	} else {
		_, _, body = doPostRequest(utils.APIBaseURLTest(), "POST", postParameters, r.config)

	}

	var parseHeader models.Response
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
func (r rede) TestCard(req *models.Payment) (*models.Response, error) {
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
