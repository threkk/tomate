// +build windows

package terminal

import (
	"syscall"
	"unsafe"
)

// From https://stackoverflow.com/questions/16569433/get-terminal-size-in-go/16576712#16576712
type winsize struct {
	Row    uint16
	Col    uint16
	Xpixel uint16
	Ypixel uint16
}

// GetColumns - Return the size of the terminal at the moment it is called.
// Defaults to 80 if error.
func GetColumns() uint {
	ws := &winsize{}
	retCode, _, errno := syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(syscall.Stdin),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(ws)))

	if int(retCode) == -1 {
		return 80
	}
	return uint(ws.Col)
}
