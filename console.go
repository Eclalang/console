package console

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strconv"
)

// Colors constants
var (
	// Reset
	ColorOff = "\033[0m" // Text Reset

	// Regular
	BLACK   = "\033[0;30m" // Black
	RED     = "\033[0;31m" // Red
	GREEN   = "\033[0;32m" // Green
	YELLOW  = "\033[0;33m" // Yellow
	BLUE    = "\033[0;34m" // Blue
	MAGENTA = "\033[0;35m" // Magenta
	CYAN    = "\033[0;36m" // Cyan
	WHITE   = "\033[0;37m" // White

	// Bold
	BBlack   = "\033[1;30m" // Black
	BRed     = "\033[1;31m" // Red
	BGreen   = "\033[1;32m" // Green
	BYellow  = "\033[1;33m" // Yellow
	BBlue    = "\033[1;34m" // Blue
	BMagenta = "\033[1;35m" // Magenta
	BCyan    = "\033[1;36m" // Cyan
	BWhite   = "\033[1;37m" // White

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

// Clear clears the console
func Clear() {
	if runtime.GOOS == "windows" {
		fmt.Print("\033[H\033[2J")
	} else {
		fmt.Print("\033[2J")
	}
}

// Input takes user input from console and returns a string
func Input() string {
	var str string
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		str = scanner.Text()
	}

	return str
}

// InputFloat takes user input from console and returns a float64
func InputFloat() (float64, error) {
	var str string
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		str = scanner.Text()
	}

	input, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0, err
	}
	return input, nil
}

// InputInt takes user input from console and returns an int
func InputInt() (int, error) {
	var str string
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		str = scanner.Text()
	}

	input, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}
	return input, nil
}

// Print prints args to console
func Print(args ...interface{}) {
	fmt.Print(args...)
}

// Printf prints args to console formatted with format
func Printf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

// PrintInColor prints args to console in color
func PrintInColor(color string, args ...interface{}) {
	fmt.Print(color)
	fmt.Print(args...)
	fmt.Print(ColorOff)
}

// Println prints args to console with a newline
func Println(args ...interface{}) {
	fmt.Println(args...)
}
