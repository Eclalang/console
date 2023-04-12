package main

import (
	console "github.com/Eclalang/console"
)

func main() {
	console.Printf("Hello, %s!\n", "world")
	console.Println("This is a test message")
	input := console.Input("Enter your name: ")
	console.Printf("Hello, %s!\n", input)
	number := console.InputInt("Enter a number: ")
	console.Printf("The number you entered is %d\n", number)
	confirmed := console.Confirm("Are you sure you want to continue?")
	if confirmed {
		console.Println("Great! Let's continue.")
	} else {
		console.Println("Okay, let's stop here.")
	}
	console.ProgressBar(50)
	console.PrintInColor("red", "This is a red message\n")
	console.PrintInColor("green", "This is a green message\n")
	confirmed = console.Confirm("Do you want to clear the console?")
	if confirmed {
		console.Clear()
	}
}
