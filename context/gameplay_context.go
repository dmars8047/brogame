package context

import "github.com/veandco/go-sdl2/sdl"

type GameplayContext struct {
	CurrentFrameTicks       uint64
	PreviousFrameTicks      uint64
	DeltaTicks              uint64
	Paused                  bool
	CurrentScene            string
	KeysPressed             []sdl.Scancode
	ControllerButtonPressed []sdl.GameControllerButton
	GameController          *sdl.GameController
}

func NewGameplayContext(controller *sdl.GameController) *GameplayContext {

	context := GameplayContext{CurrentFrameTicks: 0,
		PreviousFrameTicks:      0,
		Paused:                  false,
		CurrentScene:            "NOT_SET",
		KeysPressed:             make([]sdl.Scancode, 0),
		ControllerButtonPressed: make([]sdl.GameControllerButton, 0),
		GameController:          controller}

	return &context
}

func (context *GameplayContext) StartNewFrame() {

	context.PreviousFrameTicks = context.CurrentFrameTicks
	context.CurrentFrameTicks = sdl.GetTicks64()
	context.DeltaTicks = context.CurrentFrameTicks - context.PreviousFrameTicks
	context.KeysPressed = nil
	context.ControllerButtonPressed = nil
}

func (context *GameplayContext) IsKeyPressed(scancode sdl.Scancode, includeHeldKey bool) bool {
	if includeHeldKey {
		keyBoardState := sdl.GetKeyboardState()

		result := keyBoardState[scancode]

		return result == 1
	} else {

		var found bool = false

		for _, val := range context.KeysPressed {
			if val == scancode {
				found = true
				break
			}
		}

		return found
	}
}
