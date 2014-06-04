package clutter_helper

/*
#include <clutter/clutter.h>
#cgo pkg-config: clutter-1.0
*/
import "C"
import (
	"fmt"
	"log"
	"reflect"
	"unsafe"
)

var p = fmt.Printf

func Init() {
	var argc C.int
	C.clutter_init(&argc, nil)
}

func Main() {
	C.clutter_main()
}

func Quit() {
	C.clutter_main_quit()
}

type Actor struct {
	C *C.ClutterActor
}

func NewStage() Actor {
	return Actor{
		C: C.clutter_stage_new(),
	}
}

func NewActorFromC(p interface{}) Actor {
	return Actor{
		C: (*C.ClutterActor)(unsafe.Pointer(reflect.ValueOf(p).Pointer())),
	}
}

func (actor Actor) Show() {
	C.clutter_actor_show(actor.C)
}

func (actor Actor) AddChild(c interface{}) {
	switch child := c.(type) {
	case *C.ClutterActor:
		C.clutter_actor_add_child(actor.C, child)
	case Actor:
		C.clutter_actor_add_child(actor.C, child.C)
	}
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
