package animation

import "errors"

type Animation struct {
	Name      string
	IsLooping bool
	Frames    []AnimationFrame
}

func NewAnimation(name string, isLooping bool, frames *[]AnimationFrame) *Animation {
	return &Animation{Name: name, IsLooping: isLooping, Frames: *frames}
}

func (animation *Animation) GetFrameDuration(index int) (int, error) {
	if index >= len(animation.Frames) || index < 0 {
		return 0, errors.New("index used for frame duration retrieval is out-of-range")
	}

	return animation.Frames[index].Duration, nil
}

func (animation *Animation) GetNumFrames() int {
	return len(animation.Frames)
}
