package console

import (
	"fmt"
	_ "os"
)

// Function to print args as formated string to console
func printf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

// Function to print args to console
func println(args ...interface{}) {
	fmt.Println(args...)
}

// Function to print args to console withou newline
func print(args ...interface{}) {
	fmt.Print(args...)
}
