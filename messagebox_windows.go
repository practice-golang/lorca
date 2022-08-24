//go:build windows
// +build windows

package lorca

import (
	"syscall"
	"unsafe"
)

func messageBox(title, text string) bool {
	user32 := syscall.NewLazyDLL("user32.dll")
	messageBoxW := user32.NewProc("MessageBoxW")
	mbYesNo := 0x00000004
	mbIconQuestion := 0x00000020
	idYes := 6
	textPTR, _ := syscall.UTF16PtrFromString(text)
	titlePTR, _ := syscall.UTF16PtrFromString(title)
	ret, _, _ := messageBoxW.Call(0, uintptr(unsafe.Pointer(textPTR)),
		uintptr(unsafe.Pointer(titlePTR)), uintptr(uint(mbYesNo|mbIconQuestion)))
	return int(ret) == idYes
}
