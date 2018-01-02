package webgl

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/canvas"
	"github.com/matthewmueller/joy/dom/html"
)

type WebGLRenderingContext struct {
}

func (*WebGLRenderingContext) ActiveTexture(texture uint) {
	macro.Rewrite("$_.activeTexture($1)", texture)
}

func (*WebGLRenderingContext) AttachShader(program *WebGLProgram, shader *WebGLShader) {
	macro.Rewrite("$_.attachShader($1, $2)", program, shader)
}

func (*WebGLRenderingContext) BindAttribLocation(program *WebGLProgram, index uint, name string) {
	macro.Rewrite("$_.bindAttribLocation($1, $2, $3)", program, index, name)
}

func (*WebGLRenderingContext) BindBuffer(target uint, buffer *WebGLBuffer) {
	macro.Rewrite("$_.bindBuffer($1, $2)", target, buffer)
}

func (*WebGLRenderingContext) BindFramebuffer(target uint, framebuffer *WebGLFramebuffer) {
	macro.Rewrite("$_.bindFramebuffer($1, $2)", target, framebuffer)
}

func (*WebGLRenderingContext) BindRenderbuffer(target uint, renderbuffer *WebGLRenderbuffer) {
	macro.Rewrite("$_.bindRenderbuffer($1, $2)", target, renderbuffer)
}

func (*WebGLRenderingContext) BindTexture(target uint, texture *WebGLTexture) {
	macro.Rewrite("$_.bindTexture($1, $2)", target, texture)
}

func (*WebGLRenderingContext) BlendColor(red float32, green float32, blue float32, alpha float32) {
	macro.Rewrite("$_.blendColor($1, $2, $3, $4)", red, green, blue, alpha)
}

func (*WebGLRenderingContext) BlendEquation(mode uint) {
	macro.Rewrite("$_.blendEquation($1)", mode)
}

func (*WebGLRenderingContext) BlendEquationSeparate(modeRGB uint, modeAlpha uint) {
	macro.Rewrite("$_.blendEquationSeparate($1, $2)", modeRGB, modeAlpha)
}

func (*WebGLRenderingContext) BlendFunc(sfactor uint, dfactor uint) {
	macro.Rewrite("$_.blendFunc($1, $2)", sfactor, dfactor)
}

func (*WebGLRenderingContext) BlendFuncSeparate(srcRGB uint, dstRGB uint, srcAlpha uint, dstAlpha uint) {
	macro.Rewrite("$_.blendFuncSeparate($1, $2, $3, $4)", srcRGB, dstRGB, srcAlpha, dstAlpha)
}

func (*WebGLRenderingContext) BufferData(target uint, size interface{}, usage uint) {
	macro.Rewrite("$_.bufferData($1, $2, $3)", target, size, usage)
}

func (*WebGLRenderingContext) BufferSubData(target uint, offset int64, data interface{}) {
	macro.Rewrite("$_.bufferSubData($1, $2, $3)", target, offset, data)
}

func (*WebGLRenderingContext) CheckFramebufferStatus(target uint) (u uint) {
	macro.Rewrite("$_.checkFramebufferStatus($1)", target)
	return u
}

func (*WebGLRenderingContext) Clear(mask uint) {
	macro.Rewrite("$_.clear($1)", mask)
}

func (*WebGLRenderingContext) ClearColor(red float32, green float32, blue float32, alpha float32) {
	macro.Rewrite("$_.clearColor($1, $2, $3, $4)", red, green, blue, alpha)
}

func (*WebGLRenderingContext) ClearDepth(depth float32) {
	macro.Rewrite("$_.clearDepth($1)", depth)
}

func (*WebGLRenderingContext) ClearStencil(s int) {
	macro.Rewrite("$_.clearStencil($1)", s)
}

func (*WebGLRenderingContext) ColorMask(red bool, green bool, blue bool, alpha bool) {
	macro.Rewrite("$_.colorMask($1, $2, $3, $4)", red, green, blue, alpha)
}

func (*WebGLRenderingContext) CompileShader(shader *WebGLShader) {
	macro.Rewrite("$_.compileShader($1)", shader)
}

func (*WebGLRenderingContext) CompressedTexImage2d(target uint, level int, internalformat uint, width int, height int, border int, data []byte) {
	macro.Rewrite("$_.compressedTexImage2D($1, $2, $3, $4, $5, $6, $7)", target, level, internalformat, width, height, border, data)
}

func (*WebGLRenderingContext) CompressedTexSubImage2d(target uint, level int, xoffset int, yoffset int, width int, height int, format uint, data []byte) {
	macro.Rewrite("$_.compressedTexSubImage2D($1, $2, $3, $4, $5, $6, $7, $8)", target, level, xoffset, yoffset, width, height, format, data)
}

func (*WebGLRenderingContext) CopyTexImage2d(target uint, level int, internalformat uint, x int, y int, width int, height int, border int) {
	macro.Rewrite("$_.copyTexImage2D($1, $2, $3, $4, $5, $6, $7, $8)", target, level, internalformat, x, y, width, height, border)
}

func (*WebGLRenderingContext) CopyTexSubImage2d(target uint, level int, xoffset int, yoffset int, x int, y int, width int, height int) {
	macro.Rewrite("$_.copyTexSubImage2D($1, $2, $3, $4, $5, $6, $7, $8)", target, level, xoffset, yoffset, x, y, width, height)
}

func (*WebGLRenderingContext) CreateBuffer() (w *WebGLBuffer) {
	macro.Rewrite("$_.createBuffer()")
	return w
}

func (*WebGLRenderingContext) CreateFramebuffer() (w *WebGLFramebuffer) {
	macro.Rewrite("$_.createFramebuffer()")
	return w
}

func (*WebGLRenderingContext) CreateProgram() (w *WebGLProgram) {
	macro.Rewrite("$_.createProgram()")
	return w
}

func (*WebGLRenderingContext) CreateRenderbuffer() (w *WebGLRenderbuffer) {
	macro.Rewrite("$_.createRenderbuffer()")
	return w
}

func (*WebGLRenderingContext) CreateShader(kind uint) (w *WebGLShader) {
	macro.Rewrite("$_.createShader($1)", kind)
	return w
}

func (*WebGLRenderingContext) CreateTexture() (w *WebGLTexture) {
	macro.Rewrite("$_.createTexture()")
	return w
}

func (*WebGLRenderingContext) CullFace(mode uint) {
	macro.Rewrite("$_.cullFace($1)", mode)
}

func (*WebGLRenderingContext) DeleteBuffer(buffer *WebGLBuffer) {
	macro.Rewrite("$_.deleteBuffer($1)", buffer)
}

func (*WebGLRenderingContext) DeleteFramebuffer(framebuffer *WebGLFramebuffer) {
	macro.Rewrite("$_.deleteFramebuffer($1)", framebuffer)
}

func (*WebGLRenderingContext) DeleteProgram(program *WebGLProgram) {
	macro.Rewrite("$_.deleteProgram($1)", program)
}

func (*WebGLRenderingContext) DeleteRenderbuffer(renderbuffer *WebGLRenderbuffer) {
	macro.Rewrite("$_.deleteRenderbuffer($1)", renderbuffer)
}

func (*WebGLRenderingContext) DeleteShader(shader *WebGLShader) {
	macro.Rewrite("$_.deleteShader($1)", shader)
}

func (*WebGLRenderingContext) DeleteTexture(texture *WebGLTexture) {
	macro.Rewrite("$_.deleteTexture($1)", texture)
}

func (*WebGLRenderingContext) DepthFunc(fn uint) {
	macro.Rewrite("$_.depthFunc($1)", fn)
}

func (*WebGLRenderingContext) DepthMask(flag bool) {
	macro.Rewrite("$_.depthMask($1)", flag)
}

func (*WebGLRenderingContext) DepthRange(zNear float32, zFar float32) {
	macro.Rewrite("$_.depthRange($1, $2)", zNear, zFar)
}

func (*WebGLRenderingContext) DetachShader(program *WebGLProgram, shader *WebGLShader) {
	macro.Rewrite("$_.detachShader($1, $2)", program, shader)
}

func (*WebGLRenderingContext) Disable(capacity uint) {
	macro.Rewrite("$_.disable($1)", capacity)
}

func (*WebGLRenderingContext) DisableVertexAttribArray(index uint) {
	macro.Rewrite("$_.disableVertexAttribArray($1)", index)
}

func (*WebGLRenderingContext) DrawArrays(mode uint, first int, count int) {
	macro.Rewrite("$_.drawArrays($1, $2, $3)", mode, first, count)
}

func (*WebGLRenderingContext) DrawElements(mode uint, count int, kind uint, offset int64) {
	macro.Rewrite("$_.drawElements($1, $2, $3, $4)", mode, count, kind, offset)
}

func (*WebGLRenderingContext) Enable(capacity uint) {
	macro.Rewrite("$_.enable($1)", capacity)
}

func (*WebGLRenderingContext) EnableVertexAttribArray(index uint) {
	macro.Rewrite("$_.enableVertexAttribArray($1)", index)
}

func (*WebGLRenderingContext) Finish() {
	macro.Rewrite("$_.finish()")
}

func (*WebGLRenderingContext) Flush() {
	macro.Rewrite("$_.flush()")
}

func (*WebGLRenderingContext) FramebufferRenderbuffer(target uint, attachment uint, renderbuffertarget uint, renderbuffer *WebGLRenderbuffer) {
	macro.Rewrite("$_.framebufferRenderbuffer($1, $2, $3, $4)", target, attachment, renderbuffertarget, renderbuffer)
}

func (*WebGLRenderingContext) FramebufferTexture2d(target uint, attachment uint, textarget uint, texture *WebGLTexture, level int) {
	macro.Rewrite("$_.framebufferTexture2D($1, $2, $3, $4, $5)", target, attachment, textarget, texture, level)
}

func (*WebGLRenderingContext) FrontFace(mode uint) {
	macro.Rewrite("$_.frontFace($1)", mode)
}

func (*WebGLRenderingContext) GenerateMipmap(target uint) {
	macro.Rewrite("$_.generateMipmap($1)", target)
}

func (*WebGLRenderingContext) GetActiveAttrib(program *WebGLProgram, index uint) (w *WebGLActiveInfo) {
	macro.Rewrite("$_.getActiveAttrib($1, $2)", program, index)
	return w
}

func (*WebGLRenderingContext) GetActiveUniform(program *WebGLProgram, index uint) (w *WebGLActiveInfo) {
	macro.Rewrite("$_.getActiveUniform($1, $2)", program, index)
	return w
}

func (*WebGLRenderingContext) GetAttachedShaders(program *WebGLProgram) (w []*WebGLShader) {
	macro.Rewrite("$_.getAttachedShaders($1)", program)
	return w
}

func (*WebGLRenderingContext) GetAttribLocation(program *WebGLProgram, name string) (i int) {
	macro.Rewrite("$_.getAttribLocation($1, $2)", program, name)
	return i
}

func (*WebGLRenderingContext) GetBufferParameter(target uint, pname uint) (i interface{}) {
	macro.Rewrite("$_.getBufferParameter($1, $2)", target, pname)
	return i
}

func (*WebGLRenderingContext) GetContextAttributes() (w *WebGLContextAttributes) {
	macro.Rewrite("$_.getContextAttributes()")
	return w
}

func (*WebGLRenderingContext) GetError() (u uint) {
	macro.Rewrite("$_.getError()")
	return u
}

func (*WebGLRenderingContext) GetExtension(name string) (i interface{}) {
	macro.Rewrite("$_.getExtension($1)", name)
	return i
}

func (*WebGLRenderingContext) GetFramebufferAttachmentParameter(target uint, attachment uint, pname uint) (i interface{}) {
	macro.Rewrite("$_.getFramebufferAttachmentParameter($1, $2, $3)", target, attachment, pname)
	return i
}

func (*WebGLRenderingContext) GetParameter(pname uint) (i interface{}) {
	macro.Rewrite("$_.getParameter($1)", pname)
	return i
}

func (*WebGLRenderingContext) GetProgramInfoLog(program *WebGLProgram) (s string) {
	macro.Rewrite("$_.getProgramInfoLog($1)", program)
	return s
}

func (*WebGLRenderingContext) GetProgramParameter(program *WebGLProgram, pname uint) (i interface{}) {
	macro.Rewrite("$_.getProgramParameter($1, $2)", program, pname)
	return i
}

func (*WebGLRenderingContext) GetRenderbufferParameter(target uint, pname uint) (i interface{}) {
	macro.Rewrite("$_.getRenderbufferParameter($1, $2)", target, pname)
	return i
}

func (*WebGLRenderingContext) GetShaderInfoLog(shader *WebGLShader) (s string) {
	macro.Rewrite("$_.getShaderInfoLog($1)", shader)
	return s
}

func (*WebGLRenderingContext) GetShaderParameter(shader *WebGLShader, pname uint) (i interface{}) {
	macro.Rewrite("$_.getShaderParameter($1, $2)", shader, pname)
	return i
}

func (*WebGLRenderingContext) GetShaderPrecisionFormat(shadertype uint, precisiontype uint) (w *WebGLShaderPrecisionFormat) {
	macro.Rewrite("$_.getShaderPrecisionFormat($1, $2)", shadertype, precisiontype)
	return w
}

func (*WebGLRenderingContext) GetShaderSource(shader *WebGLShader) (s string) {
	macro.Rewrite("$_.getShaderSource($1)", shader)
	return s
}

func (*WebGLRenderingContext) GetSupportedExtensions() (s []string) {
	macro.Rewrite("$_.getSupportedExtensions()")
	return s
}

func (*WebGLRenderingContext) GetTexParameter(target uint, pname uint) (i interface{}) {
	macro.Rewrite("$_.getTexParameter($1, $2)", target, pname)
	return i
}

func (*WebGLRenderingContext) GetUniform(program *WebGLProgram, location *WebGLUniformLocation) (i interface{}) {
	macro.Rewrite("$_.getUniform($1, $2)", program, location)
	return i
}

func (*WebGLRenderingContext) GetUniformLocation(program *WebGLProgram, name string) (w *WebGLUniformLocation) {
	macro.Rewrite("$_.getUniformLocation($1, $2)", program, name)
	return w
}

func (*WebGLRenderingContext) GetVertexAttrib(index uint, pname uint) (i interface{}) {
	macro.Rewrite("$_.getVertexAttrib($1, $2)", index, pname)
	return i
}

func (*WebGLRenderingContext) GetVertexAttribOffset(index uint, pname uint) (i int64) {
	macro.Rewrite("$_.getVertexAttribOffset($1, $2)", index, pname)
	return i
}

func (*WebGLRenderingContext) Hint(target uint, mode uint) {
	macro.Rewrite("$_.hint($1, $2)", target, mode)
}

func (*WebGLRenderingContext) IsBuffer(buffer *WebGLBuffer) (b bool) {
	macro.Rewrite("$_.isBuffer($1)", buffer)
	return b
}

func (*WebGLRenderingContext) IsContextLost() (b bool) {
	macro.Rewrite("$_.isContextLost()")
	return b
}

func (*WebGLRenderingContext) IsEnabled(capacity uint) (b bool) {
	macro.Rewrite("$_.isEnabled($1)", capacity)
	return b
}

func (*WebGLRenderingContext) IsFramebuffer(framebuffer *WebGLFramebuffer) (b bool) {
	macro.Rewrite("$_.isFramebuffer($1)", framebuffer)
	return b
}

func (*WebGLRenderingContext) IsProgram(program *WebGLProgram) (b bool) {
	macro.Rewrite("$_.isProgram($1)", program)
	return b
}

func (*WebGLRenderingContext) IsRenderbuffer(renderbuffer *WebGLRenderbuffer) (b bool) {
	macro.Rewrite("$_.isRenderbuffer($1)", renderbuffer)
	return b
}

func (*WebGLRenderingContext) IsShader(shader *WebGLShader) (b bool) {
	macro.Rewrite("$_.isShader($1)", shader)
	return b
}

func (*WebGLRenderingContext) IsTexture(texture *WebGLTexture) (b bool) {
	macro.Rewrite("$_.isTexture($1)", texture)
	return b
}

func (*WebGLRenderingContext) LineWidth(width float32) {
	macro.Rewrite("$_.lineWidth($1)", width)
}

func (*WebGLRenderingContext) LinkProgram(program *WebGLProgram) {
	macro.Rewrite("$_.linkProgram($1)", program)
}

func (*WebGLRenderingContext) PixelStorei(pname uint, param int) {
	macro.Rewrite("$_.pixelStorei($1, $2)", pname, param)
}

func (*WebGLRenderingContext) PolygonOffset(factor float32, units float32) {
	macro.Rewrite("$_.polygonOffset($1, $2)", factor, units)
}

func (*WebGLRenderingContext) ReadPixels(x int, y int, width int, height int, format uint, kind uint, pixels []byte) {
	macro.Rewrite("$_.readPixels($1, $2, $3, $4, $5, $6, $7)", x, y, width, height, format, kind, pixels)
}

func (*WebGLRenderingContext) RenderbufferStorage(target uint, internalformat uint, width int, height int) {
	macro.Rewrite("$_.renderbufferStorage($1, $2, $3, $4)", target, internalformat, width, height)
}

func (*WebGLRenderingContext) SampleCoverage(value float32, invert bool) {
	macro.Rewrite("$_.sampleCoverage($1, $2)", value, invert)
}

func (*WebGLRenderingContext) Scissor(x int, y int, width int, height int) {
	macro.Rewrite("$_.scissor($1, $2, $3, $4)", x, y, width, height)
}

func (*WebGLRenderingContext) ShaderSource(shader *WebGLShader, source string) {
	macro.Rewrite("$_.shaderSource($1, $2)", shader, source)
}

func (*WebGLRenderingContext) StencilFunc(fn uint, ref int, mask uint) {
	macro.Rewrite("$_.stencilFunc($1, $2, $3)", fn, ref, mask)
}

func (*WebGLRenderingContext) StencilFuncSeparate(face uint, fn uint, ref int, mask uint) {
	macro.Rewrite("$_.stencilFuncSeparate($1, $2, $3, $4)", face, fn, ref, mask)
}

func (*WebGLRenderingContext) StencilMask(mask uint) {
	macro.Rewrite("$_.stencilMask($1)", mask)
}

func (*WebGLRenderingContext) StencilMaskSeparate(face uint, mask uint) {
	macro.Rewrite("$_.stencilMaskSeparate($1, $2)", face, mask)
}

func (*WebGLRenderingContext) StencilOp(fail uint, zfail uint, zpass uint) {
	macro.Rewrite("$_.stencilOp($1, $2, $3)", fail, zfail, zpass)
}

func (*WebGLRenderingContext) StencilOpSeparate(face uint, fail uint, zfail uint, zpass uint) {
	macro.Rewrite("$_.stencilOpSeparate($1, $2, $3, $4)", face, fail, zfail, zpass)
}

func (*WebGLRenderingContext) TexImage2d(target uint, level int, internalformat uint, format uint, kind uint, pixels *canvas.ImageData) {
	macro.Rewrite("$_.texImage2D($1, $2, $3, $4, $5, $6)", target, level, internalformat, format, kind, pixels)
}

func (*WebGLRenderingContext) TexParameterf(target uint, pname uint, param float32) {
	macro.Rewrite("$_.texParameterf($1, $2, $3)", target, pname, param)
}

func (*WebGLRenderingContext) TexParameteri(target uint, pname uint, param int) {
	macro.Rewrite("$_.texParameteri($1, $2, $3)", target, pname, param)
}

func (*WebGLRenderingContext) TexSubImage2d(target uint, level int, xoffset int, yoffset int, format uint, kind uint, pixels *canvas.ImageData) {
	macro.Rewrite("$_.texSubImage2D($1, $2, $3, $4, $5, $6, $7)", target, level, xoffset, yoffset, format, kind, pixels)
}

func (*WebGLRenderingContext) Uniform1f(location *WebGLUniformLocation, x float32) {
	macro.Rewrite("$_.uniform1f($1, $2)", location, x)
}

func (*WebGLRenderingContext) Uniform1fv(location *WebGLUniformLocation, v []float32) {
	macro.Rewrite("$_.uniform1fv($1, $2)", location, v)
}

func (*WebGLRenderingContext) Uniform1i(location *WebGLUniformLocation, x int) {
	macro.Rewrite("$_.uniform1i($1, $2)", location, x)
}

func (*WebGLRenderingContext) Uniform1iv(location *WebGLUniformLocation, v []int32) {
	macro.Rewrite("$_.uniform1iv($1, $2)", location, v)
}

func (*WebGLRenderingContext) Uniform2f(location *WebGLUniformLocation, x float32, y float32) {
	macro.Rewrite("$_.uniform2f($1, $2, $3)", location, x, y)
}

func (*WebGLRenderingContext) Uniform2fv(location *WebGLUniformLocation, v []float32) {
	macro.Rewrite("$_.uniform2fv($1, $2)", location, v)
}

func (*WebGLRenderingContext) Uniform2i(location *WebGLUniformLocation, x int, y int) {
	macro.Rewrite("$_.uniform2i($1, $2, $3)", location, x, y)
}

func (*WebGLRenderingContext) Uniform2iv(location *WebGLUniformLocation, v []int32) {
	macro.Rewrite("$_.uniform2iv($1, $2)", location, v)
}

func (*WebGLRenderingContext) Uniform3f(location *WebGLUniformLocation, x float32, y float32, z float32) {
	macro.Rewrite("$_.uniform3f($1, $2, $3, $4)", location, x, y, z)
}

func (*WebGLRenderingContext) Uniform3fv(location *WebGLUniformLocation, v []float32) {
	macro.Rewrite("$_.uniform3fv($1, $2)", location, v)
}

func (*WebGLRenderingContext) Uniform3i(location *WebGLUniformLocation, x int, y int, z int) {
	macro.Rewrite("$_.uniform3i($1, $2, $3, $4)", location, x, y, z)
}

func (*WebGLRenderingContext) Uniform3iv(location *WebGLUniformLocation, v []int32) {
	macro.Rewrite("$_.uniform3iv($1, $2)", location, v)
}

func (*WebGLRenderingContext) Uniform4f(location *WebGLUniformLocation, x float32, y float32, z float32, w float32) {
	macro.Rewrite("$_.uniform4f($1, $2, $3, $4, $5)", location, x, y, z, w)
}

func (*WebGLRenderingContext) Uniform4fv(location *WebGLUniformLocation, v []float32) {
	macro.Rewrite("$_.uniform4fv($1, $2)", location, v)
}

func (*WebGLRenderingContext) Uniform4i(location *WebGLUniformLocation, x int, y int, z int, w int) {
	macro.Rewrite("$_.uniform4i($1, $2, $3, $4, $5)", location, x, y, z, w)
}

func (*WebGLRenderingContext) Uniform4iv(location *WebGLUniformLocation, v []int32) {
	macro.Rewrite("$_.uniform4iv($1, $2)", location, v)
}

func (*WebGLRenderingContext) UniformMatrix2fv(location *WebGLUniformLocation, transpose bool, value []float32) {
	macro.Rewrite("$_.uniformMatrix2fv($1, $2, $3)", location, transpose, value)
}

func (*WebGLRenderingContext) UniformMatrix3fv(location *WebGLUniformLocation, transpose bool, value []float32) {
	macro.Rewrite("$_.uniformMatrix3fv($1, $2, $3)", location, transpose, value)
}

func (*WebGLRenderingContext) UniformMatrix4fv(location *WebGLUniformLocation, transpose bool, value []float32) {
	macro.Rewrite("$_.uniformMatrix4fv($1, $2, $3)", location, transpose, value)
}

func (*WebGLRenderingContext) UseProgram(program *WebGLProgram) {
	macro.Rewrite("$_.useProgram($1)", program)
}

func (*WebGLRenderingContext) ValidateProgram(program *WebGLProgram) {
	macro.Rewrite("$_.validateProgram($1)", program)
}

func (*WebGLRenderingContext) VertexAttrib1f(indx uint, x float32) {
	macro.Rewrite("$_.vertexAttrib1f($1, $2)", indx, x)
}

func (*WebGLRenderingContext) VertexAttrib1fv(indx uint, values []float32) {
	macro.Rewrite("$_.vertexAttrib1fv($1, $2)", indx, values)
}

func (*WebGLRenderingContext) VertexAttrib2f(indx uint, x float32, y float32) {
	macro.Rewrite("$_.vertexAttrib2f($1, $2, $3)", indx, x, y)
}

func (*WebGLRenderingContext) VertexAttrib2fv(indx uint, values []float32) {
	macro.Rewrite("$_.vertexAttrib2fv($1, $2)", indx, values)
}

func (*WebGLRenderingContext) VertexAttrib3f(indx uint, x float32, y float32, z float32) {
	macro.Rewrite("$_.vertexAttrib3f($1, $2, $3, $4)", indx, x, y, z)
}

func (*WebGLRenderingContext) VertexAttrib3fv(indx uint, values []float32) {
	macro.Rewrite("$_.vertexAttrib3fv($1, $2)", indx, values)
}

func (*WebGLRenderingContext) VertexAttrib4f(indx uint, x float32, y float32, z float32, w float32) {
	macro.Rewrite("$_.vertexAttrib4f($1, $2, $3, $4, $5)", indx, x, y, z, w)
}

func (*WebGLRenderingContext) VertexAttrib4fv(indx uint, values []float32) {
	macro.Rewrite("$_.vertexAttrib4fv($1, $2)", indx, values)
}

func (*WebGLRenderingContext) VertexAttribPointer(indx uint, size int, kind uint, normalized bool, stride int, offset int64) {
	macro.Rewrite("$_.vertexAttribPointer($1, $2, $3, $4, $5, $6)", indx, size, kind, normalized, stride, offset)
}

func (*WebGLRenderingContext) Viewport(x int, y int, width int, height int) {
	macro.Rewrite("$_.viewport($1, $2, $3, $4)", x, y, width, height)
}

func (*WebGLRenderingContext) Canvas() (canvas *html.HTMLCanvasElement) {
	macro.Rewrite("$_.canvas")
	return canvas
}

func (*WebGLRenderingContext) DrawingBufferHeight() (drawingBufferHeight int) {
	macro.Rewrite("$_.drawingBufferHeight")
	return drawingBufferHeight
}

func (*WebGLRenderingContext) DrawingBufferWidth() (drawingBufferWidth int) {
	macro.Rewrite("$_.drawingBufferWidth")
	return drawingBufferWidth
}
