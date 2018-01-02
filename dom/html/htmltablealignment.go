package html

type HTMLTableAlignment interface {
	Ch() (ch string)

	SetCh(ch string)

	ChOff() (chOff string)

	SetChOff(chOff string)

	VAlign() (vAlign string)

	SetVAlign(vAlign string)
}
