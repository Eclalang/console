package console

import (
	"bufio"
	"fmt"
	"github.com/Eclalang/Ecla/interpreter/eclaType"
	"io"
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

	// Map of variables for Ecla
	Variables = map[string]eclaType.Type{
		"COLOROFF": eclaType.String(ColorOff),
		// Regular colors
		"BLACK":   eclaType.String(BLACK),
		"RED":     eclaType.String(RED),
		"GREEN":   eclaType.String(GREEN),
		"YELLOW":  eclaType.String(YELLOW),
		"BLUE":    eclaType.String(BLUE),
		"MAGENTA": eclaType.String(MAGENTA),
		"CYAN":    eclaType.String(CYAN),
		"WHITE":   eclaType.String(WHITE),
		// Bold colors
		"B_BLACK":   eclaType.String(BBlack),
		"B_RED":     eclaType.String(BRed),
		"B_GREEN":   eclaType.String(BGreen),
		"B_YELLOW":  eclaType.String(BYellow),
		"B_BLUE":    eclaType.String(BBlue),
		"B_MAGENTA": eclaType.String(BMagenta),
		"B_CYAN":    eclaType.String(BCyan),
		"B_WHITE":   eclaType.String(BWhite),
		// Underline colors
		"U_BLACK":   eclaType.String(UBlack),
		"U_RED":     eclaType.String(URed),
		"U_GREEN":   eclaType.String(UGreen),
		"U_YELLOW":  eclaType.String(UYellow),
		"U_BLUE":    eclaType.String(UBlue),
		"U_MAGENTA": eclaType.String(UMagenta),
		"U_CYAN":    eclaType.String(UCyan),
		"U_WHITE":   eclaType.String(UWhite),
		// Background colors
		"ON_BLACK":   eclaType.String(On_Black),
		"ON_RED":     eclaType.String(On_Red),
		"ON_GREEN":   eclaType.String(On_Green),
		"ON_YELLOW":  eclaType.String(On_Yellow),
		"ON_BLUE":    eclaType.String(On_Blue),
		"ON_MAGENTA": eclaType.String(On_Magenta),
		"ON_CYAN":    eclaType.String(On_Cyan),
		"ON_WHITE":   eclaType.String(On_White),
		// High Intensity colors
		"I_BLACK":   eclaType.String(IBlack),
		"I_RED":     eclaType.String(IRed),
		"I_GREEN":   eclaType.String(IGreen),
		"I_YELLOW":  eclaType.String(IYellow),
		"I_BLUE":    eclaType.String(IBlue),
		"I_MAGENTA": eclaType.String(IMagenta),
		"I_CYAN":    eclaType.String(ICyan),
		"I_WHITE":   eclaType.String(IWhite),
		// Bold High Intensity colors
		"BI_BLACK":   eclaType.String(BIBlack),
		"BI_RED":     eclaType.String(BIRed),
		"BI_GREEN":   eclaType.String(BIGreen),
		"BI_YELLOW":  eclaType.String(BIYellow),
		"BI_BLUE":    eclaType.String(BIBlue),
		"BI_MAGENTA": eclaType.String(BIMagenta),
		"BI_CYAN":    eclaType.String(BICyan),
		"BI_WHITE":   eclaType.String(BIWhite),
		// High Intensity background colors
		"ONI_BLACK":   eclaType.String(On_IBlack),
		"ONI_RED":     eclaType.String(On_IRed),
		"ONI_GREEN":   eclaType.String(On_IGreen),
		"ONI_YELLOW":  eclaType.String(On_IYellow),
		"ONI_BLUE":    eclaType.String(On_IBlue),
		"ONI_MAGENTA": eclaType.String(On_IMagenta),
		"ONI_CYAN":    eclaType.String(On_ICyan),
		"ONI_WHITE":   eclaType.String(On_IWhite),
	}
)

// For testing purposes
var read io.Reader = os.Stdin

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
	scanner := bufio.NewScanner(read)

	if scanner.Scan() {
		str = scanner.Text()
	}

	return str
}

// InputFloat takes user input from console and returns a float64
func InputFloat() (float64, error) {
	var str string
	scanner := bufio.NewScanner(read)

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
	scanner := bufio.NewScanner(read)

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
