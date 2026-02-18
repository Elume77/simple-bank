package utils

const (
	USD = "USD"
	EUR = "EUR"
	CFA = "CFA"
)

// IsSupportedCurrency returns true if the currency is supported
func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, CFA:
		return true
	}
	return false
}
