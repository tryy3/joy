package dom

import "github.com/matthewmueller/joy/macro"

type ClientRect struct {
}

func (*ClientRect) Bottom() (bottom int) {
	macro.Rewrite("$_.bottom")
	return bottom
}

func (*ClientRect) SetBottom(bottom int) {
	macro.Rewrite("$_.bottom = $1", bottom)
}

func (*ClientRect) Height() (height float32) {
	macro.Rewrite("$_.height")
	return height
}

func (*ClientRect) Left() (left int) {
	macro.Rewrite("$_.left")
	return left
}

func (*ClientRect) SetLeft(left int) {
	macro.Rewrite("$_.left = $1", left)
}

func (*ClientRect) Right() (right int) {
	macro.Rewrite("$_.right")
	return right
}

func (*ClientRect) SetRight(right int) {
	macro.Rewrite("$_.right = $1", right)
}

func (*ClientRect) Top() (top int) {
	macro.Rewrite("$_.top")
	return top
}

func (*ClientRect) SetTop(top int) {
	macro.Rewrite("$_.top = $1", top)
}

func (*ClientRect) Width() (width float32) {
	macro.Rewrite("$_.width")
	return width
}
