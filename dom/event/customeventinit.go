package event

type CustomEventInit struct {
	*EventInit

	detail	*interface{}
}
