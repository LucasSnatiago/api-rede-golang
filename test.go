package main

import (
	"fmt"
	"strconv"

	"github.com/lusantisuper/api-rede-golang/login"
	"github.com/lusantisuper/api-rede-golang/requests"
	"github.com/lusantisuper/api-rede-golang/testapi"
)

// If you are testing remember to add 2 in the testapi.creditCard
func main() {
	user := login.ReadLoginFromJSON()

	testRequest(user)
	testIfACardIsValid(user)
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

func testIfACardIsValid(login *login.Login) {
	cartaoTeste := testapi.ReturnACardModel()
	newReference, _ := strconv.ParseInt(cartaoTeste.Reference, 0, 0)
	cartaoTeste.Reference = fmt.Sprint(newReference + 1)

	response, err := requests.TestCard(cartaoTeste, login, false)

	if err != nil {
		fmt.Println("The payment was not successful!")

	} else {
		fmt.Println("The payment worked!")

	}

	fmt.Printf("\n\n")
	fmt.Println(response)
}
