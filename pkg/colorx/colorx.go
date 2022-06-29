package colorx

import "fmt"

const (
	GreenRaw   = "\033[97;42m"
	WhiteRaw   = "\033[90;47m"
	YellowRaw  = "\033[90;43m"
	RedRaw     = "\033[97;41m"
	BlueRaw    = "\033[97;44m"
	MagentaRaw = "\033[97;45m"
	CyanRaw    = "\033[97;46m"
	ResetRaw   = "\033[0m"

	ColorBold     = 1
	ColorDarkGray = 90
)

// Colorize returns the string s wrapped in ANSI code c, unless disabled is true.
func Colorize(s interface{}, c int) string {
	//if disabled {
	//	return fmt.Sprintf("%s", s)
	//}
	return fmt.Sprintf("\x1b[%dm%v\x1b[0m", c, s)
}

func Green(v any) string {
	return fmt.Sprintf("%s%v%s", GreenRaw, v, ResetRaw)
}

func White(v any) string {
	return fmt.Sprintf("%s%v%s", WhiteRaw, v, ResetRaw)
}
func Yellow(v any) string {
	return fmt.Sprintf("%s%v%s", YellowRaw, v, ResetRaw)
}

func Red(v any) string {
	return fmt.Sprintf("%s%v%s", RedRaw, v, ResetRaw)
}

func Blue(v any) string {
	return fmt.Sprintf("%s%v%s", BlueRaw, v, ResetRaw)
}

func Magenta(v any) string {
	return fmt.Sprintf("%s%v%s", MagentaRaw, v, ResetRaw)
}

func Cyan(v any) string {
	return fmt.Sprintf("%s%v%s", CyanRaw, v, ResetRaw)
}

//
//  printf
//

func GreenPf(v any) {
	fmt.Printf("%s%v%s", GreenRaw, v, ResetRaw)
}

func WhitePf(v any) {
	fmt.Printf("%s%v%s", WhiteRaw, v, ResetRaw)
}
func YellowPf(v any) {
	fmt.Printf("%s%v%s", YellowRaw, v, ResetRaw)
}

func RedPf(v any) {
	fmt.Printf("%s%v%s", RedRaw, v, ResetRaw)
}

func BluePf(v any) {
	fmt.Printf("%s%v%s", BlueRaw, v, ResetRaw)
}

func MagentaPf(v any) {
	fmt.Printf("%s%v%s", MagentaRaw, v, ResetRaw)
}

func CyanPf(v any) {
	fmt.Printf("%s%v%s", CyanRaw, v, ResetRaw)
}

//
// println
//

func GreenPln(v any) {
	fmt.Printf("%s%v%s\n", GreenRaw, v, ResetRaw)
}

func WhitePln(v any) {
	fmt.Printf("%s%v%s\n", WhiteRaw, v, ResetRaw)
}
func YellowPln(v any) {
	fmt.Printf("%s%v%s\n", YellowRaw, v, ResetRaw)
}

func RedPln(v any) {
	fmt.Printf("%s%v%s\n", RedRaw, v, ResetRaw)
}

func BluePln(v any) {
	fmt.Printf("%s%v%s\n", BlueRaw, v, ResetRaw)
}

func MagentaPln(v any) {
	fmt.Printf("%s%v%s\n", MagentaRaw, v, ResetRaw)
}

func CyanPln(v any) {
	fmt.Printf("%s%v%s\n", CyanRaw, v, ResetRaw)
}
