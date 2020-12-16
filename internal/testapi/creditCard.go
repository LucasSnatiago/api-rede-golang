package testapi

import (
	"github.com/lusantisuper/api-rede-golang"
)

// ReturnACardModel example credit card
func ReturnACardModel() *main.Payment {
	return &main.Payment{
		Capture:                true,
		Kind:                   "credit",
		Reference:              "0",
		Amount:                 1000,
		Installments:           0,
		CardHolderName:         "Pessoa 2",
		CardNumber:             5448280000000007,
		ExpirationMonth:        1,
		ExpirationYear:         21,
		SecurityCode:           123,
		SoftDescriptor:         "Teste",
		Subscription:           false,
		DistributorAffiliation: 10007281,
	}
}
