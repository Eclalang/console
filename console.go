package console

import (
	"bufio"
	"fmt"
	"os"
)

// Function to print args as formated string to console
func Printf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

// Function to print args to console
func Println(args ...interface{}) {
	fmt.Println(args...)
}

// Function to print args to console withou newline
func Print(args ...interface{}) {
	fmt.Print(args...)
}

// Function that take user input from console
func Input(args ...interface{}) string {
	var input string
	fmt.Print(args...)
	reader := bufio.NewReader(os.Stdin)
	input, _ = reader.ReadString('\n')
	return input
}

func PrintInColor(color string, args ...interface{}) {
	switch color {
	case "red":
		fmt.Print("\033[31m")
	case "green":
		fmt.Print("\033[32m")
	case "yellow":
		fmt.Print("\033[33m")
	case "blue":
		fmt.Print("\033[34m")
	case "magenta":
		fmt.Print("\033[35m")
	case "cyan":
		fmt.Print("\033[36m")
	case "white":
		fmt.Print("\033[37m")
	case "reset":
		fmt.Print("\033[0m")
	}
	fmt.Print(args...)
	fmt.Print("\033[0m")
}
