package mediastreams


import "github.com/matthewmueller/joy/dom/doublerange"

type ConstrainDoubleRange struct {
	*doublerange.DoubleRange

	exact *float32
	ideal *float32
}
