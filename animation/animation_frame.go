package animation

import (
	"github.com/veandco/go-sdl2/sdl"
)

type AnimationFrame struct {
	Duration int
	Rect     sdl.Rect
}

func MakeAnimationFrame(x, y, w, h, duration int) AnimationFrame {
	rect := sdl.Rect{X: int32(x), Y: int32(y), W: int32(w), H: int32(h)}
	frame := AnimationFrame{Duration: duration, Rect: rect}
	return frame
}

func NewAnimationFrame(x, y, w, h, duration int32) *AnimationFrame {
	rect := sdl.Rect{X: x, Y: y, W: w, H: h}
	frame := AnimationFrame{Duration: int(duration), Rect: rect}
	return &frame
}
