//go:build linuxdesktop

package desktop

// #cgo pkg-config: gtk+-3.0
// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"

import (
	webview "github.com/webview/webview_go"
)

const GTK_WIN_POS_CENTER = 1
const GTK_WIN_POS_CENTER_ALWAYS = 3

func SetWindowPosSize(browser *webview.WebView) {
	(*browser).SetSize(window_width, window_height, webview.HintNone)
	C.gtk_window_set_position(C.toGtkWindow((*browser).Window()), C.GtkWindowPosition(GTK_WIN_POS_CENTER_ALWAYS))
	// C.gtk_window_move(C.toGtkWindow((*browser).Window()), 0, 0)
}
