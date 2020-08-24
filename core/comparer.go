package core

import (
	"fmt"
)

// PriceInBounds will check if the requested price is within range of the recommended price
func PriceInBounds(requestedPrice uint64, recommendedPrice uint64) error {
	fraction := getFraction(recommendedPrice)
	variance := getVariance(recommendedPrice, fraction)
	upperLimit := getUpperLimit(recommendedPrice, variance)
	lowerLimit := getLowerLimit(recommendedPrice, variance)

	if requestedPrice > upperLimit || requestedPrice < lowerLimit {
		return fmt.Errorf("requested Price isn't between %d and %d", lowerLimit, upperLimit)
	}

	return nil
}

func getFraction(price uint64) uint64 {
	fraction := uint64(2)
	variance := price

	for variance > 10 {
		variance = (variance / 3) * 2
		fraction++
	}

	return fraction
}

func getVariance(price uint64, fraction uint64) uint64 {
	return price / fraction
}

func getUpperLimit(price uint64, variance uint64) uint64 {
	return price + variance
}

func getLowerLimit(price uint64, variance uint64) uint64 {
	return price - variance
}
