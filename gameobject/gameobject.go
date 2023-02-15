package gameobject

import (
	"github.com/dmars8047/brogame/context"
	"github.com/veandco/go-sdl2/sdl"
)

type GameObject interface {
	IsUpdateable() bool
	Update(context *context.GameplayContext)
	IsRenderable() bool
	Render(renderer *sdl.Renderer)
	IsCollidable() bool
	HandleCollision(object *GameObject)
}
