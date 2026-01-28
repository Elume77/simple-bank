package api

import (
	"github.com/go-playground/validator/v10"
	"tutorial.sqlc.dev/app/utils"
)

var validCurrency validator.Func = func(fieldlevel validator.FieldLevel) bool {
	if currency, ok := fieldlevel.Field().Interface().(string); ok {
		// checks if currency is supported or not
		return utils.IsSupportedCurrency(currency)
	}
	return false
}
