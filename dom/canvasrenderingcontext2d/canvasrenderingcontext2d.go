package canvasrenderingcontext2d

import (
	"github.com/matthewmueller/joy/dom/canvas"
	"github.com/matthewmueller/joy/dom/htmlcanvaselement"
	"github.com/matthewmueller/joy/dom/path2d"
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/macro"
)

// CanvasRenderingContext2d struct
// js:"CanvasRenderingContext2d,omit"
type CanvasRenderingContext2d struct {
}

// BeginPath fn
// js:"beginPath"
func (*CanvasRenderingContext2d) BeginPath() {
	macro.Rewrite("$_.beginPath()")
}

// ClearRect fn
// js:"clearRect"
func (*CanvasRenderingContext2d) ClearRect(x float32, y float32, w float32, h float32) {
	macro.Rewrite("$_.clearRect($1, $2, $3, $4)", x, y, w, h)
}

// Clip fn
// js:"clip"
func (*CanvasRenderingContext2d) Clip(fillRule *canvas.CanvasFillRule) {
	macro.Rewrite("$_.clip($1)", fillRule)
}

// CreateImageData fn
// js:"createImageData"
func (*CanvasRenderingContext2d) CreateImageData(imageDataOrSw interface{}, sh *float32) (c *canvas.ImageData) {
	macro.Rewrite("$_.createImageData($1, $2)", imageDataOrSw, sh)
	return c
}

// CreateLinearGradient fn
// js:"createLinearGradient"
func (*CanvasRenderingContext2d) CreateLinearGradient(x0 float32, y0 float32, x1 float32, y1 float32) (c *canvas.CanvasGradient) {
	macro.Rewrite("$_.createLinearGradient($1, $2, $3, $4)", x0, y0, x1, y1)
	return c
}

// CreatePattern fn
// js:"createPattern"
func (*CanvasRenderingContext2d) CreatePattern(image interface{}, repetition string) (c *canvas.CanvasPattern) {
	macro.Rewrite("$_.createPattern($1, $2)", image, repetition)
	return c
}

// CreateRadialGradient fn
// js:"createRadialGradient"
func (*CanvasRenderingContext2d) CreateRadialGradient(x0 float32, y0 float32, r0 float32, x1 float32, y1 float32, r1 float32) (c *canvas.CanvasGradient) {
	macro.Rewrite("$_.createRadialGradient($1, $2, $3, $4, $5, $6)", x0, y0, r0, x1, y1, r1)
	return c
}

// DrawFocusIfNeeded fn
// js:"drawFocusIfNeeded"
func (*CanvasRenderingContext2d) DrawFocusIfNeeded(element window.Element) {
	macro.Rewrite("$_.drawFocusIfNeeded($1)", element)
}

// DrawImage fn
// js:"drawImage"
func (*CanvasRenderingContext2d) DrawImage(image interface{}, offsetX float32, offsetY float32, width *float32, height *float32, canvasOffsetX *float32, canvasOffsetY *float32, canvasImageWidth *float32, canvasImageHeight *float32) {
	macro.Rewrite("$_.drawImage($1, $2, $3, $4, $5, $6, $7, $8, $9)", image, offsetX, offsetY, width, height, canvasOffsetX, canvasOffsetY, canvasImageWidth, canvasImageHeight)
}

// Fill fn
// js:"fill"
func (*CanvasRenderingContext2d) Fill(fillRule *canvas.CanvasFillRule) {
	macro.Rewrite("$_.fill($1)", fillRule)
}

// FillRect fn
// js:"fillRect"
func (*CanvasRenderingContext2d) FillRect(x float32, y float32, w float32, h float32) {
	macro.Rewrite("$_.fillRect($1, $2, $3, $4)", x, y, w, h)
}

// FillText fn
// js:"fillText"
func (*CanvasRenderingContext2d) FillText(text string, x float32, y float32, maxWidth *float32) {
	macro.Rewrite("$_.fillText($1, $2, $3, $4)", text, x, y, maxWidth)
}

// GetImageData fn
// js:"getImageData"
func (*CanvasRenderingContext2d) GetImageData(sx float32, sy float32, sw float32, sh float32) (c *canvas.ImageData) {
	macro.Rewrite("$_.getImageData($1, $2, $3, $4)", sx, sy, sw, sh)
	return c
}

// GetLineDash fn
// js:"getLineDash"
func (*CanvasRenderingContext2d) GetLineDash() (f []float32) {
	macro.Rewrite("$_.getLineDash()")
	return f
}

// IsPointInPath fn
// js:"isPointInPath"
func (*CanvasRenderingContext2d) IsPointInPath(x float32, y float32, fillRule *canvas.CanvasFillRule) (b bool) {
	macro.Rewrite("$_.isPointInPath($1, $2, $3)", x, y, fillRule)
	return b
}

// MeasureText fn
// js:"measureText"
func (*CanvasRenderingContext2d) MeasureText(text string) (c *canvas.TextMetrics) {
	macro.Rewrite("$_.measureText($1)", text)
	return c
}

// PutImageData fn
// js:"putImageData"
func (*CanvasRenderingContext2d) PutImageData(imagedata *canvas.ImageData, dx float32, dy float32, dirtyX *float32, dirtyY *float32, dirtyWidth *float32, dirtyHeight *float32) {
	macro.Rewrite("$_.putImageData($1, $2, $3, $4, $5, $6, $7)", imagedata, dx, dy, dirtyX, dirtyY, dirtyWidth, dirtyHeight)
}

// Restore fn
// js:"restore"
func (*CanvasRenderingContext2d) Restore() {
	macro.Rewrite("$_.restore()")
}

// Rotate fn
// js:"rotate"
func (*CanvasRenderingContext2d) Rotate(angle float32) {
	macro.Rewrite("$_.rotate($1)", angle)
}

// Save fn
// js:"save"
func (*CanvasRenderingContext2d) Save() {
	macro.Rewrite("$_.save()")
}

// Scale fn
// js:"scale"
func (*CanvasRenderingContext2d) Scale(x float32, y float32) {
	macro.Rewrite("$_.scale($1, $2)", x, y)
}

// SetLineDash fn
// js:"setLineDash"
func (*CanvasRenderingContext2d) SetLineDash(segments []float32) {
	macro.Rewrite("$_.setLineDash($1)", segments)
}

// SetTransform fn
// js:"setTransform"
func (*CanvasRenderingContext2d) SetTransform(m11 float32, m12 float32, m21 float32, m22 float32, dx float32, dy float32) {
	macro.Rewrite("$_.setTransform($1, $2, $3, $4, $5, $6)", m11, m12, m21, m22, dx, dy)
}

// Stroke fn
// js:"stroke"
func (*CanvasRenderingContext2d) Stroke(path *path2d.Path2d) {
	macro.Rewrite("$_.stroke($1)", path)
}

// StrokeRect fn
// js:"strokeRect"
func (*CanvasRenderingContext2d) StrokeRect(x float32, y float32, w float32, h float32) {
	macro.Rewrite("$_.strokeRect($1, $2, $3, $4)", x, y, w, h)
}

// StrokeText fn
// js:"strokeText"
func (*CanvasRenderingContext2d) StrokeText(text string, x float32, y float32, maxWidth *float32) {
	macro.Rewrite("$_.strokeText($1, $2, $3, $4)", text, x, y, maxWidth)
}

// Transform fn
// js:"transform"
func (*CanvasRenderingContext2d) Transform(m11 float32, m12 float32, m21 float32, m22 float32, dx float32, dy float32) {
	macro.Rewrite("$_.transform($1, $2, $3, $4, $5, $6)", m11, m12, m21, m22, dx, dy)
}

// Translate fn
// js:"translate"
func (*CanvasRenderingContext2d) Translate(x float32, y float32) {
	macro.Rewrite("$_.translate($1, $2)", x, y)
}

// Arc fn
// js:"arc"
func (*CanvasRenderingContext2d) Arc(x float32, y float32, radius float32, startAngle float32, endAngle float32, anticlockwise *bool) {
	macro.Rewrite("$_.arc($1, $2, $3, $4, $5, $6)", x, y, radius, startAngle, endAngle, anticlockwise)
}

// ArcTo fn
// js:"arcTo"
func (*CanvasRenderingContext2d) ArcTo(x1 float32, y1 float32, x2 float32, y2 float32, radius float32) {
	macro.Rewrite("$_.arcTo($1, $2, $3, $4, $5)", x1, y1, x2, y2, radius)
}

// BezierCurveTo fn
// js:"bezierCurveTo"
func (*CanvasRenderingContext2d) BezierCurveTo(cp1x float32, cp1y float32, cp2x float32, cp2y float32, x float32, y float32) {
	macro.Rewrite("$_.bezierCurveTo($1, $2, $3, $4, $5, $6)", cp1x, cp1y, cp2x, cp2y, x, y)
}

// ClosePath fn
// js:"closePath"
func (*CanvasRenderingContext2d) ClosePath() {
	macro.Rewrite("$_.closePath()")
}

// Ellipse fn
// js:"ellipse"
func (*CanvasRenderingContext2d) Ellipse(x float32, y float32, radiusX float32, radiusY float32, rotation float32, startAngle float32, endAngle float32, anticlockwise *bool) {
	macro.Rewrite("$_.ellipse($1, $2, $3, $4, $5, $6, $7, $8)", x, y, radiusX, radiusY, rotation, startAngle, endAngle, anticlockwise)
}

// LineTo fn
// js:"lineTo"
func (*CanvasRenderingContext2d) LineTo(x float32, y float32) {
	macro.Rewrite("$_.lineTo($1, $2)", x, y)
}

// MoveTo fn
// js:"moveTo"
func (*CanvasRenderingContext2d) MoveTo(x float32, y float32) {
	macro.Rewrite("$_.moveTo($1, $2)", x, y)
}

// QuadraticCurveTo fn
// js:"quadraticCurveTo"
func (*CanvasRenderingContext2d) QuadraticCurveTo(cpx float32, cpy float32, x float32, y float32) {
	macro.Rewrite("$_.quadraticCurveTo($1, $2, $3, $4)", cpx, cpy, x, y)
}

// Rect fn
// js:"rect"
func (*CanvasRenderingContext2d) Rect(x float32, y float32, w float32, h float32) {
	macro.Rewrite("$_.rect($1, $2, $3, $4)", x, y, w, h)
}

// Canvas prop
// js:"canvas"
func (*CanvasRenderingContext2d) Canvas() (canvas *htmlcanvaselement.HTMLCanvasElement) {
	macro.Rewrite("$_.canvas")
	return canvas
}

// FillStyle prop
// js:"fillStyle"
func (*CanvasRenderingContext2d) FillStyle() (fillStyle interface{}) {
	macro.Rewrite("$_.fillStyle")
	return fillStyle
}

// SetFillStyle prop
// js:"fillStyle"
func (*CanvasRenderingContext2d) SetFillStyle(fillStyle interface{}) {
	macro.Rewrite("$_.fillStyle = $1", fillStyle)
}

// Font prop
// js:"font"
func (*CanvasRenderingContext2d) Font() (font string) {
	macro.Rewrite("$_.font")
	return font
}

// SetFont prop
// js:"font"
func (*CanvasRenderingContext2d) SetFont(font string) {
	macro.Rewrite("$_.font = $1", font)
}

// GlobalAlpha prop
// js:"globalAlpha"
func (*CanvasRenderingContext2d) GlobalAlpha() (globalAlpha float32) {
	macro.Rewrite("$_.globalAlpha")
	return globalAlpha
}

// SetGlobalAlpha prop
// js:"globalAlpha"
func (*CanvasRenderingContext2d) SetGlobalAlpha(globalAlpha float32) {
	macro.Rewrite("$_.globalAlpha = $1", globalAlpha)
}

// GlobalCompositeOperation prop
// js:"globalCompositeOperation"
func (*CanvasRenderingContext2d) GlobalCompositeOperation() (globalCompositeOperation string) {
	macro.Rewrite("$_.globalCompositeOperation")
	return globalCompositeOperation
}

// SetGlobalCompositeOperation prop
// js:"globalCompositeOperation"
func (*CanvasRenderingContext2d) SetGlobalCompositeOperation(globalCompositeOperation string) {
	macro.Rewrite("$_.globalCompositeOperation = $1", globalCompositeOperation)
}

// ImageSmoothingEnabled prop
// js:"imageSmoothingEnabled"
func (*CanvasRenderingContext2d) ImageSmoothingEnabled() (imageSmoothingEnabled bool) {
	macro.Rewrite("$_.imageSmoothingEnabled")
	return imageSmoothingEnabled
}

// SetImageSmoothingEnabled prop
// js:"imageSmoothingEnabled"
func (*CanvasRenderingContext2d) SetImageSmoothingEnabled(imageSmoothingEnabled bool) {
	macro.Rewrite("$_.imageSmoothingEnabled = $1", imageSmoothingEnabled)
}

// LineCap prop
// js:"lineCap"
func (*CanvasRenderingContext2d) LineCap() (lineCap string) {
	macro.Rewrite("$_.lineCap")
	return lineCap
}

// SetLineCap prop
// js:"lineCap"
func (*CanvasRenderingContext2d) SetLineCap(lineCap string) {
	macro.Rewrite("$_.lineCap = $1", lineCap)
}

// LineDashOffset prop
// js:"lineDashOffset"
func (*CanvasRenderingContext2d) LineDashOffset() (lineDashOffset float32) {
	macro.Rewrite("$_.lineDashOffset")
	return lineDashOffset
}

// SetLineDashOffset prop
// js:"lineDashOffset"
func (*CanvasRenderingContext2d) SetLineDashOffset(lineDashOffset float32) {
	macro.Rewrite("$_.lineDashOffset = $1", lineDashOffset)
}

// LineJoin prop
// js:"lineJoin"
func (*CanvasRenderingContext2d) LineJoin() (lineJoin string) {
	macro.Rewrite("$_.lineJoin")
	return lineJoin
}

// SetLineJoin prop
// js:"lineJoin"
func (*CanvasRenderingContext2d) SetLineJoin(lineJoin string) {
	macro.Rewrite("$_.lineJoin = $1", lineJoin)
}

// LineWidth prop
// js:"lineWidth"
func (*CanvasRenderingContext2d) LineWidth() (lineWidth float32) {
	macro.Rewrite("$_.lineWidth")
	return lineWidth
}

// SetLineWidth prop
// js:"lineWidth"
func (*CanvasRenderingContext2d) SetLineWidth(lineWidth float32) {
	macro.Rewrite("$_.lineWidth = $1", lineWidth)
}

// MiterLimit prop
// js:"miterLimit"
func (*CanvasRenderingContext2d) MiterLimit() (miterLimit float32) {
	macro.Rewrite("$_.miterLimit")
	return miterLimit
}

// SetMiterLimit prop
// js:"miterLimit"
func (*CanvasRenderingContext2d) SetMiterLimit(miterLimit float32) {
	macro.Rewrite("$_.miterLimit = $1", miterLimit)
}

// MsFillRule prop
// js:"msFillRule"
func (*CanvasRenderingContext2d) MsFillRule() (msFillRule *canvas.CanvasFillRule) {
	macro.Rewrite("$_.msFillRule")
	return msFillRule
}

// SetMsFillRule prop
// js:"msFillRule"
func (*CanvasRenderingContext2d) SetMsFillRule(msFillRule *canvas.CanvasFillRule) {
	macro.Rewrite("$_.msFillRule = $1", msFillRule)
}

// ShadowBlur prop
// js:"shadowBlur"
func (*CanvasRenderingContext2d) ShadowBlur() (shadowBlur float32) {
	macro.Rewrite("$_.shadowBlur")
	return shadowBlur
}

// SetShadowBlur prop
// js:"shadowBlur"
func (*CanvasRenderingContext2d) SetShadowBlur(shadowBlur float32) {
	macro.Rewrite("$_.shadowBlur = $1", shadowBlur)
}

// ShadowColor prop
// js:"shadowColor"
func (*CanvasRenderingContext2d) ShadowColor() (shadowColor string) {
	macro.Rewrite("$_.shadowColor")
	return shadowColor
}

// SetShadowColor prop
// js:"shadowColor"
func (*CanvasRenderingContext2d) SetShadowColor(shadowColor string) {
	macro.Rewrite("$_.shadowColor = $1", shadowColor)
}

// ShadowOffsetX prop
// js:"shadowOffsetX"
func (*CanvasRenderingContext2d) ShadowOffsetX() (shadowOffsetX float32) {
	macro.Rewrite("$_.shadowOffsetX")
	return shadowOffsetX
}

// SetShadowOffsetX prop
// js:"shadowOffsetX"
func (*CanvasRenderingContext2d) SetShadowOffsetX(shadowOffsetX float32) {
	macro.Rewrite("$_.shadowOffsetX = $1", shadowOffsetX)
}

// ShadowOffsetY prop
// js:"shadowOffsetY"
func (*CanvasRenderingContext2d) ShadowOffsetY() (shadowOffsetY float32) {
	macro.Rewrite("$_.shadowOffsetY")
	return shadowOffsetY
}

// SetShadowOffsetY prop
// js:"shadowOffsetY"
func (*CanvasRenderingContext2d) SetShadowOffsetY(shadowOffsetY float32) {
	macro.Rewrite("$_.shadowOffsetY = $1", shadowOffsetY)
}

// StrokeStyle prop
// js:"strokeStyle"
func (*CanvasRenderingContext2d) StrokeStyle() (strokeStyle interface{}) {
	macro.Rewrite("$_.strokeStyle")
	return strokeStyle
}

// SetStrokeStyle prop
// js:"strokeStyle"
func (*CanvasRenderingContext2d) SetStrokeStyle(strokeStyle interface{}) {
	macro.Rewrite("$_.strokeStyle = $1", strokeStyle)
}

// TextAlign prop
// js:"textAlign"
func (*CanvasRenderingContext2d) TextAlign() (textAlign string) {
	macro.Rewrite("$_.textAlign")
	return textAlign
}

// SetTextAlign prop
// js:"textAlign"
func (*CanvasRenderingContext2d) SetTextAlign(textAlign string) {
	macro.Rewrite("$_.textAlign = $1", textAlign)
}

// TextBaseline prop
// js:"textBaseline"
func (*CanvasRenderingContext2d) TextBaseline() (textBaseline string) {
	macro.Rewrite("$_.textBaseline")
	return textBaseline
}

// SetTextBaseline prop
// js:"textBaseline"
func (*CanvasRenderingContext2d) SetTextBaseline(textBaseline string) {
	macro.Rewrite("$_.textBaseline = $1", textBaseline)
}
