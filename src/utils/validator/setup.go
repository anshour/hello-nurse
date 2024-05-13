package validator

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

var CustomValidator = &Validator{Validator: SetupValidator()}

type Validator struct {
	Validator *validator.Validate
}

func (cv *Validator) Validate(i interface{}) error {
	if err := cv.Validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func SetupValidator() *validator.Validate {
	v := validator.New()

	// v.RegisterValidation("validateCategory", validateCategory)

	return v
}

// func validateCategory(fl validator.FieldLevel) bool {
// 	allowedValues := []string{
// 		"Clothing",
// 		"Accessories",
// 		"Footwear",
// 		"Beverages",
// 	}

// 	value := fl.Field().String()
// 	for _, v := range allowedValues {
// 		if strings.EqualFold(value, v) {
// 			return true
// 		}
// 	}
// 	return false
// }
