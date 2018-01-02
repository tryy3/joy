package element

type ElementTraversal interface {
	ChildElementCount() (childElementCount uint)

	FirstElementChild() (firstElementChild Element)

	LastElementChild() (lastElementChild Element)

	NextElementSibling() (nextElementSibling Element)

	PreviousElementSibling() (previousElementSibling Element)
}
