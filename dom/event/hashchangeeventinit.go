package event

type HashChangeEventInit struct {
	*EventInit

	newURL	*string
	oldURL	*string
}
