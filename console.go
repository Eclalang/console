package console

import (
	"fmt"
	"runtime"
	"strings"
)

var (
	// Reset
	Color_Off = "\033[0m" // Text Reset

	// Regular Colors
	BLACK   = "\033[0;30m" // Black
	RED     = "\033[0;31m" // Red
	GREEN   = "\033[0;32m" // Green
	YELLOW  = "\033[0;33m" // Yellow
	BLUE    = "\033[0;34m" // Blue
	MAGENTA = "\033[0;35m" // Magenta
	CYAN    = "\033[0;36m" // Cyan
	WHITE   = "\033[0;37m" // White

	// Bold
	BBlack  = "\033[1;30m" // Black
	BRed    = "\033[1;31m" // Red
	BGreen  = "\033[1;32m" // Green
	BYellow = "\033[1;33m" // Yellow
	BBlue   = "\033[1;34m" // Blue
	BPurple = "\033[1;35m" // Magenta
	BCyan   = "\033[1;36m" // Cyan
	BWhite  = "\033[1;37m" // White

	// Underline
	UBlack   = "\033[4;30m" // Black
	URed     = "\033[4;31m" // Red
	UGreen   = "\033[4;32m" // Green
	UYellow  = "\033[4;33m" // Yellow
	UBlue    = "\033[4;34m" // Blue
	UMagenta = "\033[4;35m" // Magenta
	UCyan    = "\033[4;36m" // Cyan
	UWhite   = "\033[4;37m" // White

	// Background
	On_Black   = "\033[40m" // Black
	On_Red     = "\033[41m" // Red
	On_Green   = "\033[42m" // Green
	On_Yellow  = "\033[43m" // Yellow
	On_Blue    = "\033[44m" // Blue
	On_Magenta = "\033[45m" // Magenta
	On_Cyan    = "\033[46m" // Cyan
	On_White   = "\033[47m" // White

	// High Intensity
	IBlack   = "\033[0;90m" // Black
	IRed     = "\033[0;91m" // Red
	IGreen   = "\033[0;92m" // Green
	IYellow  = "\033[0;93m" // Yellow
	IBlue    = "\033[0;94m" // Blue
	IMagenta = "\033[0;95m" // Magenta
	ICyan    = "\033[0;96m" // Cyan
	IWhite   = "\033[0;97m" // White

	// Bold High Intensity
	BIBlack   = "\033[1;90m" // Black
	BIRed     = "\033[1;91m" // Red
	BIGreen   = "\033[1;92m" // Green
	BIYellow  = "\033[1;93m" // Yellow
	BIBlue    = "\033[1;94m" // Blue
	BIMagenta = "\033[1;95m" // Magenta
	BICyan    = "\033[1;96m" // Cyan
	BIWhite   = "\033[1;97m" // White

	// High Intensity backgrounds
	On_IBlack   = "\033[0;100m" // Black
	On_IRed     = "\033[0;101m" // Red
	On_IGreen   = "\033[0;102m" // Green
	On_IYellow  = "\033[0;103m" // Yellow
	On_IBlue    = "\033[0;104m" // Blue
	On_IMagenta = "\033[0;105m" // Magenta
	On_ICyan    = "\033[0;106m" // Cyan
	On_IWhite   = "\033[0;107m" // White

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
	Print(args...)
	var input string
	fmt.Scanln(&input)
	return input
}

// Function to read Int input
func InputInt(prompt string) int {
	var input int
	fmt.Print(prompt)
	fmt.Scanln(&input)
	return input
}

// Function to read Float input
func InputFloat(prompt string) float64 {
	var input float64
	fmt.Print(prompt)
	fmt.Scanln(&input)
	return input
}

// Function to get user confirmation
func Confirm(prompt string) bool {
	var response string
	fmt.Print(prompt + " [y/n]: ")
	fmt.Scanln(&response)
	return strings.ToLower(response) == "y" || response == ""
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

func ProgressHotBar(percent int) {
	barLength := 20
	numBars := int(float64(percent) / 100 * float64(barLength))

	// Set the color based on the progress percentage
	var color string
	if percent >= 90 {
		color = GREEN // green
	} else if percent >= 50 {
		color = YELLOW // yellow
	} else {
		color = RED // red
	}
	fmt.Print(color)
	fmt.Print("[")
	for i := 0; i < barLength; i++ {
		if i < numBars {
			fmt.Print("=")
		} else {
			fmt.Print(" ")
		}
	}
	fmt.Printf("] %d%%%s\n", percent, "\033[0m") // Reset color at the end
}

func ProgressBarInColor(percent int, color string) {
	barLength := 20
	numBars := int(float64(percent) / 100 * float64(barLength))

	// Set the color based on the progress percentage
	switch color {
	case "red":
		color = RED
	case "green":
		color = GREEN
	case "yellow":
		color = YELLOW
	case "blue":
		color = BLUE
	case "magneta":
		color = MAGENTA
	case "cyan":
		color = CYAN
	case "white":
		color = WHITE
	case "black":
		color = BLACK
	}

	fmt.Print(color)
	fmt.Print("[")
	for i := 0; i < barLength; i++ {
		if i < numBars {
			fmt.Print("=")
		} else {
			fmt.Print(" ")
		}
	}
	fmt.Printf("] %d%%%s\n", percent, "\033[0m") // Reset color at the end
}

// Function to clear console
func Clear() {
	if runtime.GOOS == "windows" {
		fmt.Print("\033[H\033[2J")
	} else {
		fmt.Print("\033[2J")
	}
}

func PrintInColor(color string, args ...interface{}) {
	switch color {
	case "red":
		fmt.Print(RED)
	case "green":
		fmt.Print(GREEN)
	case "yellow":
		fmt.Print(YELLOW)
	case "blue":
		fmt.Print(BLUE)
	case "magenta":
		fmt.Print(MAGENTA)
	case "cyan":
		fmt.Print(CYAN)
	case "white":
		fmt.Print(WHITE)
	case "reset":
		fmt.Print(Color_Off)
	}
	fmt.Print(args...)
	fmt.Print(Color_Off)
}
