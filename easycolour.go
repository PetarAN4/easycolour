package easycolour

import (
	"fmt"
)

const (
	escape  string = "\033["
	normal  string = "0"
	bold    string = "1"
	black   string = ";90m"
	red     string = ";91m"
	green   string = ";92m"
	yellow  string = ";93m"
	blue    string = ";94m"
	magenta string = ";95m"
	cyan    string = ";96m"
	white   string = ";97m"
	end     string = "0m"
)

type format string

const (
	//Normal default mode
	Normal format = "0"
	//Bold bold mode
	Bold format = "1"
	//Underline underline mode
	Underline format = "4"
	//Negative negative mode
	Negative format = "7"
)

//Black returns s in black colour
func Black(format format, s string) string {
	return fmt.Sprintf("%s%s%s%s%s%s", escape, format, black, s, escape, end)
}

//Red returns s in red colour
func Red(format format, s string) string {
	return fmt.Sprintf("%s%s%s%s%s%s", escape, format, red, s, escape, end)
}

//Green returns s in green colour
func Green(format format, s string) string {
	return fmt.Sprintf("%s%s%s%s%s%s", escape, format, green, s, escape, end)
}

//Yellow returns s in yellow colour
func Yellow(format format, s string) string {
	return fmt.Sprintf("%s%s%s%s%s%s", escape, format, yellow, s, escape, end)
}

//Blue returns s in blue colour
func Blue(format format, s string) string {
	return fmt.Sprintf("%s%s%s%s%s%s", escape, format, blue, s, escape, end)
}

//Magenta returns s in magenta colour
func Magenta(format format, s string) string {
	return fmt.Sprintf("%s%s%s%s%s%s", escape, format, magenta, s, escape, end)
}

//Cyan returns s in cyan colour
func Cyan(format format, s string) string {
	return fmt.Sprintf("%s%s%s%s%s%s", escape, format, cyan, s, escape, end)
}

//White returns s in white colour
func White(format format, s string) string {
	return fmt.Sprintf("%s%s%s%s%s%s", escape, format, white, s, escape, end)
}
