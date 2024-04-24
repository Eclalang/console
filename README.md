# Console library

[![Go Report Card](https://goreportcard.com/badge/github.com/Eclalang/console)](https://goreportcard.com/report/github.com/Eclalang/console)
[![codecov](https://codecov.io/gh/Eclalang/console/graph/badge.svg?token=YNCIYERVBO)](https://codecov.io/gh/Eclalang/console)

## Candidate functions :

|  Func name   |                   Prototype                   |                       Description                        | Comments |
|:------------:|:---------------------------------------------:|:--------------------------------------------------------:|:--------:|
|    Clear     |                    Clear()                    |                    Clears the console                    |   N/A    |
|    Input     |               Input() string {}               |    Takes user input from console and returns a string    |   N/A    |
|  InputFloat  |        InputFloat() (float, error) {}         |   Takes user input from console and returns a float64    |   N/A    |
|   InputInt   |          InputInt() (int, error) {}           |     Takes user input from console and returns an int     |   N/A    |
|    Print     |             Print(args ...any) {}             |                  Prints args to console                  |   N/A    |
|    Printf    |     Printf(format string, args ...any) {}     |       Prints args to console formatted with format       |   N/A    |
| PrintInColor |  PrintInColor(color string, args ...any) {}   |             Prints args to console in color              |   N/A    |
|   Println    |            Println(args ...any) {}            |          Prints args to console with a newline           |   N/A    |
|    SPrint    |         SPrint(args ...any) string {}         | Returns a string with the args using the default formats |   N/A    |
|   SPrintf    | SPrintf(format string, args ...any) string {} |   Returns a string with the args formatted with format   |   N/A    |
|   SPrintln   |        SPrintln(args ...any) string {}        | Returns a string with the args and a newline at the end  |   N/A    |

## Colors constants :

### Regular colors

| Color Name |   Detail    |
|:----------:|:-----------:|
|  COLOROFF  | Reset color |
|   BLACK    |    Black    |
|    BLUE    |    Blue     |
|    CYAN    |    Cyan     |
|   GREEN    |    Green    |
|  MAGENTA   |   Magenta   |
|    RED     |     Red     |
|   WHITE    |    White    |
|   YELLOW   |   Yellow    |

### Bold colors

| Color Name | Detail  |
|:----------:|:-------:|
|  B_BLACK   |  Black  |
|   B_BLUE   |  Blue   |
|   B_CYAN   |  Cyan   |
|  B_GREEN   |  Green  |
| B_MAGENTA  | Magenta |
|   B_RED    |   Red   |
|  B_WHITE   |  White  |
|  B_YELLOW  | Yellow  |

### Underline colors

| Color Name | Detail  |
|:----------:|:-------:|
|  U_BLACK   |  Black  |
|   U_BLUE   |  Blue   |
|   U_CYAN   |  Cyan   |
|  U_GREEN   |  Green  |
| U_MAGENTA  | Magenta |
|   U_RED    |   Red   |
|  U_WHITE   |  White  |
|  U_YELLOW  | Yellow  |

### Background colors

| Color Name | Detail  |
|:----------:|:-------:|
|  ON_BLACK  |  Black  |
|  ON_BLUE   |  Blue   |
|  ON_CYAN   |  Cyan   |
|  ON_GREEN  |  Green  |
| ON_MAGENTA | Magenta |
|   ON_RED   |   Red   |
|  ON_WHITE  |  White  |
| ON_YELLOW  | Yellow  |

### High intensity colors

| Color Name | Detail  |
|:----------:|:-------:|
|  I_BLACK   |  Black  |
|   I_BLUE   |  Blue   |
|   I_CYAN   |  Cyan   |
|  I_GREEN   |  Green  |
| I_MAGENTA  | Magenta |
|   I_RED    |   Red   |
|  I_WHITE   |  White  |
|  I_YELLOW  | Yellow  |

### Bold high intensity colors

| Color Name | Detail  |
|:----------:|:-------:|
|  BI_BLACK  |  Black  |
|  BI_BLUE   |  Blue   |
|  BI_CYAN   |  Cyan   |
|  BI_GREEN  |  Green  |
| BI_MAGENTA | Magenta |
|   BI_RED   |   Red   |
|  BI_WHITE  |  White  |
| BI_YELLOW  | Yellow  |

### High intensity background colors

| Color Name  | Detail  |
|:-----------:|:-------:|
|  ONI_BLACK  |  Black  |
|  ONI_BLUE   |  Blue   |
|  ONI_CYAN   |  Cyan   |
|  ONI_GREEN  |  Green  |
| ONI_MAGENTA | Magenta |
|   ONI_RED   |   Red   |
|  ONI_WHITE  |  White  |
| ONI_YELLOW  | Yellow  |
