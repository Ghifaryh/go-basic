package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("==============================")
	fmt.Println("    Belajar error Handling    ")
	fmt.Println("==============================")

	basicError()
	customError()
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

func basicError() {
	fmt.Println("1. Basic")
	fmt.Println("--------")

	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	_, err = divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}

func customError() {
	fmt.Println()
	fmt.Println("2. Custom Error")
	fmt.Println("---------------")

	err := validateAge(200)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Age is valid")
	}

	err = validateAge(25)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Age is valid")
	}
}

type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("Validasi Gagal pada field '%s' : '%s'", e.Field, e.Message)
}

func validateAge(age int) error {
	if age < 0 {
		return &ValidationError{
			Field:   "age",
			Message: "age must be more than 0",
		}
	}

	if age > 150 {
		return &ValidationError{
			Field:   "age",
			Message: "age must be less than 150",
		}
	}
	return nil
}
