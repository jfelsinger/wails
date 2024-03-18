//go:build windows

package w32

import (
	"syscall"
)

var (
	wtsapi32 = syscall.NewLazyDLL("wtsapi32.dll")

	procWTSQuerySessionInformation       = wtsapi32.NewProc("WTSQuerySessionInformation")
	procWTSFreeMemoryEx                  = wtsapi32.NewProc("WTSFreeMemoryEx")
	procWTSRegisterSessionNotification   = wtsapi32.NewProc("WTSRegisterSessionNotification")
	procWTSUnRegisterSessionNotification = wtsapi32.NewProc("WTSUnRegisterSessionNotification")
)

func WTSRegisterSessionNotification(hwnd HWND) bool {
	ret, _, _ := procWTSRegisterSessionNotification.Call(
		uintptr(hwnd),
		uintptr(NOTIFY_FOR_THIS_SESSION))
	return ret != 0
}

func WTSUnRegisterSessionNotification(hwnd HWND) bool {
	ret, _, _ := procWTSUnRegisterSessionNotification.Call(
		uintptr(hwnd))
	return ret != 0
}
