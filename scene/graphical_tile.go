package scene

import (
	"github.com/veandco/go-sdl2/sdl"
)

type GraphicalTile struct {
	sourceRect, destRect sdl.Rect
	texture              *sdl.Texture
}

func NewGraphicalTile(sourceX, sourceY, destX, destY, tileSize int32, texture *sdl.Texture) *GraphicalTile {

	source := sdl.Rect{X: sourceX, Y: sourceY, W: tileSize, H: tileSize}
	destination := sdl.Rect{X: destX, Y: destY, W: tileSize, H: tileSize}

	tile := GraphicalTile{source, destination, texture}

	return &tile
}

func (tile *GraphicalTile) Render(renderer *sdl.Renderer) {
	renderer.Copy(tile.texture, &tile.sourceRect, &tile.destRect)
}
