package asset

import (
	"github.com/veandco/go-sdl2/sdl"
)

// asset types
type AssetType int

const (
	Font AssetType = iota
	SoundEffect
	Music
	FontTexture
	Texture
	AnimationData
	Unknown
)

// A request for an asset of the specified type
type AssetRequest struct {
	Id                 string
	Value              string
	FontId             string
	FontSize           int
	Color              sdl.Color
	Type               AssetType
	IsAnimationLooping bool
}

// Generates an asset request for a sdl texture.
func NewTextureRequest(id, value string) *AssetRequest {
	return &AssetRequest{Id: id, Value: value, Type: Texture}
}

// Generates an asset request for a sdl texture.
func NewFontTextureRequest(id, value, fontId string, colorR, colorG, colorB, colorA uint8) *AssetRequest {
	color := sdl.Color{R: colorR, G: colorG, B: colorB, A: colorA}
	return &AssetRequest{Id: id, Value: value, FontId: fontId, Type: FontTexture, Color: color}
}

// Generates an asset request for a sdl font.
func NewFontRequest(id, value string, fontSize int) *AssetRequest {
	return &AssetRequest{Id: id, Value: value, FontSize: fontSize, Type: Font}
}

func NewMusicRequest(id, value string) *AssetRequest {
	return &AssetRequest{Id: id, Value: value, Type: Music}
}
