package util

// Constants for all supported currencies
const (
	ETH = "以太坊"
	BTC = "比特币"
)

// IsSupportedCurrency 是否支持该币种
func IsSupportedCurrency(currency string) bool {
	switch currency {
	case ETH, BTC:
		return true
	}
	return false
}
