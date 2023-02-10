package asset

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/dmars8047/brogame/animation"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type AssetManager struct {
	Textures      map[string]*sdl.Texture
	Fonts         map[string]*ttf.Font
	Music         map[string]*mix.Music
	SoundEffects  map[string]*mix.Chunk
	AnimationData map[string]*animation.Animation
}

// Creates and initializes a new AssetManager.
func NewAssetManager() *AssetManager {

	assetManager := AssetManager{
		Textures:      make(map[string]*sdl.Texture),
		Fonts:         make(map[string]*ttf.Font),
		Music:         make(map[string]*mix.Music),
		SoundEffects:  make(map[string]*mix.Chunk),
		AnimationData: make(map[string]*animation.Animation),
	}

	return &assetManager
}

// Loads a texture into the asset manager. If it already exists, then no action will be taken unless the overite parameter is set to 'true'.
func (assetManager *AssetManager) LoadTexture(renderer *sdl.Renderer, request *AssetRequest, overwrite bool) error {

	// Make sure the request is for a texture type
	if request.Type != Texture && request.Type != FontTexture {
		return errors.New("asset load request is not a texture type")
	}

	value, present := assetManager.Textures[request.Id]

	if present {
		if overwrite {
			err := value.Destroy()
			if err != nil {
				return err
			}

			delete(assetManager.Textures, request.Id)
		} else {
			// texture already loaded
			return nil
		}
	}

	if request.Type == Texture {
		surface, loadErr := img.Load(request.Value)

		if loadErr != nil {
			return fmt.Errorf("unable to load image '%s'.  SDL_image Error: %s", request.Value, loadErr.Error())
		}

		texture, textureErr := renderer.CreateTextureFromSurface(surface)

		if textureErr != nil {
			return fmt.Errorf("unable to create texture from '%s'. SDL Error: %s", request.Value, textureErr.Error())
		}

		surface.Free()
		assetManager.Textures[request.Id] = texture

	} else if request.Type == FontTexture {

		font := assetManager.Fonts[request.FontId]

		surface, surfaceErr := font.RenderUTF8Solid(request.Value, request.Color)

		if surfaceErr != nil {
			return fmt.Errorf("unable to render text surface! SDL_ttf Error: %s", surfaceErr.Error())
		}

		texture, textureErr := renderer.CreateTextureFromSurface(surface)

		if textureErr != nil {
			return fmt.Errorf("unable to create texture from rendered text: '%s'. SDL_ttf error: %s", request.Value, textureErr.Error())
		}

		surface.Free()
		assetManager.Textures[request.Id] = texture
	}

	return nil
}

// Clears all textures and frees their memory. Errors only arise when destroying pointers to nil or invalid textures.
func (assetManager *AssetManager) ClearTextures() error {
	for key, val := range assetManager.Textures {
		if val != nil {
			err := val.Destroy()
			if err != nil {
				return err
			}

			delete(assetManager.Textures, key)
		}
	}

	return nil
}

// Loads a font into the asset manager. If it already exists, then no action will be taken unless the overwite parameter is 'true'.
func (assetManager *AssetManager) LoadFont(request *AssetRequest, overwrite bool) error {

	if request.Type != Font {
		return errors.New("asset load request is not a font type")
	}

	font, present := assetManager.Fonts[request.Id]

	if present {
		if overwrite {
			font.Close()
			delete(assetManager.Fonts, request.Id)
		} else {
			// font already loaded
			return nil
		}
	}

	font, fontErr := ttf.OpenFont(request.Value, request.FontSize)

	if fontErr != nil {
		return fmt.Errorf("failed to load font: '%s'. SDL error: %s", request.Value, fontErr.Error())
	}

	assetManager.Fonts[request.Id] = font

	return nil
}

// Clears all fonts and frees their memory.
func (assetManager *AssetManager) ClearFonts() {
	for k, v := range assetManager.Fonts {
		v.Close()
		delete(assetManager.Fonts, k)
	}
}

// Loads music into the asset manager. If it already exists, then no action will be taken unless the overwite parameter is 'true'.
func (assetManager *AssetManager) LoadMusic(request *AssetRequest, overwrite bool) error {

	if request.Type != Music {
		return fmt.Errorf("asset load request is not a music type")
	}

	music, present := assetManager.Music[request.Id]

	if present {
		if overwrite {
			music.Free()
			delete(assetManager.Music, request.Id)
		} else {
			// music already loaded
			return nil
		}
	}

	music, loadErr := mix.LoadMUS(request.Value)

	if loadErr != nil {
		return fmt.Errorf("failed to load music! SDL_mixer Error: %s", loadErr)
	}

	assetManager.Music[request.Id] = music

	return nil
}

// Clears all music assets and frees their memory
func (assetManager *AssetManager) ClearMusic() {
	for k, v := range assetManager.Music {
		v.Free()
		delete(assetManager.Music, k)
	}
}

// Loads a sound effect into the asset manager. If it already exists, then no action will be taken unless the overwite parameter is 'true'.
func (assetManager *AssetManager) LoadSoundEffect(request *AssetRequest, overwrite bool) error {

	if request.Type != SoundEffect {
		return errors.New("asset load request is not a sound effect type")
	}

	sfxChunk, present := assetManager.SoundEffects[request.Id]

	if present {
		if overwrite {
			sfxChunk.Free()
			delete(assetManager.SoundEffects, request.Id)
		} else {
			// sound effect already loaded
			return nil
		}
	}

	sfxChunk, loadErr := mix.LoadWAV(request.Value)

	if loadErr != nil {
		return fmt.Errorf("failed to load sound effect! SDL_mixer error: %s", loadErr.Error())
	}

	assetManager.SoundEffects[request.Id] = sfxChunk

	return nil
}

// Clears all sound effect assets and frees their memory.
func (assetManager *AssetManager) ClearSoundEffects() {
	for k, v := range assetManager.SoundEffects {
		v.Free()
		delete(assetManager.SoundEffects, k)
	}
}

// Loads animation data into the asset manager. If it already exists, then no action will be taken unless the overwite parameter is 'true'. Expects an animation data file which is in the 'bro' format.
func (assetManager *AssetManager) LoadAnimation(request *AssetRequest, overwrite bool) error {

	if request.Type != AnimationData {
		return errors.New("asset load request is not an animation type")
	}

	_, present := assetManager.AnimationData[request.Id]

	if present {
		if overwrite {
			delete(assetManager.AnimationData, request.Id)
		} else {
			// animation already loaded
			return nil
		}
	}

	file, err := os.Open(request.Value)

	if err != nil {
		return fmt.Errorf("%s could not be opened", request.Value)
	}

	defer file.Close()

	frames := make([]animation.AnimationFrame, 0)

	scanner := bufio.NewScanner(file)
	lineNum := 1
	for scanner.Scan() {
		if lineNum > 2 {
			rawAnimationData := strings.Split(scanner.Text(), ",")
			parsedAnimationData, err := parseRawAnimationData(&rawAnimationData)

			if err != nil {
				return fmt.Errorf("parsing error occured in animation file '%s' on line %d", request.Value, lineNum)
			}

			// # FORMAT: FRAME_X, FRAME_Y, FRAME_W, FRAME_H, DURATION
			frames = append(frames, animation.NewAnimationFrame(parsedAnimationData[0], parsedAnimationData[1], parsedAnimationData[2], parsedAnimationData[3], parsedAnimationData[4]))
		}
		lineNum++
	}

	assetManager.AnimationData[request.Id] = animation.NewAnimation(request.Id, request.IsAnimationLooping, &frames)

	return nil
}

// Clears all animations from the asset manager. If it already exists, then no action will be taken unless the overwite parameter is 'true'.
func (assetManager *AssetManager) ClearAnimations(request *AssetRequest, overwrite bool) {
	for k := range assetManager.AnimationData {
		delete(assetManager.AnimationData, k)
	}
}

func parseRawAnimationData(rawAnimationData *[]string) (*[5]int, error) {
	data := [5]int{0, 0, 0, 0, 0}

	for i, rawVal := range *rawAnimationData {
		val, err := strconv.Atoi(rawVal)

		if err != nil {
			return nil, err
		}

		data[i] = val
	}

	return &data, nil
}
