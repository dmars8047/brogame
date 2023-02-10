package brogame

import (
	"fmt"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

const APP_NAME = "TEST"
const SCREEN_WIDTH = 960
const SCREEN_HEIGHT = 768
const BACKGROUND_RED = 17
const BACKGROUND_GREEN = 17
const BACKGROUND_BLUE = 17
const BACKGROUND_ALPHA = 255

var window *sdl.Window
var renderer *sdl.Renderer
var gameController *sdl.GameController

func initSDL() {

	// initilize sdl
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Println("Error: could not initilize SDL!")
		panic(err)
	}

	// setup render scaling to use
	if !sdl.SetHint(sdl.HINT_RENDER_SCALE_QUALITY, "2") {
		fmt.Println("Warning: Best texture filtering not enabled!")
	} else if !sdl.SetHint(sdl.HINT_RENDER_SCALE_QUALITY, "1") {
		fmt.Println("Warning: Linear texture filtering not enabled!")
	}

	// set controller if available (just the first one)

	if sdl.NumJoysticks() < 1 {
		fmt.Println("No controllers/joysticks connected!")
	} else {
		gameController = sdl.GameControllerOpen(0)

		if gameController == nil {
			err := sdl.GetError()
			if err != nil {
				fmt.Printf("Warning: Detected controller could not be opened! Error: %s\n", err.Error())
			} else {
				fmt.Println("Warning: Detected controller could not be opened!")
			}
		}
	}

	// initialize window
	var windowErr error

	window, windowErr = sdl.CreateWindow(APP_NAME, sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED,
		SCREEN_WIDTH, SCREEN_HEIGHT, sdl.WINDOW_SHOWN)

	if windowErr != nil {
		fmt.Println("Error: could not initilize window!")
		panic(windowErr)
	}

	var rendererErr error
	var renderFlags uint32 = sdl.RENDERER_PRESENTVSYNC | sdl.RENDERER_ACCELERATED | sdl.RENDERER_SOFTWARE

	renderer, rendererErr = sdl.CreateRenderer(window, 0, renderFlags)

	if rendererErr != nil {
		fmt.Println("Error: SDL Renderer could not be created!")
		panic(rendererErr)
	}

	// set default draw color of the renderer
	renderer.SetDrawColor(BACKGROUND_RED, BACKGROUND_GREEN, BACKGROUND_BLUE, BACKGROUND_ALPHA)

	// initialize SDL_IMAGE
	if err := img.Init(img.INIT_PNG); err != nil {
		fmt.Println("Error: could not initialize SDL_IMAGE!")
		panic(err)
	}

	// initialize SDL_TTF
	if err := ttf.Init(); err != nil {
		fmt.Println("Error: could not initialize SDL_TTF!")
		panic(err)
	}

	// initialize SDL_MIXER with CD quality audio, stereo sound
	if err := mix.OpenAudio(44100, mix.DEFAULT_FORMAT, 2, 2048); err != nil {
		fmt.Println("Error: could not initialize SDL_MIXER!")
		panic(err)
	}
}

func close() {
	fmt.Println("Closing SDL!")

	if gameController != nil {
		gameController.Close()
	}

	if renderer != nil {
		renderer.Destroy()
	}
	if window != nil {
		window.Destroy()
	}

	ttf.Quit()
	mix.Quit()
	img.Quit()
	sdl.Quit()
}

func main() {
	defer close()
	initSDL()

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.KeyboardEvent:
				if t.Keysym.Sym == sdl.K_ESCAPE {
					running = false
				} else {
					break
				}
			case *sdl.QuitEvent:
				running = false
			}

			// clear screen
			renderer.Clear()

			renderer.Present()
		}
	}
}
