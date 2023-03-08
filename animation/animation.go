package animation

type Animation struct {
	Name      string
	IsLooping bool
	Frames    *[]AnimationFrame
}

func NewAnimation(name string, isLooping bool, frames *[]AnimationFrame) *Animation {
	return &Animation{Name: name, IsLooping: isLooping, Frames: frames}
}
