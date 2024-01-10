package main

import (
	"fmt"
	"github.com/Eclalang/console"
)

func main() {
	TestInputFunctions()
}

func TestInputFunctions() {
	fmt.Println("Test for all Input functions")
	fmt.Println("----------------------------------------")

	fmt.Println("Test function Input()")
	fmt.Print("Write your input: ")
	input := console.Input()
	fmt.Printf("Your input is: %s \n", input)

	fmt.Println("----------------------------------------")

	fmt.Println("Test function InputFloat()")
	fmt.Print("Write your input: ")
	inputFloat, err := console.InputFloat()
	if err != "" {
		fmt.Printf("Error: %s \n", err)
	}
	fmt.Printf("Your input is: %f \n", inputFloat)

	fmt.Println("----------------------------------------")

	fmt.Println("Test function InputInt()")
	fmt.Print("Write your input: ")
	inputInt, err := console.InputInt()
	if err != "" {
		fmt.Printf("Error: %s \n", err)
	}
	fmt.Printf("Your input is: %d", inputInt)
}
