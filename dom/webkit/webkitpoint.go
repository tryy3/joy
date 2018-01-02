package webkit

import "github.com/matthewmueller/joy/macro"

func New(x *float32, y *float32) *WebKitPoint {
	macro.Rewrite("new WebKitPoint($1, $2)", x, y)
	return &WebKitPoint{}
}

type WebKitPoint struct {
}

func (*WebKitPoint) X() (x float32) {
	macro.Rewrite("$_.x")
	return x
}

func (*WebKitPoint) SetX(x float32) {
	macro.Rewrite("$_.x = $1", x)
}

func (*WebKitPoint) Y() (y float32) {
	macro.Rewrite("$_.y")
	return y
}

func (*WebKitPoint) SetY(y float32) {
	macro.Rewrite("$_.y = $1", y)
}
