package console

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
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

// Function to read Int input
func InputInt(prompt string) int {
	var input int
	fmt.Print(prompt)
	fmt.Scan(&input)
	return input
}

// Function to read Float input
func InputFloat(prompt string) float64 {
	var input float64
	fmt.Print(prompt)
	fmt.Scan(&input)
	return input
}

// Function to get user confirmation
func Confirm(prompt string) bool {
	var response string
	fmt.Print(prompt + " [y/n]: ")
	fmt.Scan(&response)
	return strings.ToLower(response) == "y"
}

// Function to display a progress bar
func ProgressBar(percent int) {
	barLength := 20
	numBars := int(float64(percent) / 100 * float64(barLength))
	fmt.Print("[")
	for i := 0; i < barLength; i++ {
		if i < numBars {
			fmt.Print("=")
		} else {
			fmt.Print(" ")
		}
	}
	fmt.Printf("] %d%%\n", percent)
}

// Function to clear console
func Clear() {
	cmd := exec.Command("clear")
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
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
