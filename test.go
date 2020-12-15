package main

import (
	"fmt"

	"github.com/lusantisuper/api-rede-golang/login"
	"github.com/lusantisuper/api-rede-golang/requests"
	"github.com/lusantisuper/api-rede-golang/testapi"
)

func main() {
	user := login.ReadLoginFromJSON()

	testRequest(user)
}

func testRequest(login *login.Login) {
	cartaoTeste := testapi.ReturnACardModel()

	response, err := requests.Pay(cartaoTeste, login, false)

	if err != nil {
		fmt.Println("The payment was not successful!")

	} else {
		fmt.Println("The payment worked!")

	}

	fmt.Printf("\n\n")
	fmt.Println(response)
}