package scene

import (
	"github.com/dmars8047/brogame/context"
	"github.com/veandco/go-sdl2/sdl"
)

type Scene[T any] interface {
	Init(renderer *sdl.Renderer)
	GetName() string
	Update(context *context.GameplayContext, gameParams T)
	Render(renderer *sdl.Renderer)
	Unload()
}
