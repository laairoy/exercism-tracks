package thefarm

import (
	"errors"
	"fmt"
)

// define the 'DivideFood' function
func DivideFood(fc FodderCalculator, ncows int) (float64, error) {
  amount, err := fc.FodderAmount(ncows)
  if err != nil {
    return 0.0, err
  }
  factor, err := fc.FatteningFactor()
  if err != nil {
    return 0.0, err
  }
  return float64(amount*factor/float64(ncows)), nil
}

// define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(fc FodderCalculator, ncows int) (float64, error){
  if ncows <= 0 {
    return 0.0, errors.New("invalid number of cows") 
  }

  return DivideFood(fc, ncows)
}

// define the 'ValidateNumberOfCows' function
type NumberOfCowsError struct {
  number int
  message string
}

func (e *NumberOfCowsError) Error() string {
  return fmt.Sprintf("%d cows are invalid: %s", e.number, e.message)
}

func ValidateNumberOfCows(ncows int) error {
  if ncows < 0 {
    return &NumberOfCowsError{number: ncows, message: "there are no negative cows"}
  } else if ncows == 0 {
    return &NumberOfCowsError{number: ncows, message: "no cows don't need food"}
  }
  return nil
}
// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
