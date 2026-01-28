package utils

const (
	USD = "USD"
	EUR = "EUR"
	CFA = "CFA"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, CFA:
		return true

	}
	return false
}
