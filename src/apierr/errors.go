package errors

import "errors"

const INSUFFICIENTPARAMETERS string = "Some necessary parameters are missing!"
const WRONGAMOUNT string = "The amount passed is out of range!"
const WRONGDATENUMBER string = "The date passed is out of range!"

func ApiErr(s string) error {
	return errors.New(s)
}