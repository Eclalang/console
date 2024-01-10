package console

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestClear(t *testing.T) {
	file, err := os.OpenFile("unit_test_files/clear.txt", os.O_RDWR, 0777)
	if err != nil {
		t.Errorf("Error opening file: %v", err)
	}
	err = file.Truncate(0)
	if err != nil {
		t.Errorf("Error truncating file: %v", err)
	}
	os.Stdout = file

	Clear()

	err = os.Stdout.Sync()
	if err != nil {
		t.Errorf("Error syncing stdout: %v", err)
	}
	err = os.Stdout.Close()
	if err != nil {
		t.Errorf("Error closing stdout: %v", err)
	}

	content, err := os.ReadFile("unit_test_files/clear.txt")
	if err != nil {
		t.Errorf("Error reading file: %v", err)
	}

	expected := "\u001B[H\u001B[2J"
	if string(content) != expected {
		t.Errorf("Expected %v, got %v", expected, string(content))
	}
}

func TestPrint(t *testing.T) {
	str := "Test123"
	x := 123
	y := 123.456
	check := true

	file, err := os.OpenFile("unit_test_files/print.txt", os.O_RDWR, 0777)
	if err != nil {
		t.Errorf("Error opening file: %v", err)
	}
	err = file.Truncate(0)
	if err != nil {
		t.Errorf("Error truncating file: %v", err)
	}
	os.Stdout = file

	Print(str, "\n", x, "\n", y, "\n", check)

	err = os.Stdout.Sync()
	if err != nil {
		t.Errorf("Error syncing stdout: %v", err)
	}
	err = os.Stdout.Close()
	if err != nil {
		t.Errorf("Error closing stdout: %v", err)
	}

	content, err := os.ReadFile("unit_test_files/print.txt")
	if err != nil {
		t.Errorf("Error reading file: %v", err)
	}

	contentSplit := strings.Split(string(content), "\n")
	for i := 0; i < len(contentSplit); i++ {
		switch i {
		case 0:
			if contentSplit[i] != str {
				t.Errorf("Error Print(%s), got %s", str, contentSplit[i])
			}
		case 1:
			if contentSplit[i] != fmt.Sprintf("%d", x) {
				t.Errorf("Error Print(%d), got %s", x, contentSplit[i])
			}
		case 2:
			if contentSplit[i] != fmt.Sprintf("%g", y) {
				t.Errorf("Error Print(%g), got %s", y, contentSplit[i])
			}
		case 3:
			if contentSplit[i] != fmt.Sprintf("%t", check) {
				t.Errorf("Error Print(%t), got %s", check, contentSplit[i])
			}
		}
	}
}

func TestPrintf(t *testing.T) {
	str := "Test123"
	x := 123
	y := 123.456
	check := true

	file, err := os.OpenFile("unit_test_files/print.txt", os.O_RDWR, 0777)
	if err != nil {
		t.Errorf("Error opening file: %v", err)
	}
	err = file.Truncate(0)
	if err != nil {
		t.Errorf("Error truncating file: %v", err)
	}
	os.Stdout = file

	Printf("%s\n%d\n%g\n%t\n", str, x, y, check)

	err = os.Stdout.Sync()
	if err != nil {
		t.Errorf("Error syncing stdout: %v", err)
	}
	err = os.Stdout.Close()
	if err != nil {
		t.Errorf("Error closing stdout: %v", err)
	}

	content, err := os.ReadFile("unit_test_files/print.txt")
	if err != nil {
		t.Errorf("Error reading file: %v", err)
	}

	contentSplit := strings.Split(string(content), "\n")
	for i := 0; i < len(contentSplit); i++ {
		switch i {
		case 0:
			if contentSplit[i] != str {
				t.Errorf("Error Printf(%s), got %s", str, contentSplit[i])
			}
		case 1:
			if contentSplit[i] != fmt.Sprintf("%d", x) {
				t.Errorf("Error Printf(%d), got %s", x, contentSplit[i])
			}
		case 2:
			if contentSplit[i] != fmt.Sprintf("%g", y) {
				t.Errorf("Error Printf(%g), got %s", y, contentSplit[i])
			}
		case 3:
			if contentSplit[i] != fmt.Sprintf("%t", check) {
				t.Errorf("Error Printf(%t), got %s", check, contentSplit[i])
			}
		}
	}
}

func TestPrintInColor(t *testing.T) {
	str := "Test123\n"

	file, err := os.OpenFile("unit_test_files/print.txt", os.O_RDWR, 0777)
	if err != nil {
		t.Errorf("Error opening file: %v", err)
	}
	err = file.Truncate(0)
	if err != nil {
		t.Errorf("Error truncating file: %v", err)
	}
	os.Stdout = file

	PrintInColor(BLACK, str)
	PrintInColor(BLUE, str)
	PrintInColor(CYAN, str)
	PrintInColor(GREEN, str)
	PrintInColor(MAGENTA, str)
	PrintInColor(RED, str)
	PrintInColor(WHITE, str)
	PrintInColor(YELLOW, str)

	PrintInColor(BBlack, str)
	PrintInColor(BBlue, str)
	PrintInColor(BCyan, str)
	PrintInColor(BGreen, str)
	PrintInColor(BMagenta, str)
	PrintInColor(BRed, str)
	PrintInColor(BWhite, str)
	PrintInColor(BYellow, str)

	PrintInColor(UBlack, str)
	PrintInColor(UBlue, str)
	PrintInColor(UCyan, str)
	PrintInColor(UGreen, str)
	PrintInColor(UMagenta, str)
	PrintInColor(URed, str)
	PrintInColor(UWhite, str)
	PrintInColor(UYellow, str)

	PrintInColor(On_Black, str)
	PrintInColor(On_Blue, str)
	PrintInColor(On_Cyan, str)
	PrintInColor(On_Green, str)
	PrintInColor(On_Magenta, str)
	PrintInColor(On_Red, str)
	PrintInColor(On_White, str)
	PrintInColor(On_Yellow, str)

	PrintInColor(IBlack, str)
	PrintInColor(IBlue, str)
	PrintInColor(ICyan, str)
	PrintInColor(IGreen, str)
	PrintInColor(IMagenta, str)
	PrintInColor(IRed, str)
	PrintInColor(IWhite, str)
	PrintInColor(IYellow, str)

	PrintInColor(BIBlack, str)
	PrintInColor(BIBlue, str)
	PrintInColor(BICyan, str)
	PrintInColor(BIGreen, str)
	PrintInColor(BIMagenta, str)
	PrintInColor(BIRed, str)
	PrintInColor(BIWhite, str)
	PrintInColor(BIYellow, str)

	PrintInColor(On_IBlack, str)
	PrintInColor(On_IBlue, str)
	PrintInColor(On_ICyan, str)
	PrintInColor(On_IGreen, str)
	PrintInColor(On_IMagenta, str)
	PrintInColor(On_IRed, str)
	PrintInColor(On_IWhite, str)
	PrintInColor(On_IYellow, str)

	PrintInColor(ColorOff, str)

	err = os.Stdout.Sync()
	if err != nil {
		t.Errorf("Error syncing stdout: %v", err)
	}
	err = os.Stdout.Close()
	if err != nil {
		t.Errorf("Error closing stdout: %v", err)
	}

	content, err := os.ReadFile("unit_test_files/print.txt")
	if err != nil {
		t.Errorf("Error reading file: %v", err)
	}
	contentSplit := strings.Split(string(content), "\033[0m")
	for i := 0; i < len(contentSplit); i++ {
		switch i {
		case 0:
			if contentSplit[i] != BLACK+str {
				t.Errorf("Error PrintInColor(BLACK, %s), got %s", str, contentSplit[i])
			}
		case 1:
			if contentSplit[i] != BLUE+str {
				t.Errorf("Error PrintInColor(BLUE, %s), got %s", str, contentSplit[i])
			}
		case 2:
			if contentSplit[i] != CYAN+str {
				t.Errorf("Error PrintInColor(CYAN, %s), got %s", str, contentSplit[i])
			}
		case 3:
			if contentSplit[i] != GREEN+str {
				t.Errorf("Error PrintInColor(GREEN, %s), got %s", str, contentSplit[i])
			}
		case 4:
			if contentSplit[i] != MAGENTA+str {
				t.Errorf("Error PrintInColor(MAGENTA, %s), got %s", str, contentSplit[i])
			}
		case 5:
			if contentSplit[i] != RED+str {
				t.Errorf("Error PrintInColor(RED, %s), got %s", str, contentSplit[i])
			}
		case 6:
			if contentSplit[i] != WHITE+str {
				t.Errorf("Error PrintInColor(WHITE, %s), got %s", str, contentSplit[i])
			}
		case 7:
			if contentSplit[i] != YELLOW+str {
				t.Errorf("Error PrintInColor(YELLOW, %s), got %s", str, contentSplit[i])
			}
		case 8:
			if contentSplit[i] != BBlack+str {
				t.Errorf("Error PrintInColor(BBlack, %s), got %s", str, contentSplit[i])
			}
		case 9:
			if contentSplit[i] != BBlue+str {
				t.Errorf("Error PrintInColor(BBlue, %s), got %s", str, contentSplit[i])
			}
		case 10:
			if contentSplit[i] != BCyan+str {
				t.Errorf("Error PrintInColor(BCyan, %s), got %s", str, contentSplit[i])
			}
		case 11:
			if contentSplit[i] != BGreen+str {
				t.Errorf("Error PrintInColor(BGreen, %s), got %s", str, contentSplit[i])
			}
		case 12:
			if contentSplit[i] != BMagenta+str {
				t.Errorf("Error PrintInColor(BMagenta, %s), got %s", str, contentSplit[i])
			}
		case 13:
			if contentSplit[i] != BRed+str {
				t.Errorf("Error PrintInColor(BRed, %s), got %s", str, contentSplit[i])
			}
		case 14:
			if contentSplit[i] != BWhite+str {
				t.Errorf("Error PrintInColor(BWhite, %s), got %s", str, contentSplit[i])
			}
		case 15:
			if contentSplit[i] != BYellow+str {
				t.Errorf("Error PrintInColor(BYellow, %s), got %s", str, contentSplit[i])
			}
		case 16:
			if contentSplit[i] != UBlack+str {
				t.Errorf("Error PrintInColor(UBlack, %s), got %s", str, contentSplit[i])
			}
		case 17:
			if contentSplit[i] != UBlue+str {
				t.Errorf("Error PrintInColor(UBlue, %s), got %s", str, contentSplit[i])
			}
		case 18:
			if contentSplit[i] != UCyan+str {
				t.Errorf("Error PrintInColor(UCyan, %s), got %s", str, contentSplit[i])
			}
		case 19:
			if contentSplit[i] != UGreen+str {
				t.Errorf("Error PrintInColor(UGreen, %s), got %s", str, contentSplit[i])
			}
		case 20:
			if contentSplit[i] != UMagenta+str {
				t.Errorf("Error PrintInColor(UMagenta, %s), got %s", str, contentSplit[i])
			}
		case 21:
			if contentSplit[i] != URed+str {
				t.Errorf("Error PrintInColor(URed, %s), got %s", str, contentSplit[i])
			}
		case 22:
			if contentSplit[i] != UWhite+str {
				t.Errorf("Error PrintInColor(UWhite, %s), got %s", str, contentSplit[i])
			}
		case 23:
			if contentSplit[i] != UYellow+str {
				t.Errorf("Error PrintInColor(UYellow, %s), got %s", str, contentSplit[i])
			}
		case 24:
			if contentSplit[i] != On_Black+str {
				t.Errorf("Error PrintInColor(On_Black, %s), got %s", str, contentSplit[i])
			}
		case 25:
			if contentSplit[i] != On_Blue+str {
				t.Errorf("Error PrintInColor(On_Blue, %s), got %s", str, contentSplit[i])
			}
		case 26:
			if contentSplit[i] != On_Cyan+str {
				t.Errorf("Error PrintInColor(On_Cyan, %s), got %s", str, contentSplit[i])
			}
		case 27:
			if contentSplit[i] != On_Green+str {
				t.Errorf("Error PrintInColor(On_Green, %s), got %s", str, contentSplit[i])
			}
		case 28:
			if contentSplit[i] != On_Magenta+str {
				t.Errorf("Error PrintInColor(On_Magenta, %s), got %s", str, contentSplit[i])
			}
		case 29:
			if contentSplit[i] != On_Red+str {
				t.Errorf("Error PrintInColor(On_Red, %s), got %s", str, contentSplit[i])
			}
		case 30:
			if contentSplit[i] != On_White+str {
				t.Errorf("Error PrintInColor(On_White, %s), got %s", str, contentSplit[i])
			}
		case 31:
			if contentSplit[i] != On_Yellow+str {
				t.Errorf("Error PrintInColor(On_Yellow, %s), got %s", str, contentSplit[i])
			}
		case 32:
			if contentSplit[i] != IBlack+str {
				t.Errorf("Error PrintInColor(IBlack, %s), got %s", str, contentSplit[i])
			}
		case 33:
			if contentSplit[i] != IBlue+str {
				t.Errorf("Error PrintInColor(IBlue, %s), got %s", str, contentSplit[i])
			}
		case 34:
			if contentSplit[i] != ICyan+str {
				t.Errorf("Error PrintInColor(ICyan, %s), got %s", str, contentSplit[i])
			}
		case 35:
			if contentSplit[i] != IGreen+str {
				t.Errorf("Error PrintInColor(IGreen, %s), got %s", str, contentSplit[i])
			}
		case 36:
			if contentSplit[i] != IMagenta+str {
				t.Errorf("Error PrintInColor(IMagenta, %s), got %s", str, contentSplit[i])
			}
		case 37:
			if contentSplit[i] != IRed+str {
				t.Errorf("Error PrintInColor(IRed, %s), got %s", str, contentSplit[i])
			}
		case 38:
			if contentSplit[i] != IWhite+str {
				t.Errorf("Error PrintInColor(IWhite, %s), got %s", str, contentSplit[i])
			}
		case 39:
			if contentSplit[i] != IYellow+str {
				t.Errorf("Error PrintInColor(IYellow, %s), got %s", str, contentSplit[i])
			}
		case 40:
			if contentSplit[i] != BIBlack+str {
				t.Errorf("Error PrintInColor(BIBlack, %s), got %s", str, contentSplit[i])
			}
		case 41:
			if contentSplit[i] != BIBlue+str {
				t.Errorf("Error PrintInColor(BIBlue, %s), got %s", str, contentSplit[i])
			}
		case 42:
			if contentSplit[i] != BICyan+str {
				t.Errorf("Error PrintInColor(BICyan, %s), got %s", str, contentSplit[i])
			}
		case 43:
			if contentSplit[i] != BIGreen+str {
				t.Errorf("Error PrintInColor(BIGreen, %s), got %s", str, contentSplit[i])
			}
		case 44:
			if contentSplit[i] != BIMagenta+str {
				t.Errorf("Error PrintInColor(BIMagenta, %s), got %s", str, contentSplit[i])
			}
		case 45:
			if contentSplit[i] != BIRed+str {
				t.Errorf("Error PrintInColor(BIRed, %s), got %s", str, contentSplit[i])
			}
		case 46:
			if contentSplit[i] != BIWhite+str {
				t.Errorf("Error PrintInColor(BIWhite, %s), got %s", str, contentSplit[i])
			}
		case 47:
			if contentSplit[i] != BIYellow+str {
				t.Errorf("Error PrintInColor(BIYellow, %s), got %s", str, contentSplit[i])
			}
		case 48:
			if contentSplit[i] != On_IBlack+str {
				t.Errorf("Error PrintInColor(On_IBlack, %s), got %s", str, contentSplit[i])
			}
		case 49:
			if contentSplit[i] != On_IBlue+str {
				t.Errorf("Error PrintInColor(On_IBlue, %s), got %s", str, contentSplit[i])
			}
		case 50:
			if contentSplit[i] != On_ICyan+str {
				t.Errorf("Error PrintInColor(On_ICyan, %s), got %s", str, contentSplit[i])
			}
		case 51:
			if contentSplit[i] != On_IGreen+str {
				t.Errorf("Error PrintInColor(On_IGreen, %s), got %s", str, contentSplit[i])
			}
		case 52:
			if contentSplit[i] != On_IMagenta+str {
				t.Errorf("Error PrintInColor(On_IMagenta, %s), got %s", str, contentSplit[i])
			}
		case 53:
			if contentSplit[i] != On_IRed+str {
				t.Errorf("Error PrintInColor(On_IRed, %s), got %s", str, contentSplit[i])
			}
		case 54:
			if contentSplit[i] != On_IWhite+str {
				t.Errorf("Error PrintInColor(On_IWhite, %s), got %s", str, contentSplit[i])
			}
		case 55:
			if contentSplit[i] != On_IYellow+str {
				t.Errorf("Error PrintInColor(On_IYellow, %s), got %s", str, contentSplit[i])
			}
		case 56:
			if contentSplit[i] != str {
				t.Errorf("Error PrintInColor(ColorOff, %s), got %s", str, contentSplit[i])
			}
		}
	}
}

func TestPrintln(t *testing.T) {
	str := "Test123"
	x := 123
	y := 123.456
	check := true

	file, err := os.OpenFile("unit_test_files/print.txt", os.O_RDWR, 0777)
	if err != nil {
		t.Errorf("Error opening file: %v", err)
	}
	err = file.Truncate(0)
	if err != nil {
		t.Errorf("Error truncating file: %v", err)
	}
	os.Stdout = file

	Println(str)
	Println(x)
	Println(y)
	Println(check)

	err = os.Stdout.Sync()
	if err != nil {
		t.Errorf("Error syncing stdout: %v", err)
	}
	err = os.Stdout.Close()
	if err != nil {
		t.Errorf("Error closing stdout: %v", err)
	}

	content, err := os.ReadFile("unit_test_files/print.txt")
	if err != nil {
		t.Errorf("Error reading file: %v", err)
	}

	contentSplit := strings.Split(string(content), "\n")
	for i := 0; i < len(contentSplit); i++ {
		switch i {
		case 0:
			if contentSplit[i] != str {
				t.Errorf("Error Println(%s), got %s", str, contentSplit[i])
			}
		case 1:
			if contentSplit[i] != fmt.Sprintf("%d", x) {
				t.Errorf("Error Println(%d), got %s", x, contentSplit[i])
			}
		case 2:
			if contentSplit[i] != fmt.Sprintf("%g", y) {
				t.Errorf("Error Println(%g), got %s", y, contentSplit[i])
			}
		case 3:
			if contentSplit[i] != fmt.Sprintf("%t", check) {
				t.Errorf("Error Println(%t), got %s", check, contentSplit[i])
			}
		}
	}
}
