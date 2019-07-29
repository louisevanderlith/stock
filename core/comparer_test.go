package core

import (
	"testing"
)

func TestIsValidPrice_MustBeValid(t *testing.T) {
	inputPrice := int64(21)
	recommendedPrice := int64(25)

	err := PriceInBounds(inputPrice, recommendedPrice)

	if err != nil {
		t.Error(err)
	}
}

func TestIsValidPrice_MustBeInvalid(t *testing.T) {
	inputPrice := int64(85)
	recommendedPrice := int64(25)

	err := PriceInBounds(inputPrice, recommendedPrice)

	if err != nil {
		t.Error(err)
	}
}

func TestGetFraction_ShouldBe14(t *testing.T) {
	inputPrice := int64(1296)
	expectedFraction := int64(14)

	actualFraction := getFraction(inputPrice)

	if expectedFraction != actualFraction {
		t.Errorf("Fraction was not valid, got: %d, want: %d", actualFraction, expectedFraction)
	}
}

func TestGetFraction_ShouldBe9(t *testing.T) {
	inputPrice := int64(171)
	expectedFraction := int64(9)

	actualFraction := getFraction(inputPrice)

	if expectedFraction != actualFraction {
		t.Errorf("Fraction was not valid, got: %d, want: %d", actualFraction, expectedFraction)
	}
}

func TestGetFraction_ShouldBe2(t *testing.T) {
	inputPrice := int64(7)
	expectedFraction := int64(2)

	actualFraction := getFraction(inputPrice)

	if expectedFraction != actualFraction {
		t.Errorf("Fraction was not valid, got: %d, want: %d", actualFraction, expectedFraction)
	}
}

func TestGetFraction_ShouldBe9Between(t *testing.T) {
	inputPrice := int64(245)
	expectedFraction := int64(9)

	actualFraction := getFraction(inputPrice)

	if expectedFraction != actualFraction {
		t.Errorf("Fraction was not valid, got: %d, want: %d", actualFraction, expectedFraction)
	}
}
