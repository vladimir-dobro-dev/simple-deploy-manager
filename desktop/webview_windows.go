//go:build windowsdesktop

package desktop

import (
	"syscall"

	webview "github.com/webview/webview_go"
)

func SetWindowPosSize(browser *webview.WebView) {
	monitor_width, monitor_height := getMonitorSize()
	x := (monitor_width - window_width) / 2
	y := (monitor_height - window_height) / 2
	window := (*browser).Window()
	setWindowPos((uintptr)(window), 0, x, y, window_width, window_height, SWP_SHOWWINDOW)
}

func getMonitorSize() (int, int) {
	libuser32 := loadLibrary("user32.dll")
	getSystemMetrics := getProcAddress(libuser32, "GetSystemMetrics")
	width, _, _ := syscall.SyscallN(getSystemMetrics, SM_CXSCREEN)
	height, _, _ := syscall.SyscallN(getSystemMetrics, SM_CYSCREEN)
	return int(width), int(height)
}

func setWindowPos(hWnd uintptr, hWndInsertAfter uintptr, x, y, width, height int, flags uint32) bool {
	libuser32 := loadLibrary("user32.dll")
	setWindowPos := getProcAddress(libuser32, "SetWindowPos")
	ret, _, _ := syscall.SyscallN(setWindowPos,
		uintptr(hWnd),
		uintptr(hWndInsertAfter),
		uintptr(x),
		uintptr(y),
		uintptr(width),
		uintptr(height),
		uintptr(flags),
	)

	return ret != 0
}

func loadLibrary(name string) uintptr {
	lib, err := syscall.LoadLibrary(name)
	if err != nil {
		panic(err)
	}

	return uintptr(lib)
}

func getProcAddress(lib uintptr, name string) uintptr {
	addr, err := syscall.GetProcAddress(syscall.Handle(lib), name)
	if err != nil {
		panic(err)
	}

	return uintptr(addr)
}

// SetWindowPos flags
const (
	SM_CXSCREEN = 0
	SM_CYSCREEN = 1
)

// SetWindowPos flags
const (
	SWP_ASYNCWINDOWPOS = 0x4000
	SWP_DEFERERASE     = 0x2000
	SWP_DRAWFRAME      = 0x0020
	SWP_FRAMECHANGED   = 0x0020
	SWP_HIDEWINDOW     = 0x0080
	SWP_NOACTIVATE     = 0x0010
	SWP_NOCOPYBITS     = 0x0100
	SWP_NOMOVE         = 0x0002
	SWP_NOOWNERZORDER  = 0x0200
	SWP_NOREDRAW       = 0x0008
	SWP_NOREPOSITION   = SWP_NOOWNERZORDER
	SWP_NOSENDCHANGING = 0x0400
	SWP_NOSIZE         = 0x0001
	SWP_NOZORDER       = 0x0004
	SWP_SHOWWINDOW     = 0x0040
)
