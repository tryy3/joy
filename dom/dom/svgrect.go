package dom

import "github.com/matthewmueller/joy/macro"

type SVGRect struct {
}

func (*SVGRect) Height() (height float32) {
	macro.Rewrite("$_.height")
	return height
}

func (*SVGRect) SetHeight(height float32) {
	macro.Rewrite("$_.height = $1", height)
}

func (*SVGRect) Width() (width float32) {
	macro.Rewrite("$_.width")
	return width
}

func (*SVGRect) SetWidth(width float32) {
	macro.Rewrite("$_.width = $1", width)
}

func (*SVGRect) X() (x float32) {
	macro.Rewrite("$_.x")
	return x
}

func (*SVGRect) SetX(x float32) {
	macro.Rewrite("$_.x = $1", x)
}

func (*SVGRect) Y() (y float32) {
	macro.Rewrite("$_.y")
	return y
}

func (*SVGRect) SetY(y float32) {
	macro.Rewrite("$_.y = $1", y)
}
