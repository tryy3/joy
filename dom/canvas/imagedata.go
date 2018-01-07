package canvas

import "github.com/matthewmueller/joy/macro"

// NewImageData fn
func NewImageData(width uint, height uint) *ImageData {
	macro.Rewrite("new ImageData($1, $2)", width, height)
	return &ImageData{}
}

// ImageData struct
// js:"ImageData,omit"
type ImageData struct {
}

// Data prop
// js:"data"
func (*ImageData) Data() (data []uint8) {
	macro.Rewrite("$_.data")
	return data
}

// Height prop
// js:"height"
func (*ImageData) Height() (height uint) {
	macro.Rewrite("$_.height")
	return height
}

// Width prop
// js:"width"
func (*ImageData) Width() (width uint) {
	macro.Rewrite("$_.width")
	return width
}
