package animation

import (
	"errors"

	"github.com/veandco/go-sdl2/sdl"
)

type Animator struct {
	frameTimeElapsed, currentFrameIndex int
	animations                          map[string]Animation
	currentAnimation                    *Animation
}

func NewAnimator(animations map[string]Animation) *Animator {
	var animator = Animator{frameTimeElapsed: 0, currentFrameIndex: 0, animations: animations, currentAnimation: nil}
	return &animator
}

func MakeAnimator(animations map[string]Animation) Animator {
	var animator = Animator{frameTimeElapsed: 0, currentFrameIndex: 0, animations: animations, currentAnimation: nil}
	return animator
}

func (animator *Animator) SetCurrentAnimation(animationId *string) error {
	animation, ok := animator.animations[*animationId]

	if ok {
		animator.currentAnimation = &animation
		return nil
	} else {
		return errors.New("'%s' is not an animation loaded into animator")
	}
}

func (animator *Animator) GetCurrentAnimationSourceRect() (*sdl.Rect, error) {
	if animator.currentAnimation == nil {
		return nil, errors.New("animator does not have a current animation set")
	}

	return &animator.currentAnimation.Frames[animator.currentFrameIndex].Rect, nil
}

func (animator *Animator) Update(deltaTicks int) error {
	animator.frameTimeElapsed += deltaTicks

	currentFrameDuration, err := animator.currentAnimation.GetFrameDuration(animator.currentFrameIndex)

	if err != nil {
		return err
	}

	if animator.frameTimeElapsed > currentFrameDuration {
		if animator.currentAnimation.GetNumFrames() <= animator.currentFrameIndex+1 {
			if animator.currentAnimation.IsLooping {
				animator.currentFrameIndex = 0
			}
		} else {
			animator.currentFrameIndex++
		}

		animator.frameTimeElapsed -= currentFrameDuration
	}

	return nil
}
