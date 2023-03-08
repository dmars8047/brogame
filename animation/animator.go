package animation

import (
	"errors"

	"github.com/veandco/go-sdl2/sdl"
)

type Animator struct {
	frameTimeElapsed, currentFrameIndex int
	animations                          map[string]*Animation
	currentAnimation                    *Animation
}

func NewAnimator(animations map[string]*Animation) *Animator {
	var animator = Animator{frameTimeElapsed: 0, currentFrameIndex: 0, animations: animations, currentAnimation: nil}
	return &animator
}

func (animator *Animator) SetCurrentAnimation(animationId *string) error {
	animation, ok := animator.animations[*animationId]

	if ok {
		animator.currentAnimation = animation
		return nil
	} else {
		return errors.New("'%s' is not an animation loaded into animator")
	}
}

func (animator *Animator) GetCurrentAnimationSourceRect() (*sdl.Rect, error) {
	if animator.currentAnimation == nil {
		return nil, errors.New("animator does not have a current animation set.")
	}

	frames := *animator.currentAnimation.Frames

	return &frames[animator.currentFrameIndex].Rect, nil
}

func (animator *Animator) Update(deltaTicks int) {

}
