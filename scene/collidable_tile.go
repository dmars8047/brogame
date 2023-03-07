package scene

import "github.com/veandco/go-sdl2/sdl"

type CollidableTile struct {
	rect sdl.Rect
}

func NewCollidableTile(x, y, w, h int32) *CollidableTile {
	var tile = CollidableTile{
		rect: sdl.Rect{X: x, Y: y, W: w, H: h},
	}

	return &tile
}

func (tile *CollidableTile) IsColliding(rect *sdl.Rect) bool {

	if tile.rect.X == rect.X && tile.rect.Y == rect.Y {
		return true
	}

	if tile.rect.X <= rect.X && tile.rect.X+tile.rect.W > rect.X {
		if tile.rect.Y <= rect.Y && tile.rect.Y+tile.rect.H > rect.Y {
			return true
		}
	}

	if rect.X <= tile.rect.X && rect.X+rect.W > tile.rect.X {
		if rect.Y <= tile.rect.Y && rect.Y+rect.H > tile.rect.Y {
			return true
		}
	}

	return false
}
