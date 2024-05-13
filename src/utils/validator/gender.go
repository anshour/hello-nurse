package validator

import (
	"errors"
)

const (
	MALE   string = "male"
	FEMALE string = "female"
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
