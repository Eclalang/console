# Console library

## Candidate functions :

|  Func name   |                     Prototype                      |                     Description                     | Comments |
|:------------:|:--------------------------------------------------:|:---------------------------------------------------:|:--------:|
|    Clear     |                      Clear()                       |                 Clears the console                  |   N/A    |
|    Input     |                 Input() string {}                  | Takes user input from console and returns a string  |   N/A    |
|  InputFloat  |           InputFloat() (float, error) {}           | Takes user input from console and returns a float64 |   N/A    |
|   InputInt   |             InputInt() (int, error) {}             |  Takes user input from console and returns an int   |   N/A    |
|    Print     |           Print(args ...interface{}) {}            |               Prints args to console                |   N/A    |
|    Printf    |   Printf(format string, args ...interface{}) {}    |    Prints args to console formatted with format     |   N/A    |
| PrintInColor | PrintInColor(color string, args ...interface{}) {} |           Prints args to console in color           |   N/A    |
|   Println    |          Println(args ...interface{}) {}           |        Prints args to console with a newline        |   N/A    |

## Colors constants :

### Regular colors

| Color Name |   Detail    |
|:----------:|:-----------:|
|  ColorOff  | Reset color |
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
|   BBlack   |  Black  |
|   BBlue    |  Blue   |
|   BCyan    |  Cyan   |
|   BGreen   |  Green  |
|  BMagenta  | Magenta |
|    BRed    |   Red   |
|   BWhite   |  White  |
|  BYellow   | Yellow  |

### Underline colors

| Color Name | Detail  |
|:----------:|:-------:|
|   UBlack   |  Black  |
|   UBlue    |  Blue   |
|   UCyan    |  Cyan   |
|   UGreen   |  Green  |
|  UMagenta  | Magenta |
|    URed    |   Red   |
|   UWhite   |  White  |
|  UYellow   | Yellow  |

### Background colors

| Color Name | Detail  |
|:----------:|:-------:|
|  On_Black  |  Black  |
|  On_Blue   |  Blue   |
|  On_Cyan   |  Cyan   |
|  On_Green  |  Green  |
| On_Magenta | Magenta |
|   On_Red   |   Red   |
|  On_White  |  White  |
| On_Yellow  | Yellow  |

### High intensity colors

| Color Name | Detail  |
|:----------:|:-------:|
|   IBlack   |  Black  |
|   IBlue    |  Blue   |
|   ICyan    |  Cyan   |
|   IGreen   |  Green  |
|  IMagenta  | Magenta |
|    IRed    |   Red   |
|   IWhite   |  White  |
|  IYellow   | Yellow  |

### Bold high intensity colors

| Color Name | Detail  |
|:----------:|:-------:|
|  BIBlack   |  Black  |
|   BIBlue   |  Blue   |
|   BICyan   |  Cyan   |
|  BIGreen   |  Green  |
| BIMagenta  | Magenta |
|   BIRed    |   Red   |
|  BIWhite   |  White  |
|  BIYellow  | Yellow  |

### High intensity background colors

| Color Name  | Detail  |
|:-----------:|:-------:|
|  On_IBlack  |  Black  |
|  On_IBlue   |  Blue   |
|  On_ICyan   |  Cyan   |
|  On_IGreen  |  Green  |
| On_IMagenta | Magenta |
|   On_IRed   |   Red   |
|  On_IWhite  |  White  |
| On_IYellow  | Yellow  |