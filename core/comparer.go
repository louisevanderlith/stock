package core

import (
	"fmt"
)

// PriceInBounds will check if the requested price is within range of the recommended price
func PriceInBounds(requestedPrice int64, recommendedPrice int64) error {
	fraction := getFraction(recommendedPrice)
	variance := getVariance(recommendedPrice, fraction)
	upperLimit := getUpperLimit(recommendedPrice, variance)
	lowerLimit := getLowerLimit(recommendedPrice, variance)

	if requestedPrice > upperLimit || requestedPrice < lowerLimit {
		return fmt.Errorf("requested Price isn't between %d and %d", lowerLimit, upperLimit)
	}

	return nil
}

func getFraction(price int64) int64 {
	fraction := int64(2)
	variance := price

	for variance > 10 {
		variance = (variance / 3) * 2
		fraction++
	}

	return fraction
}

func getVariance(price int64, fraction int64) int64 {
	return price / fraction
}

func getUpperLimit(price int64, variance int64) int64 {
	return price + variance
}

func getLowerLimit(price int64, variance int64) int64 {
	return price - variance
}
