package validator

import (
	"errors"
	"regexp"
)

func ValidatePhoneNumber(phone string) error {
	re := regexp.MustCompile(`^\+(?:[0-9]-? ?){6,14}[0-9]$`)

	if !re.MatchString(phone) {
		return errors.New("phone number must contains country code")
	}

	return nil

}
