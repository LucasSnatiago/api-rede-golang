package main

import (
	"fmt"
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
		DistributorAffiliation: 123123,
	}

	fmt.Println(cartaoTeste.ToJson())
}