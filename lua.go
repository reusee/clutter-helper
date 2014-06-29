package clutter_helper

/*
#include <clutter/clutter.h>

gboolean is_actor(void *o) {
	return CLUTTER_IS_ACTOR(o);
}

gboolean is_action(void *o) {
	return CLUTTER_IS_ACTION(o);
}

gboolean is_constraint(void *o) {
	return CLUTTER_IS_CONSTRAINT(o);
}

gboolean is_effect(void *o) {
	return CLUTTER_IS_EFFECT(o);
}

*/
import "C"
import (
	"log"
	"unsafe"

	"github.com/reusee/lgo"
)

func FromLua(code string) map[string]unsafe.Pointer {
	lua := lgo.NewLua()
	bindings := make(map[string]unsafe.Pointer)

	processActorArgs := func(actor *C.ClutterActor, args map[interface{}]interface{}) unsafe.Pointer {
		pointer := unsafe.Pointer(actor)
		var minFilter, magFilter *C.ClutterScalingFilter
		var scaleX C.gdouble = 1.0
		var scaleY C.gdouble = 1.0
		var translationX, translationY, translationZ C.gfloat
		for k, v := range args {
			switch key := k.(type) {
			case string:
				switch key {
				case "id":
					bindings[v.(string)] = pointer
				case "bgcolor", "background_color":
					C.clutter_actor_set_background_color(actor, NewColorFromString(v.(string)))
				case "clip_rect":
					clip := (*C.ClutterRect)(v.(unsafe.Pointer))
					C.clutter_actor_set_clip(actor,
						C.gfloat(clip.origin.x),
						C.gfloat(clip.origin.y),
						C.gfloat(clip.size.width),
						C.gfloat(clip.size.height))
				case "content":
					C.clutter_actor_set_content(actor, (*C.ClutterContent)(v.(unsafe.Pointer)))
				case "content_gravity":
					switch v.(string) {
					case "top-left":
						C.clutter_actor_set_content_gravity(actor, C.CLUTTER_CONTENT_GRAVITY_TOP_LEFT)
					case "top":
						C.clutter_actor_set_content_gravity(actor, C.CLUTTER_CONTENT_GRAVITY_TOP)
					case "top-right":
						C.clutter_actor_set_content_gravity(actor, C.CLUTTER_CONTENT_GRAVITY_TOP_RIGHT)
					case "left":
						C.clutter_actor_set_content_gravity(actor, C.CLUTTER_CONTENT_GRAVITY_LEFT)
					case "center":
						C.clutter_actor_set_content_gravity(actor, C.CLUTTER_CONTENT_GRAVITY_CENTER)
					case "right":
						C.clutter_actor_set_content_gravity(actor, C.CLUTTER_CONTENT_GRAVITY_RIGHT)
					case "bottom-left":
						C.clutter_actor_set_content_gravity(actor, C.CLUTTER_CONTENT_GRAVITY_BOTTOM_LEFT)
					case "bottom":
						C.clutter_actor_set_content_gravity(actor, C.CLUTTER_CONTENT_GRAVITY_BOTTOM)
					case "bottom-right":
						C.clutter_actor_set_content_gravity(actor, C.CLUTTER_CONTENT_GRAVITY_BOTTOM_RIGHT)
					case "resize-fill":
						C.clutter_actor_set_content_gravity(actor, C.CLUTTER_CONTENT_GRAVITY_RESIZE_FILL)
					case "resize-aspect":
						C.clutter_actor_set_content_gravity(actor, C.CLUTTER_CONTENT_GRAVITY_RESIZE_ASPECT)
					default:
						log.Fatalf("unknown content gravity: %s", v.(string))
					}
				case "content_repeat":
					switch v.(string) {
					case "none":
						C.clutter_actor_set_content_repeat(actor, C.CLUTTER_REPEAT_NONE)
					case "x", "x-axis":
						C.clutter_actor_set_content_repeat(actor, C.CLUTTER_REPEAT_X_AXIS)
					case "y", "y-axis":
						C.clutter_actor_set_content_repeat(actor, C.CLUTTER_REPEAT_Y_AXIS)
					case "both":
						C.clutter_actor_set_content_repeat(actor, C.CLUTTER_REPEAT_BOTH)
					default:
						log.Fatalf("unknown content repeat: %s", v.(string))
					}
				case "x", "fixed_x":
					C.clutter_actor_set_x(actor, C.gfloat(v.(float64)))
				case "y", "fixed_y":
					C.clutter_actor_set_y(actor, C.gfloat(v.(float64)))
				case "height":
					C.clutter_actor_set_height(actor, C.gfloat(v.(float64)))
				case "layout", "layout_manager":
					C.clutter_actor_set_layout_manager(actor, (*C.ClutterLayoutManager)(v.(unsafe.Pointer)))
				case "mag_filter", "magnification_filter":
					magFilter = (*C.ClutterScalingFilter)(v.(unsafe.Pointer))
					C.clutter_actor_set_content_scaling_filters(actor, *minFilter, *magFilter)
				case "margin_bottom":
					C.clutter_actor_set_margin_bottom(actor, C.gfloat(v.(float64)))
				case "margin_left":
					C.clutter_actor_set_margin_left(actor, C.gfloat(v.(float64)))
				case "margin_right":
					C.clutter_actor_set_margin_right(actor, C.gfloat(v.(float64)))
				case "margin_top":
					C.clutter_actor_set_margin_top(actor, C.gfloat(v.(float64)))
				//TODO min height
				//TODO min width
				case "min_filter", "minification_filter":
					minFilter = (*C.ClutterScalingFilter)(v.(unsafe.Pointer))
					C.clutter_actor_set_content_scaling_filters(actor, *minFilter, *magFilter)
				case "name":
					C.clutter_actor_set_name(actor, toGStr(v.(string)))
				//TODO natural height
				//TODO natural width
				case "offscreen_redirect":
					switch v.(string) {
					case "auto", "automatic-for-opacity":
						C.clutter_actor_set_offscreen_redirect(actor, C.CLUTTER_OFFSCREEN_REDIRECT_AUTOMATIC_FOR_OPACITY)
					case "always":
						C.clutter_actor_set_offscreen_redirect(actor, C.CLUTTER_OFFSCREEN_REDIRECT_ALWAYS)
					default:
						log.Fatalf("unknown offscreen redirect option: %s", v.(string))
					}
				case "opacity":
					C.clutter_actor_set_opacity(actor, C.guint8(uint8(v.(float64))))
				case "pivot_point":
					point := (*C.ClutterPoint)(v.(unsafe.Pointer))
					C.clutter_actor_set_pivot_point(actor, C.gfloat(point.x), C.gfloat(point.y))
				case "pivot_point_z":
					C.clutter_actor_set_pivot_point_z(actor, C.gfloat(v.(float64)))
				case "position":
					point := (*C.ClutterPoint)(v.(unsafe.Pointer))
					C.clutter_actor_set_position(actor, C.gfloat(point.x), C.gfloat(point.y))
				case "reactive":
					b := C.FALSE
					if v.(bool) {
						b = C.TRUE
					}
					C.clutter_actor_set_reactive(actor, C.gboolean(b))
				case "request_mode":
					switch v.(string) {
					case "height", "height_for_width":
						C.clutter_actor_set_request_mode(actor, C.CLUTTER_REQUEST_HEIGHT_FOR_WIDTH)
					case "width", "width_for_height":
						C.clutter_actor_set_request_mode(actor, C.CLUTTER_REQUEST_WIDTH_FOR_HEIGHT)
					default:
						log.Fatalf("unknown request mode: %s", v.(string))
					}
				case "rotation_angle_x":
					C.clutter_actor_set_rotation_angle(actor, C.CLUTTER_X_AXIS, C.gdouble(v.(float64)))
				case "rotation_angle_y":
					C.clutter_actor_set_rotation_angle(actor, C.CLUTTER_Y_AXIS, C.gdouble(v.(float64)))
				case "rotation_angle_z":
					C.clutter_actor_set_rotation_angle(actor, C.CLUTTER_Z_AXIS, C.gdouble(v.(float64)))
				case "scale_x":
					scaleX = C.gdouble(v.(float64))
					C.clutter_actor_set_scale(actor, scaleX, scaleY)
				case "scale_y":
					scaleY = C.gdouble(v.(float64))
					C.clutter_actor_set_scale(actor, scaleX, scaleY)
				case "scale_z":
					C.clutter_actor_set_scale_z(actor, C.gdouble(v.(float64)))
				case "size":
					size := (*C.ClutterSize)(v.(unsafe.Pointer))
					C.clutter_actor_set_size(actor, C.gfloat(size.width), C.gfloat(size.height))
				case "text_direction", "text_dir":
					switch v.(string) {
					case "default":
						C.clutter_actor_set_text_direction(actor, C.CLUTTER_TEXT_DIRECTION_DEFAULT)
					case "left", "left-to-right", "ltr":
						C.clutter_actor_set_text_direction(actor, C.CLUTTER_TEXT_DIRECTION_LTR)
					case "right", "right-to-left", "rtl":
						C.clutter_actor_set_text_direction(actor, C.CLUTTER_TEXT_DIRECTION_RTL)
					default:
						log.Fatalf("unknown text direction: %s", v.(string))
					}
				//TODO transform
				case "translation_x":
					translationX = C.gfloat(v.(float64))
					C.clutter_actor_set_translation(actor, translationX, translationY, translationZ)
				case "translation_y":
					translationY = C.gfloat(v.(float64))
					C.clutter_actor_set_translation(actor, translationX, translationY, translationZ)
				case "translation_z":
					translationZ = C.gfloat(v.(float64))
					C.clutter_actor_set_translation(actor, translationX, translationY, translationZ)
				case "visible":
					if v.(bool) {
						C.clutter_actor_show(actor)
					} else {
						C.clutter_actor_hide(actor)
					}
				case "width":
					C.clutter_actor_set_width(actor, C.gfloat(v.(float64)))
				case "x_align":
					switch v.(string) {
					case "fill":
						C.clutter_actor_set_x_align(actor, C.CLUTTER_ACTOR_ALIGN_FILL)
					case "start":
						C.clutter_actor_set_x_align(actor, C.CLUTTER_ACTOR_ALIGN_START)
					case "center":
						C.clutter_actor_set_x_align(actor, C.CLUTTER_ACTOR_ALIGN_CENTER)
					case "end":
						C.clutter_actor_set_x_align(actor, C.CLUTTER_ACTOR_ALIGN_END)
					default:
						log.Fatalf("unknown x align: %s", v.(string))
					}
				case "x_expand":
					b := C.FALSE
					if v.(bool) {
						b = C.TRUE
					}
					C.clutter_actor_set_x_expand(actor, C.gboolean(b))
				case "y_align":
					switch v.(string) {
					case "fill":
						C.clutter_actor_set_y_align(actor, C.CLUTTER_ACTOR_ALIGN_FILL)
					case "start":
						C.clutter_actor_set_y_align(actor, C.CLUTTER_ACTOR_ALIGN_START)
					case "center":
						C.clutter_actor_set_y_align(actor, C.CLUTTER_ACTOR_ALIGN_CENTER)
					case "end":
						C.clutter_actor_set_y_align(actor, C.CLUTTER_ACTOR_ALIGN_END)
					default:
						log.Fatalf("unknown y align: %s", v.(string))
					}
				case "y_expand":
					b := C.FALSE
					if v.(bool) {
						b = C.TRUE
					}
					C.clutter_actor_set_y_expand(actor, C.gboolean(b))
				case "z", "z_position":
					C.clutter_actor_set_z_position(actor, C.gfloat(v.(float64)))
				}
			case float64:
				switch value := v.(type) {
				case unsafe.Pointer:
					//TODO add tests
					if IsActor(value) { // actor child
						C.clutter_actor_add_child(actor, (*C.ClutterActor)(value))
					} else if IsAction(value) { // action
						C.clutter_actor_add_action(actor, (*C.ClutterAction)(value))
					} else if IsConstraint(value) { // constraint
						C.clutter_actor_add_constraint(actor, (*C.ClutterConstraint)(value))
					} else if IsEffect(value) { // effect
						C.clutter_actor_add_effect(actor, (*C.ClutterEffect)(value))
					} else {
						log.Fatalf("unknown subelement type")
					}
				}
			}
		}
		return pointer
	}

	processBoxArgs := func(box *C.ClutterLayoutManager, args map[string]interface{}) {
		//TODO
	}

	lua.RegisterFunctions(map[string]interface{}{

		// actors

		"Actor": func(args map[interface{}]interface{}) unsafe.Pointer {
			actor := C.clutter_actor_new()
			return processActorArgs(actor, args)
		},

		"Stage": func(args map[interface{}]interface{}) unsafe.Pointer {
			actor := C.clutter_stage_new()
			C.clutter_actor_show(actor)
			pointer := processActorArgs(actor, args)
			stage := (*C.ClutterStage)(unsafe.Pointer(actor))
			for k, v := range args {
				//TODO other properties
				switch key := k.(type) {
				case string:
					switch key {
					case "title":
						C.clutter_stage_set_title(stage, toGStr(v.(string)))
					}
				}
			}
			return pointer
		},

		"Text": func(args map[interface{}]interface{}) unsafe.Pointer {
			actor := C.clutter_text_new()
			pointer := processActorArgs(actor, args)
			text := (*C.ClutterText)(unsafe.Pointer(actor))
			for k, v := range args {
				//TODO other properties
				switch key := k.(type) {
				case string:
					switch key {
					case "text":
						C.clutter_text_set_text(text, toGStr(v.(string)))
					case "use_markup":
						b := C.FALSE
						if v.(bool) {
							b = C.TRUE
						}
						C.clutter_text_set_use_markup(text, C.gboolean(b))
					case "markup":
						C.clutter_text_set_markup(text, toGStr(v.(string)))
					case "color":
						C.clutter_text_set_color(text, NewColorFromString(v.(string)))
					}
				case float64:
				}
			}
			return pointer
		},

		// data structures

		"Point": func(args []float64) unsafe.Pointer {
			var point C.ClutterPoint
			point.x = C.float(args[0])
			point.y = C.float(args[1])
			return unsafe.Pointer(&point)
		},

		"Size": func(args []float64) unsafe.Pointer {
			var size C.ClutterSize
			size.width = C.float(args[0])
			size.height = C.float(args[1])
			return unsafe.Pointer(&size)
		},

		"Rect": func(args []float64) unsafe.Pointer {
			var rect C.ClutterRect
			var point C.ClutterPoint
			var size C.ClutterSize
			point.x = C.float(args[0])
			point.y = C.float(args[1])
			size.width = C.float(args[2])
			size.height = C.float(args[3])
			rect.origin = point
			rect.size = size
			return unsafe.Pointer(&rect)
		},

		// contents

		"Image": func(args []string) unsafe.Pointer {
			//TODO set data
			image := C.clutter_image_new()
			return unsafe.Pointer(image)
		},

		// layouts

		"Box": func(args map[string]interface{}) unsafe.Pointer {
			box := C.clutter_box_layout_new()
			processBoxArgs(box, args)
			return unsafe.Pointer(box)
		},

		"HBox": func(args map[string]interface{}) unsafe.Pointer {
			box := C.clutter_box_layout_new()
			C.clutter_box_layout_set_orientation((*C.ClutterBoxLayout)(unsafe.Pointer(box)), C.CLUTTER_ORIENTATION_HORIZONTAL)
			processBoxArgs(box, args)
			return unsafe.Pointer(box)
		},

		"VBox": func(args map[string]interface{}) unsafe.Pointer {
			box := C.clutter_box_layout_new()
			C.clutter_box_layout_set_orientation((*C.ClutterBoxLayout)(unsafe.Pointer(box)), C.CLUTTER_ORIENTATION_VERTICAL)
			processBoxArgs(box, args)
			return unsafe.Pointer(box)
		},

		// effects

		"Blur": func(args map[string]interface{}) unsafe.Pointer {
			effect := C.clutter_blur_effect_new()
			return unsafe.Pointer(effect)
		},

		// constraints
		"Align": func(args map[string]interface{}) unsafe.Pointer {
			constraint := C.clutter_align_constraint_new(nil, C.CLUTTER_ALIGN_X_AXIS, 0)
			align := (*C.ClutterAlignConstraint)(unsafe.Pointer(constraint))
			for key, v := range args {
				switch key {
				case "source":
					C.clutter_align_constraint_set_source(align, (*C.ClutterActor)(bindings[v.(string)]))
				case "axis":
					switch v.(string) {
					case "x", "X":
						C.clutter_align_constraint_set_align_axis(align, C.CLUTTER_ALIGN_X_AXIS)
					case "y", "Y":
						C.clutter_align_constraint_set_align_axis(align, C.CLUTTER_ALIGN_Y_AXIS)
					case "both", "BOTH", "Both":
						C.clutter_align_constraint_set_align_axis(align, C.CLUTTER_ALIGN_BOTH)
					}
				case "factor":
					C.clutter_align_constraint_set_factor(align, C.gfloat(v.(float64)))
				}
			}
			return unsafe.Pointer(constraint)
		},
	})
	lua.RunString(code)
	return bindings
}

func IsActor(o unsafe.Pointer) bool {
	return C.is_actor(o) == C.TRUE
}

func IsAction(o unsafe.Pointer) bool {
	return C.is_action(o) == C.TRUE
}

func IsConstraint(o unsafe.Pointer) bool {
	return C.is_constraint(o) == C.TRUE
}

func IsEffect(o unsafe.Pointer) bool {
	return C.is_effect(o) == C.TRUE
}
