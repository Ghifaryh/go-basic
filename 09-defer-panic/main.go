package main

import "fmt"

func main() {
	defer fmt.Println("World!")

	fmt.Println("Hello")

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered From", r)
		}
	}()

	panic("Ayo Panic!") // like an instant stoper

	defer fmt.Println("One")
	defer fmt.Println("Two")
	defer fmt.Println("Three")

}
