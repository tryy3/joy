package mediastreams

type ConstrainLongRange struct {
	*LongRange

	exact	*int
	ideal	*int
}
