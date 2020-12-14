package main

import (
	"github.com/lusantisuper/api-rede-golang/login"
	"github.com/lusantisuper/api-rede-golang/requests"
	"github.com/lusantisuper/api-rede-golang/testAPI"
)

func main() {
	user := login.ReadLoginFromJSON()

	testRequest(user)
}

func testRequest(login *login.Login) {
	cartaoTeste := testAPI.ReturnACardModel()

	requests.Pay(cartaoTeste, login)
}
