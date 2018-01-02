package canvas

import "github.com/matthewmueller/joy/macro"

func New(width uint, height uint) *ImageData {
	macro.Rewrite("new ImageData($1, $2)", width, height)
	return &ImageData{}
}

type ImageData struct {
}

func (*ImageData) Data() (data []uint8) {
	macro.Rewrite("$_.data")
	return data
}

func (*ImageData) Height() (height uint) {
	macro.Rewrite("$_.height")
	return height
}

func (*ImageData) Width() (width uint) {
	macro.Rewrite("$_.width")
	return width
}
