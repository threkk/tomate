package main

import (
	"fmt"
	"strings"
	"time"
)

// http://www.theasciicode.com.ar/extended-ascii-code/black-square-ascii-code-254.html
const (
	low    = "░"
	medium = "▒"
	high   = "▓"
	full   = "█"
)

func nanoTosec(nano int64) int64 {
	return nano * 1000000000
}

// Bar - Interface which represents any type of interface for the pomodoro.
type Bar interface {
	Frame(current int64, total int64) string
}

// RedBlackBar - Bar with a red filling and dark grey background.
type RedBlackBar struct {
	Size int
}

// Frame - Displays a frame.
func (ui *RedBlackBar) Frame(current int64, total int64) string {
	d := Bold(time.Duration(nanoTosec(total - current)).String())

	size := ui.Size - (12 + len(d))
	free := (float64(total-current) / float64(total))
	percentage := (1 - free) * 100

	length := int(float64(size) * free)
	filled := strings.Repeat(Red(high), size-length)
	toFill := strings.Repeat(DarkGrey(low), length)

	bar := fmt.Sprintf("  %3.0f%% | %s%s | %s", percentage, filled, toFill, d)

	return bar
}

// MinimalMagentaBar - Magenta bar with nothing else.
type MinimalMagentaBar struct {
	Size int
}

// Frame - Display a frame.
func (ui *MinimalMagentaBar) Frame(current int64, total int64) string {
	free := (float64(total-current) / float64(total))

	length := int(float64(ui.Size) * free)
	filled := strings.Repeat(Magenta(full), ui.Size-length)
	toFill := strings.Repeat(LightGrey(low), length)

	bar := fmt.Sprintf("%s%s", filled, toFill)

	return bar
}

// HomeBrewBar - Classics never die.
type HomeBrewBar struct {
	Size int
}

// Frame - Displays a frame.
func (ui *HomeBrewBar) Frame(current int64, total int64) string {
	d := Bold(time.Duration(nanoTosec(total - current)).String())

	size := ui.Size - (12 + len(d))
	free := (float64(total-current) / float64(total))
	percentage := (1 - free) * 100

	length := int(float64(size) * free)
	filled := strings.Repeat(Green(full), size-length-1) + Blink(LightGreen(full))
	toFill := strings.Repeat(DarkGrey(low), length)

	bar := fmt.Sprintf("  %3.0f%% | %s%s | %s", percentage, filled, toFill, d)

	return bar
}
