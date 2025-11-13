package main

import (
	"02-workspace/utils"
	"fmt"
)

func main(){
	fmt.Println("Demo package")
	greeting:= utils.SayHello("gip")
	fmt.Println(greeting)
	fmt.Println()

	fmt.Println("Demo Greet")
	greetingDynamic:= utils.Greetings("Morning", "gip")
	fmt.Println(greetingDynamic)
	fmt.Println()

	fmt.Println("Parameter Sprintf")
	var sfString = fmt.Sprintf("String: %s", "Gip")
	fmt.Println(sfString) // String: Gip

	var sfNumber = fmt.Sprintf("Number: %d", 100)
	fmt.Println(sfNumber) // Number: 100

	var sfFloat = fmt.Sprintf("Float: %.2f", 3.14159)
	fmt.Println(sfFloat) // Float: 3.14

	fmt.Println()
}