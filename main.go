package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/lusantisuper/api-rede-golang/structs"
)

func main() {
	cartaoTeste := &structs.Request{
		Capture:                false,
		Kind:                   "credit",
		Reference:              "123",
		Amount:                 100,
		Installments:           0,
		CardHolderName:         "Pessoa 1",
		CardNumber:             2131230,
		ExpirationMonth:        10,
		ExpirationYear:         24,
		SecurityCode:           234,
		SoftDescriptor:         "Teste de cartão de credito",
		Subscription:           false,
		DistributorAffiliation: 10007281,
	}

	var jsonStr, _ = cartaoTeste.ToJSON()
	fmt.Println(string(jsonStr))

	req, err := http.NewRequest("POST", structs.APIBaseURL()+"/transactions", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		fmt.Println(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}
