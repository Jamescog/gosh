package utils

import (
	"fmt"
	"runtime"
)

var CAN_PRINT_COLOR = false

func InitColor() {
	if runtime.GOOS == "windows" {
		CAN_PRINT_COLOR = enableVTProcessing()
	} else {
		CAN_PRINT_COLOR = true
	}
}

func GetColored(text, colorName string) string {
	if !CAN_PRINT_COLOR {
		return text
	}

	colorMap := map[string]string{
		"red":     "\033[31m",
		"green":   "\033[32m",
		"yellow":  "\033[33m",
		"blue":    "\033[34m",
		"magenta": "\033[35m",
		"cyan":    "\033[36m",
		"gray":    "\033[37m",
		"white":   "\033[97m",
	}

	code, ok := colorMap[colorName]
	if !ok {
		return text
	}

	return fmt.Sprintf("%s%s\033[0m", code, text)
}

func Red(text string) string     { return GetColored(text, "red") }
func Green(text string) string   { return GetColored(text, "green") }
func Yellow(text string) string  { return GetColored(text, "yellow") }
func Blue(text string) string    { return GetColored(text, "blue") }
func Magenta(text string) string { return GetColored(text, "magenta") }
func Cyan(text string) string    { return GetColored(text, "cyan") }
func Gray(text string) string    { return GetColored(text, "gray") }
func White(text string) string   { return GetColored(text, "white") }
