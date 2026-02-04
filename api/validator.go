package api

import (
	"github.com/go-playground/validator/v10"
	"tutorial.sqlc.dev/app/utils"
)

// validCurrency is a custom validator for the 'currency' tag
var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		// check if currency is supported
		return utils.IsSupportedCurrency(currency)
	}
	return false
}
