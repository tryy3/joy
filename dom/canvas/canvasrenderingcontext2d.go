package canvas

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/element"
	"github.com/matthewmueller/joy/dom/html"
)

type CanvasRenderingContext2d struct {
}

func (*CanvasRenderingContext2d) BeginPath() {
	macro.Rewrite("$_.beginPath()")
}

func (*CanvasRenderingContext2d) ClearRect(x float32, y float32, w float32, h float32) {
	macro.Rewrite("$_.clearRect($1, $2, $3, $4)", x, y, w, h)
}

func (*CanvasRenderingContext2d) Clip(fillRule *CanvasFillRule) {
	macro.Rewrite("$_.clip($1)", fillRule)
}

func (*CanvasRenderingContext2d) CreateImageData(imageDataOrSw interface{}, sh *float32) (i *ImageData) {
	macro.Rewrite("$_.createImageData($1, $2)", imageDataOrSw, sh)
	return i
}

func (*CanvasRenderingContext2d) CreateLinearGradient(x0 float32, y0 float32, x1 float32, y1 float32) (c *CanvasGradient) {
	macro.Rewrite("$_.createLinearGradient($1, $2, $3, $4)", x0, y0, x1, y1)
	return c
}

func (*CanvasRenderingContext2d) CreatePattern(image interface{}, repetition string) (c *CanvasPattern) {
	macro.Rewrite("$_.createPattern($1, $2)", image, repetition)
	return c
}

func (*CanvasRenderingContext2d) CreateRadialGradient(x0 float32, y0 float32, r0 float32, x1 float32, y1 float32, r1 float32) (c *CanvasGradient) {
	macro.Rewrite("$_.createRadialGradient($1, $2, $3, $4, $5, $6)", x0, y0, r0, x1, y1, r1)
	return c
}

func (*CanvasRenderingContext2d) DrawFocusIfNeeded(element element.Element) {
	macro.Rewrite("$_.drawFocusIfNeeded($1)", element)
}

func (*CanvasRenderingContext2d) DrawImage(image interface{}, offsetX float32, offsetY float32, width *float32, height *float32, canvasOffsetX *float32, canvasOffsetY *float32, canvasImageWidth *float32, canvasImageHeight *float32) {
	macro.Rewrite("$_.drawImage($1, $2, $3, $4, $5, $6, $7, $8, $9)", image, offsetX, offsetY, width, height, canvasOffsetX, canvasOffsetY, canvasImageWidth, canvasImageHeight)
}

func (*CanvasRenderingContext2d) Fill(fillRule *CanvasFillRule) {
	macro.Rewrite("$_.fill($1)", fillRule)
}

func (*CanvasRenderingContext2d) FillRect(x float32, y float32, w float32, h float32) {
	macro.Rewrite("$_.fillRect($1, $2, $3, $4)", x, y, w, h)
}

func (*CanvasRenderingContext2d) FillText(text string, x float32, y float32, maxWidth *float32) {
	macro.Rewrite("$_.fillText($1, $2, $3, $4)", text, x, y, maxWidth)
}

func (*CanvasRenderingContext2d) GetImageData(sx float32, sy float32, sw float32, sh float32) (i *ImageData) {
	macro.Rewrite("$_.getImageData($1, $2, $3, $4)", sx, sy, sw, sh)
	return i
}

func (*CanvasRenderingContext2d) GetLineDash() (f []float32) {
	macro.Rewrite("$_.getLineDash()")
	return f
}

func (*CanvasRenderingContext2d) IsPointInPath(x float32, y float32, fillRule *CanvasFillRule) (b bool) {
	macro.Rewrite("$_.isPointInPath($1, $2, $3)", x, y, fillRule)
	return b
}

func (*CanvasRenderingContext2d) MeasureText(text string) (t *TextMetrics) {
	macro.Rewrite("$_.measureText($1)", text)
	return t
}

func (*CanvasRenderingContext2d) PutImageData(imagedata *ImageData, dx float32, dy float32, dirtyX *float32, dirtyY *float32, dirtyWidth *float32, dirtyHeight *float32) {
	macro.Rewrite("$_.putImageData($1, $2, $3, $4, $5, $6, $7)", imagedata, dx, dy, dirtyX, dirtyY, dirtyWidth, dirtyHeight)
}

func (*CanvasRenderingContext2d) Restore() {
	macro.Rewrite("$_.restore()")
}

func (*CanvasRenderingContext2d) Rotate(angle float32) {
	macro.Rewrite("$_.rotate($1)", angle)
}

func (*CanvasRenderingContext2d) Save() {
	macro.Rewrite("$_.save()")
}

func (*CanvasRenderingContext2d) Scale(x float32, y float32) {
	macro.Rewrite("$_.scale($1, $2)", x, y)
}

func (*CanvasRenderingContext2d) SetLineDash(segments []float32) {
	macro.Rewrite("$_.setLineDash($1)", segments)
}

func (*CanvasRenderingContext2d) SetTransform(m11 float32, m12 float32, m21 float32, m22 float32, dx float32, dy float32) {
	macro.Rewrite("$_.setTransform($1, $2, $3, $4, $5, $6)", m11, m12, m21, m22, dx, dy)
}

func (*CanvasRenderingContext2d) Stroke(path *Path2d) {
	macro.Rewrite("$_.stroke($1)", path)
}

func (*CanvasRenderingContext2d) StrokeRect(x float32, y float32, w float32, h float32) {
	macro.Rewrite("$_.strokeRect($1, $2, $3, $4)", x, y, w, h)
}

func (*CanvasRenderingContext2d) StrokeText(text string, x float32, y float32, maxWidth *float32) {
	macro.Rewrite("$_.strokeText($1, $2, $3, $4)", text, x, y, maxWidth)
}

func (*CanvasRenderingContext2d) Transform(m11 float32, m12 float32, m21 float32, m22 float32, dx float32, dy float32) {
	macro.Rewrite("$_.transform($1, $2, $3, $4, $5, $6)", m11, m12, m21, m22, dx, dy)
}

func (*CanvasRenderingContext2d) Translate(x float32, y float32) {
	macro.Rewrite("$_.translate($1, $2)", x, y)
}

func (*CanvasRenderingContext2d) Arc(x float32, y float32, radius float32, startAngle float32, endAngle float32, anticlockwise *bool) {
	macro.Rewrite("$_.arc($1, $2, $3, $4, $5, $6)", x, y, radius, startAngle, endAngle, anticlockwise)
}

func (*CanvasRenderingContext2d) ArcTo(x1 float32, y1 float32, x2 float32, y2 float32, radius float32) {
	macro.Rewrite("$_.arcTo($1, $2, $3, $4, $5)", x1, y1, x2, y2, radius)
}

func (*CanvasRenderingContext2d) BezierCurveTo(cp1x float32, cp1y float32, cp2x float32, cp2y float32, x float32, y float32) {
	macro.Rewrite("$_.bezierCurveTo($1, $2, $3, $4, $5, $6)", cp1x, cp1y, cp2x, cp2y, x, y)
}

func (*CanvasRenderingContext2d) ClosePath() {
	macro.Rewrite("$_.closePath()")
}

func (*CanvasRenderingContext2d) Ellipse(x float32, y float32, radiusX float32, radiusY float32, rotation float32, startAngle float32, endAngle float32, anticlockwise *bool) {
	macro.Rewrite("$_.ellipse($1, $2, $3, $4, $5, $6, $7, $8)", x, y, radiusX, radiusY, rotation, startAngle, endAngle, anticlockwise)
}

func (*CanvasRenderingContext2d) LineTo(x float32, y float32) {
	macro.Rewrite("$_.lineTo($1, $2)", x, y)
}

func (*CanvasRenderingContext2d) MoveTo(x float32, y float32) {
	macro.Rewrite("$_.moveTo($1, $2)", x, y)
}

func (*CanvasRenderingContext2d) QuadraticCurveTo(cpx float32, cpy float32, x float32, y float32) {
	macro.Rewrite("$_.quadraticCurveTo($1, $2, $3, $4)", cpx, cpy, x, y)
}

func (*CanvasRenderingContext2d) Rect(x float32, y float32, w float32, h float32) {
	macro.Rewrite("$_.rect($1, $2, $3, $4)", x, y, w, h)
}

func (*CanvasRenderingContext2d) Canvas() (canvas *html.HTMLCanvasElement) {
	macro.Rewrite("$_.canvas")
	return canvas
}

func (*CanvasRenderingContext2d) FillStyle() (fillStyle interface{}) {
	macro.Rewrite("$_.fillStyle")
	return fillStyle
}

func (*CanvasRenderingContext2d) SetFillStyle(fillStyle interface{}) {
	macro.Rewrite("$_.fillStyle = $1", fillStyle)
}

func (*CanvasRenderingContext2d) Font() (font string) {
	macro.Rewrite("$_.font")
	return font
}

func (*CanvasRenderingContext2d) SetFont(font string) {
	macro.Rewrite("$_.font = $1", font)
}

func (*CanvasRenderingContext2d) GlobalAlpha() (globalAlpha float32) {
	macro.Rewrite("$_.globalAlpha")
	return globalAlpha
}

func (*CanvasRenderingContext2d) SetGlobalAlpha(globalAlpha float32) {
	macro.Rewrite("$_.globalAlpha = $1", globalAlpha)
}

func (*CanvasRenderingContext2d) GlobalCompositeOperation() (globalCompositeOperation string) {
	macro.Rewrite("$_.globalCompositeOperation")
	return globalCompositeOperation
}

func (*CanvasRenderingContext2d) SetGlobalCompositeOperation(globalCompositeOperation string) {
	macro.Rewrite("$_.globalCompositeOperation = $1", globalCompositeOperation)
}

func (*CanvasRenderingContext2d) ImageSmoothingEnabled() (imageSmoothingEnabled bool) {
	macro.Rewrite("$_.imageSmoothingEnabled")
	return imageSmoothingEnabled
}

func (*CanvasRenderingContext2d) SetImageSmoothingEnabled(imageSmoothingEnabled bool) {
	macro.Rewrite("$_.imageSmoothingEnabled = $1", imageSmoothingEnabled)
}

func (*CanvasRenderingContext2d) LineCap() (lineCap string) {
	macro.Rewrite("$_.lineCap")
	return lineCap
}

func (*CanvasRenderingContext2d) SetLineCap(lineCap string) {
	macro.Rewrite("$_.lineCap = $1", lineCap)
}

func (*CanvasRenderingContext2d) LineDashOffset() (lineDashOffset float32) {
	macro.Rewrite("$_.lineDashOffset")
	return lineDashOffset
}

func (*CanvasRenderingContext2d) SetLineDashOffset(lineDashOffset float32) {
	macro.Rewrite("$_.lineDashOffset = $1", lineDashOffset)
}

func (*CanvasRenderingContext2d) LineJoin() (lineJoin string) {
	macro.Rewrite("$_.lineJoin")
	return lineJoin
}

func (*CanvasRenderingContext2d) SetLineJoin(lineJoin string) {
	macro.Rewrite("$_.lineJoin = $1", lineJoin)
}

func (*CanvasRenderingContext2d) LineWidth() (lineWidth float32) {
	macro.Rewrite("$_.lineWidth")
	return lineWidth
}

func (*CanvasRenderingContext2d) SetLineWidth(lineWidth float32) {
	macro.Rewrite("$_.lineWidth = $1", lineWidth)
}

func (*CanvasRenderingContext2d) MiterLimit() (miterLimit float32) {
	macro.Rewrite("$_.miterLimit")
	return miterLimit
}

func (*CanvasRenderingContext2d) SetMiterLimit(miterLimit float32) {
	macro.Rewrite("$_.miterLimit = $1", miterLimit)
}

func (*CanvasRenderingContext2d) MsFillRule() (msFillRule *CanvasFillRule) {
	macro.Rewrite("$_.msFillRule")
	return msFillRule
}

func (*CanvasRenderingContext2d) SetMsFillRule(msFillRule *CanvasFillRule) {
	macro.Rewrite("$_.msFillRule = $1", msFillRule)
}

func (*CanvasRenderingContext2d) ShadowBlur() (shadowBlur float32) {
	macro.Rewrite("$_.shadowBlur")
	return shadowBlur
}

func (*CanvasRenderingContext2d) SetShadowBlur(shadowBlur float32) {
	macro.Rewrite("$_.shadowBlur = $1", shadowBlur)
}

func (*CanvasRenderingContext2d) ShadowColor() (shadowColor string) {
	macro.Rewrite("$_.shadowColor")
	return shadowColor
}

func (*CanvasRenderingContext2d) SetShadowColor(shadowColor string) {
	macro.Rewrite("$_.shadowColor = $1", shadowColor)
}

func (*CanvasRenderingContext2d) ShadowOffsetX() (shadowOffsetX float32) {
	macro.Rewrite("$_.shadowOffsetX")
	return shadowOffsetX
}

func (*CanvasRenderingContext2d) SetShadowOffsetX(shadowOffsetX float32) {
	macro.Rewrite("$_.shadowOffsetX = $1", shadowOffsetX)
}

func (*CanvasRenderingContext2d) ShadowOffsetY() (shadowOffsetY float32) {
	macro.Rewrite("$_.shadowOffsetY")
	return shadowOffsetY
}

func (*CanvasRenderingContext2d) SetShadowOffsetY(shadowOffsetY float32) {
	macro.Rewrite("$_.shadowOffsetY = $1", shadowOffsetY)
}

func (*CanvasRenderingContext2d) StrokeStyle() (strokeStyle interface{}) {
	macro.Rewrite("$_.strokeStyle")
	return strokeStyle
}

func (*CanvasRenderingContext2d) SetStrokeStyle(strokeStyle interface{}) {
	macro.Rewrite("$_.strokeStyle = $1", strokeStyle)
}

func (*CanvasRenderingContext2d) TextAlign() (textAlign string) {
	macro.Rewrite("$_.textAlign")
	return textAlign
}

func (*CanvasRenderingContext2d) SetTextAlign(textAlign string) {
	macro.Rewrite("$_.textAlign = $1", textAlign)
}

func (*CanvasRenderingContext2d) TextBaseline() (textBaseline string) {
	macro.Rewrite("$_.textBaseline")
	return textBaseline
}

func (*CanvasRenderingContext2d) SetTextBaseline(textBaseline string) {
	macro.Rewrite("$_.textBaseline = $1", textBaseline)
}
