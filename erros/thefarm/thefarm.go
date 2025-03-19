package thefarm

import (
	"errors"
	"fmt"
)

// FodderCalculator interface
type FodderCalculator interface {
	FodderAmount(int) (float64, error)
	FatteningFactor() (float64, error)
}

type InvalidCowsError struct {
	cows          int
	customMessage string
}

func (e *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", e.cows, e.customMessage)
}

// DivideFood calculates the amount of food per cow
func DivideFood(fc FodderCalculator, cows int) (float64, error) {
	// Get total fodder amount
	totalFodder, err := fc.FodderAmount(cows)
	if err != nil {
		return 0, err
	}

	// Get fattening factor
	factor, err := fc.FatteningFactor()
	if err != nil {
		return 0, err
	}

	// Calculate and return food per cow
	return (totalFodder * factor) / float64(cows), nil
}
func ValidateInputAndDivideFood(fc FodderCalculator, cows int) (float64, error) {
	if cows <= 0 {
		return 0, errors.New("invalid number of cows")
	}
	return DivideFood(fc, cows)
}

func ValidateNumberOfCows(cows int) error {
	if cows < 0 {
		return &InvalidCowsError{
			cows:          cows,
			customMessage: "there are no negative cows",
		}

	} else if cows == 0 {
		return &InvalidCowsError{
			cows:          cows,
			customMessage: "no cows don't need food",
		}
	}
	return nil
}

// Example usage
type ExampleFodderCalculator struct{}

func (e ExampleFodderCalculator) FodderAmount(cows int) (float64, error) {
	return 50, nil // Example value
}

func (e ExampleFodderCalculator) FatteningFactor() (float64, error) {
	return 1.5, nil // Example value
}

// func main() {
// 	fc := ExampleFodderCalculator{}
// 	result, err := DivideFood(fc, 5)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 	} else {
// 		fmt.Println("Food per cow:", result)
// 	}
// }
