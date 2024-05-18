package validator

import (
	"errors"
	"hello-nurse/src/constants"
)

const (
	MALE   string = constants.GENDER_MALE
	FEMALE string = constants.GENDER_FEMALE
)

func ValidateGender(status string) error {
	switch status {
	case MALE,
		FEMALE:

		return nil
	default:
		return errors.New("invalid Gender Type, please check your Gender")
	}
}
