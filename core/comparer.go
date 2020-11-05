package core

import (
	"fmt"
)

// PriceInBounds will check if the requested price is within range of the recommended price
func PriceInBounds(requestedPrice float32, recommendedPrice float32) error {
	fraction := getFraction(recommendedPrice)
	variance := getVariance(recommendedPrice, fraction)
	upperLimit := getUpperLimit(recommendedPrice, variance)
	lowerLimit := getLowerLimit(recommendedPrice, variance)

	if requestedPrice > upperLimit || requestedPrice < lowerLimit {
		return fmt.Errorf("requested Price isn't between %d and %d", lowerLimit, upperLimit)
	}

	return nil
}

func getFraction(price float32) float32 {
	fraction := float32(2)
	variance := price

	for variance > 10 {
		variance = (variance / 3) * 2
		fraction++
	}

	return fraction
}

func getVariance(price float32, fraction float32) float32 {
	return price / fraction
}

func getUpperLimit(price float32, variance float32) float32 {
	return price + variance
}

func getLowerLimit(price float32, variance float32) float32 {
	return price - variance
}
