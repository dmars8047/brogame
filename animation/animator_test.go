package animation_test

import (
	"testing"

	"github.com/dmars8047/brogame/animation"
)

func TestNewAnimator(t *testing.T) {
	animationId := "testAnimation"
	animationMap := make(map[string]animation.Animation)

	frames := make([]animation.AnimationFrame, 1)

	frames[0] = animation.MakeAnimationFrame(0, 0, 0, 0, 100)

	animationMap[animationId] = animation.MakeAnimation(animationId, false, frames)

	animator := animation.NewAnimator(animationMap)

	if animator == nil {
		t.Errorf("NewAnimator produced nil result")
	}
}

func TestSetCurrentAnimation(t *testing.T) {
	animationId_1 := "testAnimation1"
	animationId_2 := "testAnimation2"
	animationMap := make(map[string]animation.Animation, 2)

	frames := make([]animation.AnimationFrame, 2)

	frames[0] = animation.MakeAnimationFrame(0, 0, 0, 0, 100)

	animationMap[animationId_1] = animation.MakeAnimation(animationId_1, false, frames)
	animationMap[animationId_2] = animation.MakeAnimation(animationId_2, false, frames)

	animator := animation.MakeAnimator(animationMap)

	err := animator.SetCurrentAnimation(&animationId_1)

	if err != nil {
		t.Errorf("animator.SetCurrentAnimation(%s) resulted in error: %s", animationId_1, err.Error())
	}

	err = animator.SetCurrentAnimation(&animationId_2)

	if err != nil {
		t.Errorf("animator.SetCurrentAnimation(%s) resulted in error: %s", animationId_2, err.Error())
	}
}

func BenchmarkMakeAnimator(b *testing.B) {
	animationId := "testAnimation"
	animationMap := make(map[string]animation.Animation)

	frames := make([]animation.AnimationFrame, 1)

	frames[0] = animation.MakeAnimationFrame(0, 0, 0, 0, 100)

	animationMap[animationId] = animation.MakeAnimation(animationId, false, frames)

	for i := 0; i < b.N; i++ {
		animation.MakeAnimator(animationMap)
	}
}

func BenchmarkNewAnimator(b *testing.B) {
	animationId := "testAnimation"
	animationMap := make(map[string]animation.Animation)

	frames := make([]animation.AnimationFrame, 1)

	frames[0] = animation.MakeAnimationFrame(0, 0, 0, 0, 100)

	animationMap[animationId] = animation.MakeAnimation(animationId, false, frames)

	for i := 0; i < b.N; i++ {
		animation.NewAnimator(animationMap)
	}
}
