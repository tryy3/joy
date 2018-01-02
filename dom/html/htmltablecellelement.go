package html

type HTMLTableCellElement interface {
	HTMLElement

	Ch() (ch string)

	SetCh(ch string)

	ChOff() (chOff string)

	SetChOff(chOff string)

	VAlign() (vAlign string)

	SetVAlign(vAlign string)

	Abbr() (abbr string)

	SetAbbr(abbr string)

	Align() (align string)

	SetAlign(align string)

	Axis() (axis string)

	SetAxis(axis string)

	BgColor() (bgColor interface{})

	SetBgColor(bgColor interface{})

	CellIndex() (cellIndex int)

	ColSpan() (colSpan uint)

	SetColSpan(colSpan uint)

	Headers() (headers string)

	SetHeaders(headers string)

	Height() (height interface{})

	SetHeight(height interface{})

	NoWrap() (noWrap bool)

	SetNoWrap(noWrap bool)

	RowSpan() (rowSpan uint)

	SetRowSpan(rowSpan uint)

	Scope() (scope string)

	SetScope(scope string)

	Width() (width string)

	SetWidth(width string)
}
