package validator

import (
	"errors"
	"hello-nurse/src/constants"
)

const (
	IT    string = constants.ROLE_IT
	NURSE string = constants.ROLE_NURSE
)

func ValidateCategory(status string) error {
	switch status {
	case IT,
		NURSE:

		return nil
	default:
		return errors.New("invalid Category Role, please check your Role")
	}
}
