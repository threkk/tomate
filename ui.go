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
	free := (float64(total-current) / float64(total))
	percentage := (1 - free) * 100

	length := int(float64(ui.Size) * free)
	filled := strings.Repeat(Red(high), ui.Size-length)
	toFill := strings.Repeat(DarkGrey(low), length)

	d := Bold(time.Duration(nanoTosec(total - current)).String())

	bar := fmt.Sprintf("  %2.0f%% | %s%s | %s", percentage, filled, toFill, d)

	return bar
}
