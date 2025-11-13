package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// score := 85
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide your score")
		return
	}

	score, err := strconv.ParseFloat(arguments[1],64)

	if err != nil {
		fmt.Println(err)
	}

	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Println("Grade: B")
	} else if score >= 70 {
		fmt.Println("Grade: C")
	} else if score >= 60 {
		fmt.Println("Grade: D")
	} else {
		fmt.Println("Grade: F")
	}
}

// go run ./04-control-flow/main.go 70
