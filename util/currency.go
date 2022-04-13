package util

var currencyList = []string{"USD", "EUR", "CAD"}

func IsSupportedCurrency(currency string) bool {
	for _, v := range currencyList {
		if currency == v {
			return true
		}
	}
	return false
}
