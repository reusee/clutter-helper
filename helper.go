package clutter_helper

/*
#include <clutter/clutter.h>
#include <clutter-gtk/clutter-gtk.h>
#include <gtk/gtk.h>
#cgo pkg-config: clutter-1.0 clutter-gtk-1.0 gtk+-3.0
*/
import "C"
import (
	"fmt"
	"log"
	"unsafe"
)

var p = fmt.Printf

func Init() {
	var argc C.int
	C.gtk_clutter_init(&argc, nil)
}

func Main() {
	C.gtk_main()
}

func Quit() {
	C.gtk_main_quit()
}

func NewColorFromString(s string) *C.ClutterColor {
	var color C.ClutterColor
	if C.clutter_color_from_string(&color, toGStr(s)) != C.TRUE {
		log.Fatalf("wrong color format %s", s)
	}
	return &color
}

var _gstrs = make(map[string]*C.gchar)

func toGStr(s string) *C.gchar {
	if gstr, ok := _gstrs[s]; ok {
		return gstr
	}
	gstr := (*C.gchar)(unsafe.Pointer(C.CString(s)))
	_gstrs[s] = gstr
	return gstr
}
