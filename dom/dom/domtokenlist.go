package dom

type DOMTokenList interface {
	Add(token string)

	Contains(token string) (b bool)

	Item(index uint) (s string)

	Remove(token string)

	Toggle(token string, force *bool) (b bool)

	ToString() (s string)

	Length() (length uint)
}
