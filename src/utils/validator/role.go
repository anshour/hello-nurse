package validator

import (
	"errors"
)

const (
	IT    string = "It"
	NURSE string = "Nurse"
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
