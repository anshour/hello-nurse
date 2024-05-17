package validator

import (
	"errors"
	"fmt"
	"strconv"
)

func ValidateNip(nip int64, nipTarget string) error {
	nipStr := strconv.FormatInt(nip, 10)
	if nipStr[0:3] != nipTarget {
		message := fmt.Sprint("Nip must start with ", nipTarget)
		return errors.New(message)

	}

	return nil

}
