package utilitiesCalculate

import (
	"fmt"
	"kkl-v2/core/logs"
	"math"
)

// CalculateDiscount calculates the discounted price given the original price and discount percentage.
func CalculateDiscount(originalPrice float64, discountPercentage float64) (float64, float64) {
	discount := originalPrice * (discountPercentage / 100)
	discountedPrice := originalPrice - discount
	return discount, discountedPrice
}

// CalculateDiscount calculates the discounted price based on the original price and discount percentage.
func CalculatDisCount(originalPrice float64, discountPercentage float64) float64 {
	logs.Info(fmt.Sprintf("[CalculatDisCount] - originalPrice=%v", originalPrice))
	logs.Info(fmt.Sprintf("[CalculatDisCount] - discountPercentage=%v", discountPercentage))
	discountedPrice := originalPrice - (originalPrice * discountPercentage / 100)

	return discountedPrice
}

// CalculateVAT calculates the VAT amount given the amount and VAT rate
func CalculateVAT(amount float64, vatRate float64) float64 {
	logs.Info(fmt.Sprintf("[CalculateVAT] - amount=%v", amount))
	logs.Info(fmt.Sprintf("[CalculateVAT] - vatRate=%v", vatRate))
	return amount * vatRate / 100
}

// CalculateTotalSummary calculates the VAT amount given the amount and VAT rate
func CalculateTotalSummary(totalAmount float64, vatAmount float64) float64 {
	logs.Info(fmt.Sprintf("[CalculateTotalSummary] - totalAmount=%v", totalAmount))
	logs.Info(fmt.Sprintf("[CalculateTotalSummary] - vatAmount=%v", vatAmount))
	return totalAmount + vatAmount
}

// Calculate the exchanged amount
func ExchangeLAK(amount float64, rate float64) float64 {
	exchangedAmount := amount / rate
	return truncateDecimal(exchangedAmount, 2)
}

func truncateDecimal(num float64, precision int) float64 {
	shift := math.Pow(10, float64(precision))
	return math.Floor(num*shift) / shift
}
