package main

// #cgo pkg-config: gdk-3.0 gio-2.0 glib-2.0 gobject-2.0 gtk+-3.0
// #include <liveUpdate.h>
import "C"

import (
	"os/exec"
	"time"
	"unsafe"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

func liveUpdate() {
	setAppSym("", true, false)
	exec.Command("cp", "./index.theme", tmpBuild+"/index.theme").Run()
	time.Sleep(time.Second * 5)
	// IconTheme is a representation of GTK's GtkIconTheme
	type IconTheme struct {
		Theme *C.GtkIconTheme
	}

	// IconThemeGetForScreen is a wrapper around gtk_icon_theme_get_for_screen().
	getIconTheme := func() (*IconTheme, error) {
		screen, _ := gdk.ScreenGetDefault()
		scr, _ := gtk.IconThemeGetForScreen(*screen)
		c := C.gtk_icon_theme_set_custom_theme(unsafe.Pointer(scr)), (*C.gchar)(C.CString(""))
		if c == nil {
			return nil, nil
		}
		return &IconTheme{c}, nil
	}
	gtk.Init(nil)
	getIconTheme()
}
