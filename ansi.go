package main

import (
	"fmt"
)

// Based on http://wiki.bash-hackers.org/scripting/terminalcodes
const (
	defaulForeground  = "\x1b[39m"
	defaultBackground = "\x1b[49m"

	// Effects
	reset        = "\x1b[0m"
	bold         = "\x1b[1m"
	faint        = "\x1b[2m"
	standout     = "\x1b[3m"
	underline    = "\x1b[4m"
	blink        = "\x1b[5m"
	reverse      = "\x1b[7m"
	hidden       = "\x1b[8m"
	standoutOff  = "\x1b[23m"
	underlineOff = "\x1b[24m"
	blinkOff     = "\x1b[25m"
	reverseOff   = "\x1b[27m"

	// Colors
	black   = "\x1b[30m"
	red     = "\x1b[31m"
	green   = "\x1b[32m"
	yellow  = "\x1b[33m"
	blue    = "\x1b[34m"
	magenta = "\x1b[35m"
	cyan    = "\x1b[36m"
	white   = "\x1b[37m"

	// Bright colors
	bBlack   = "\x1b[90m"
	bRed     = "\x1b[91m"
	bGreen   = "\x1b[92m"
	bYellow  = "\x1b[93m"
	bBlue    = "\x1b[94m"
	bMagenta = "\x1b[95m"
	bCyan    = "\x1b[96m"
	bWhite   = "\x1b[97m"

	// Backgrounds
	blackBg   = "\x1b[40m"
	redBg     = "\x1b[41m"
	greenBg   = "\x1b[42m"
	yellowBg  = "\x1b[43m"
	blueBg    = "\x1b[44m"
	magentaBg = "\x1b[45m"
	cyanBg    = "\x1b[46m"
	whiteBg   = "\x1b[47m"

	// Light backgrounds
	bBlackBg   = "\x1b[100m"
	bRedBg     = "\x1b[101m"
	bGreenBg   = "\x1b[102m"
	bYellowBg  = "\x1b[103m"
	bBlueBg    = "\x1b[104m"
	bMagentaBg = "\x1b[105m"
	bCyanBg    = "\x1b[106m"
	bWhiteBg   = "\x1b[107m"
)

func apply(str string, code string) string {
	return fmt.Sprintf("%s%s%s", code, str, reset)
}

// Bold - Prints the text in bold.
func Bold(str string) string {
	return apply(str, bold)
}

// Red - Colors the output in red.
func Red(str string) string {
	return apply(str, bRed)
}

// DarkGrey - Colors the output in dark grey.
func DarkGrey(str string) string {
	return apply(str, bBlack)
}
