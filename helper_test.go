package clutter_helper

import (
	"testing"

	g "../g-helper"
)

func TestBasic(t *testing.T) {
	Init()
	bindings := FromLua(`
	Stage {
		id = "stage",
		Actor{
			id = "root",
			bgcolor = "#0099CC",
			width = 300,
			height = 300,
			clip_rect = Rect{0, 0, 300, 200},
			content_gravity = "top-right",
			x = 300,
			y = 50,
			request_mode = "height",
			rotation_angle_x = 30,
			rotation_angle_y = 30,
			rotation_angle_z = 30,
			scale_x = 1.2,
			scale_y = 1.3,
			scale_z = 1.4,
			translation_x = 5,
			translation_y = 6,
			translation_z = 7,
			visible = true,
			layout = Box{
			},
			Actor{
				bgcolor = "#00CC99",
				size = Size{50, 50},
				margin_bottom = 30,
				margin_left = 100,
				margin_right = 5,
				margin_top = 50,
				name = "c1",
			},
			Actor{
				content = Image{"a.png"},
				content_repeat = "both",
			},
			Actor{
				bgcolor = "#9900CC",
				width = 30,
				height = 60,
				offscreen_redirect = "always",
				opacity = 30,
				pivot_point = Point{30, 20},
				pivot_point_z = 3,
				position = Point{30, 30},
				reactive = true,
			},
			Text{
				text_dir = "right",
				text = "你好啊",
			},
		}
	}
	`)
	g.GConnect(bindings["stage"], "destroy", func() {
		Quit()
	})
	Main()
}
