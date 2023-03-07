package scene_test

import (
	"fmt"
	"testing"

	"github.com/dmars8047/brogame/scene"
	"github.com/veandco/go-sdl2/sdl"
)

func TestNewCollidableTile(t *testing.T) {
	var tests = []struct {
		x, y, w, h int32
	}{
		{0, 0, 0, 0},
		{32, 48, 64, 72},
		{-32, -48, -64, -72},
	}

	for _, test := range tests {
		testname := fmt.Sprintf("x:%d,y:%d,w:%d,h:%d}", test.x, test.y, test.w, test.h)

		t.Run(testname, func(t *testing.T) {
			tile := scene.NewCollidableTile(test.x, test.y, test.w, test.h)
			if tile == nil {
				t.Errorf("Collidable tile is nil.")
			}
		})
	}
}

func TestIsColliding(t *testing.T) {

	var tests = []struct {
		a, b   sdl.Rect
		result bool
	}{
		{
			a:      sdl.Rect{X: 0, Y: 0, W: 32, H: 32},
			b:      sdl.Rect{X: 0, Y: 0, W: 32, H: 32},
			result: true,
		},
		{
			a:      sdl.Rect{X: 0, Y: 0, W: 32, H: 32},
			b:      sdl.Rect{X: 1, Y: 1, W: 32, H: 32},
			result: true,
		},
		{
			a:      sdl.Rect{X: 1, Y: 1, W: 32, H: 32},
			b:      sdl.Rect{X: 0, Y: 0, W: 32, H: 32},
			result: true,
		},
		{
			a:      sdl.Rect{X: 64, Y: 64, W: 32, H: 32},
			b:      sdl.Rect{X: 0, Y: 0, W: 32, H: 32},
			result: false,
		},
		{
			a:      sdl.Rect{X: 32, Y: 0, W: 32, H: 32},
			b:      sdl.Rect{X: 0, Y: 0, W: 32, H: 32},
			result: false,
		},
		{
			a:      sdl.Rect{X: 0, Y: 0, W: 32, H: 32},
			b:      sdl.Rect{X: 64, Y: 64, W: 32, H: 32},
			result: false,
		},
		{
			a:      sdl.Rect{X: 0, Y: 0, W: 32, H: 32},
			b:      sdl.Rect{X: 32, Y: 0, W: 32, H: 32},
			result: false,
		},
		{
			a:      sdl.Rect{X: 0, Y: 0, W: 64, H: 32},
			b:      sdl.Rect{X: 32, Y: 0, W: 32, H: 32},
			result: true,
		},
		{
			a:      sdl.Rect{X: 32, Y: 0, W: 32, H: 32},
			b:      sdl.Rect{X: 0, Y: 0, W: 64, H: 32},
			result: true,
		},
	}

	for _, test := range tests {
		var collidableTile = scene.NewCollidableTile(test.a.X, test.a.Y, test.a.W, test.a.H)

		testname := fmt.Sprintf("A{x:%d,y:%d,w:%d,h:%d}_B{x:%d,y:%d,w:%d,h:%d}",
			test.a.X, test.a.Y, test.a.W, test.a.H, test.b.X, test.b.Y, test.b.W, test.b.H)

		t.Run(testname, func(t *testing.T) {
			result := collidableTile.IsColliding(&test.b)
			if result != test.result {
				t.Errorf("got %t, expected %t", result, test.result)
			}
		})
	}
}
