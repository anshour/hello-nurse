package validator

import (
	"errors"
	"fmt"
	"hello-nurse/src/constants"
	"strconv"
)

func ValidateNip(nip int64, nipTarget string) error {
	nipStr := strconv.FormatInt(nip, 10)
	if nipStr[0:3] != constants.NipNurse {
		return errors.New(fmt.Sprintln("Nip must start with $1", nipTarget))

	}

	return nil

}
