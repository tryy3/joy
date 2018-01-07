package animation

import "github.com/matthewmueller/joy/dom/event"

type AnimationEventInit struct {
	*event.EventInit

	animationName *string
	elapsedTime   *float32
}
