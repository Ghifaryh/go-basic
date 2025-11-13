package utils

import "fmt"

func SayHello(name string) string {
	// return "Hello, " + name + "!"
	// or
	return fmt.Sprintf("Hellow, %s!", name)
}

func Greetings(timeOfDay, name string ) string {
	return greet(timeOfDay, name)
}

func greet(timeOfDay, name string) string {
	return fmt.Sprintf("Good %s, %s!", timeOfDay, name)
}