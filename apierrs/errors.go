package apierrs

import "errors"

// INSUFFICIENTPARAMETERS Insuficient paremeters passed
const INSUFFICIENTPARAMETERS string = "Some necessary parameters are missing!"

// WRONGAMOUNT The money value is not correct
const WRONGAMOUNT string = "The amount passed is out of range!"

// WRONGDATENUMBER The date passed is out of the range
const WRONGDATENUMBER string = "The date passed is out of range!"

// APIErr return a error type
func APIErr(s string) error {
	return errors.New(s)
}
