//go:build windowsdesktop

package desktop

import (
	webview "github.com/webview/webview_go"
)

var window_width int = 400
var window_height int = 600

func Run(url string) {
	browser := webview.New(false)
	defer browser.Destroy()
	SetWindowPosSize(&browser)
	browser.SetTitle("Simple deploy manager")
	browser.Navigate(url)
	browser.Run()
}
