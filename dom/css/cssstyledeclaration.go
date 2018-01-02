package css

import "github.com/matthewmueller/joy/macro"

type CSSStyleDeclaration struct {
}

func (*CSSStyleDeclaration) GetPropertyPriority(propertyName string) (s string) {
	macro.Rewrite("$_.getPropertyPriority($1)", propertyName)
	return s
}

func (*CSSStyleDeclaration) GetPropertyValue(propertyName string) (s string) {
	macro.Rewrite("$_.getPropertyValue($1)", propertyName)
	return s
}

func (*CSSStyleDeclaration) Item(index uint) (s string) {
	macro.Rewrite("$_.item($1)", index)
	return s
}

func (*CSSStyleDeclaration) RemoveProperty(propertyName string) (s string) {
	macro.Rewrite("$_.removeProperty($1)", propertyName)
	return s
}

func (*CSSStyleDeclaration) SetProperty(propertyName string, value string, priority *string) {
	macro.Rewrite("$_.setProperty($1, $2, $3)", propertyName, value, priority)
}

func (*CSSStyleDeclaration) AlignContent() (alignContent string) {
	macro.Rewrite("$_.alignContent")
	return alignContent
}

func (*CSSStyleDeclaration) SetAlignContent(alignContent string) {
	macro.Rewrite("$_.alignContent = $1", alignContent)
}

func (*CSSStyleDeclaration) AlignItems() (alignItems string) {
	macro.Rewrite("$_.alignItems")
	return alignItems
}

func (*CSSStyleDeclaration) SetAlignItems(alignItems string) {
	macro.Rewrite("$_.alignItems = $1", alignItems)
}

func (*CSSStyleDeclaration) AlignmentBaseline() (alignmentBaseline string) {
	macro.Rewrite("$_.alignmentBaseline")
	return alignmentBaseline
}

func (*CSSStyleDeclaration) SetAlignmentBaseline(alignmentBaseline string) {
	macro.Rewrite("$_.alignmentBaseline = $1", alignmentBaseline)
}

func (*CSSStyleDeclaration) AlignSelf() (alignSelf string) {
	macro.Rewrite("$_.alignSelf")
	return alignSelf
}

func (*CSSStyleDeclaration) SetAlignSelf(alignSelf string) {
	macro.Rewrite("$_.alignSelf = $1", alignSelf)
}

func (*CSSStyleDeclaration) Animation() (animation string) {
	macro.Rewrite("$_.animation")
	return animation
}

func (*CSSStyleDeclaration) SetAnimation(animation string) {
	macro.Rewrite("$_.animation = $1", animation)
}

func (*CSSStyleDeclaration) AnimationDelay() (animationDelay string) {
	macro.Rewrite("$_.animationDelay")
	return animationDelay
}

func (*CSSStyleDeclaration) SetAnimationDelay(animationDelay string) {
	macro.Rewrite("$_.animationDelay = $1", animationDelay)
}

func (*CSSStyleDeclaration) AnimationDirection() (animationDirection string) {
	macro.Rewrite("$_.animationDirection")
	return animationDirection
}

func (*CSSStyleDeclaration) SetAnimationDirection(animationDirection string) {
	macro.Rewrite("$_.animationDirection = $1", animationDirection)
}

func (*CSSStyleDeclaration) AnimationDuration() (animationDuration string) {
	macro.Rewrite("$_.animationDuration")
	return animationDuration
}

func (*CSSStyleDeclaration) SetAnimationDuration(animationDuration string) {
	macro.Rewrite("$_.animationDuration = $1", animationDuration)
}

func (*CSSStyleDeclaration) AnimationFillMode() (animationFillMode string) {
	macro.Rewrite("$_.animationFillMode")
	return animationFillMode
}

func (*CSSStyleDeclaration) SetAnimationFillMode(animationFillMode string) {
	macro.Rewrite("$_.animationFillMode = $1", animationFillMode)
}

func (*CSSStyleDeclaration) AnimationIterationCount() (animationIterationCount string) {
	macro.Rewrite("$_.animationIterationCount")
	return animationIterationCount
}

func (*CSSStyleDeclaration) SetAnimationIterationCount(animationIterationCount string) {
	macro.Rewrite("$_.animationIterationCount = $1", animationIterationCount)
}

func (*CSSStyleDeclaration) AnimationName() (animationName string) {
	macro.Rewrite("$_.animationName")
	return animationName
}

func (*CSSStyleDeclaration) SetAnimationName(animationName string) {
	macro.Rewrite("$_.animationName = $1", animationName)
}

func (*CSSStyleDeclaration) AnimationPlayState() (animationPlayState string) {
	macro.Rewrite("$_.animationPlayState")
	return animationPlayState
}

func (*CSSStyleDeclaration) SetAnimationPlayState(animationPlayState string) {
	macro.Rewrite("$_.animationPlayState = $1", animationPlayState)
}

func (*CSSStyleDeclaration) AnimationTimingFunction() (animationTimingFunction string) {
	macro.Rewrite("$_.animationTimingFunction")
	return animationTimingFunction
}

func (*CSSStyleDeclaration) SetAnimationTimingFunction(animationTimingFunction string) {
	macro.Rewrite("$_.animationTimingFunction = $1", animationTimingFunction)
}

func (*CSSStyleDeclaration) BackfaceVisibility() (backfaceVisibility string) {
	macro.Rewrite("$_.backfaceVisibility")
	return backfaceVisibility
}

func (*CSSStyleDeclaration) SetBackfaceVisibility(backfaceVisibility string) {
	macro.Rewrite("$_.backfaceVisibility = $1", backfaceVisibility)
}

func (*CSSStyleDeclaration) Background() (background string) {
	macro.Rewrite("$_.background")
	return background
}

func (*CSSStyleDeclaration) SetBackground(background string) {
	macro.Rewrite("$_.background = $1", background)
}

func (*CSSStyleDeclaration) BackgroundAttachment() (backgroundAttachment string) {
	macro.Rewrite("$_.backgroundAttachment")
	return backgroundAttachment
}

func (*CSSStyleDeclaration) SetBackgroundAttachment(backgroundAttachment string) {
	macro.Rewrite("$_.backgroundAttachment = $1", backgroundAttachment)
}

func (*CSSStyleDeclaration) BackgroundClip() (backgroundClip string) {
	macro.Rewrite("$_.backgroundClip")
	return backgroundClip
}

func (*CSSStyleDeclaration) SetBackgroundClip(backgroundClip string) {
	macro.Rewrite("$_.backgroundClip = $1", backgroundClip)
}

func (*CSSStyleDeclaration) BackgroundColor() (backgroundColor string) {
	macro.Rewrite("$_.backgroundColor")
	return backgroundColor
}

func (*CSSStyleDeclaration) SetBackgroundColor(backgroundColor string) {
	macro.Rewrite("$_.backgroundColor = $1", backgroundColor)
}

func (*CSSStyleDeclaration) BackgroundImage() (backgroundImage string) {
	macro.Rewrite("$_.backgroundImage")
	return backgroundImage
}

func (*CSSStyleDeclaration) SetBackgroundImage(backgroundImage string) {
	macro.Rewrite("$_.backgroundImage = $1", backgroundImage)
}

func (*CSSStyleDeclaration) BackgroundOrigin() (backgroundOrigin string) {
	macro.Rewrite("$_.backgroundOrigin")
	return backgroundOrigin
}

func (*CSSStyleDeclaration) SetBackgroundOrigin(backgroundOrigin string) {
	macro.Rewrite("$_.backgroundOrigin = $1", backgroundOrigin)
}

func (*CSSStyleDeclaration) BackgroundPosition() (backgroundPosition string) {
	macro.Rewrite("$_.backgroundPosition")
	return backgroundPosition
}

func (*CSSStyleDeclaration) SetBackgroundPosition(backgroundPosition string) {
	macro.Rewrite("$_.backgroundPosition = $1", backgroundPosition)
}

func (*CSSStyleDeclaration) BackgroundPositionX() (backgroundPositionX string) {
	macro.Rewrite("$_.backgroundPositionX")
	return backgroundPositionX
}

func (*CSSStyleDeclaration) SetBackgroundPositionX(backgroundPositionX string) {
	macro.Rewrite("$_.backgroundPositionX = $1", backgroundPositionX)
}

func (*CSSStyleDeclaration) BackgroundPositionY() (backgroundPositionY string) {
	macro.Rewrite("$_.backgroundPositionY")
	return backgroundPositionY
}

func (*CSSStyleDeclaration) SetBackgroundPositionY(backgroundPositionY string) {
	macro.Rewrite("$_.backgroundPositionY = $1", backgroundPositionY)
}

func (*CSSStyleDeclaration) BackgroundRepeat() (backgroundRepeat string) {
	macro.Rewrite("$_.backgroundRepeat")
	return backgroundRepeat
}

func (*CSSStyleDeclaration) SetBackgroundRepeat(backgroundRepeat string) {
	macro.Rewrite("$_.backgroundRepeat = $1", backgroundRepeat)
}

func (*CSSStyleDeclaration) BackgroundSize() (backgroundSize string) {
	macro.Rewrite("$_.backgroundSize")
	return backgroundSize
}

func (*CSSStyleDeclaration) SetBackgroundSize(backgroundSize string) {
	macro.Rewrite("$_.backgroundSize = $1", backgroundSize)
}

func (*CSSStyleDeclaration) BaselineShift() (baselineShift string) {
	macro.Rewrite("$_.baselineShift")
	return baselineShift
}

func (*CSSStyleDeclaration) SetBaselineShift(baselineShift string) {
	macro.Rewrite("$_.baselineShift = $1", baselineShift)
}

func (*CSSStyleDeclaration) Border() (border string) {
	macro.Rewrite("$_.border")
	return border
}

func (*CSSStyleDeclaration) SetBorder(border string) {
	macro.Rewrite("$_.border = $1", border)
}

func (*CSSStyleDeclaration) BorderBottom() (borderBottom string) {
	macro.Rewrite("$_.borderBottom")
	return borderBottom
}

func (*CSSStyleDeclaration) SetBorderBottom(borderBottom string) {
	macro.Rewrite("$_.borderBottom = $1", borderBottom)
}

func (*CSSStyleDeclaration) BorderBottomColor() (borderBottomColor string) {
	macro.Rewrite("$_.borderBottomColor")
	return borderBottomColor
}

func (*CSSStyleDeclaration) SetBorderBottomColor(borderBottomColor string) {
	macro.Rewrite("$_.borderBottomColor = $1", borderBottomColor)
}

func (*CSSStyleDeclaration) BorderBottomLeftRadius() (borderBottomLeftRadius string) {
	macro.Rewrite("$_.borderBottomLeftRadius")
	return borderBottomLeftRadius
}

func (*CSSStyleDeclaration) SetBorderBottomLeftRadius(borderBottomLeftRadius string) {
	macro.Rewrite("$_.borderBottomLeftRadius = $1", borderBottomLeftRadius)
}

func (*CSSStyleDeclaration) BorderBottomRightRadius() (borderBottomRightRadius string) {
	macro.Rewrite("$_.borderBottomRightRadius")
	return borderBottomRightRadius
}

func (*CSSStyleDeclaration) SetBorderBottomRightRadius(borderBottomRightRadius string) {
	macro.Rewrite("$_.borderBottomRightRadius = $1", borderBottomRightRadius)
}

func (*CSSStyleDeclaration) BorderBottomStyle() (borderBottomStyle string) {
	macro.Rewrite("$_.borderBottomStyle")
	return borderBottomStyle
}

func (*CSSStyleDeclaration) SetBorderBottomStyle(borderBottomStyle string) {
	macro.Rewrite("$_.borderBottomStyle = $1", borderBottomStyle)
}

func (*CSSStyleDeclaration) BorderBottomWidth() (borderBottomWidth string) {
	macro.Rewrite("$_.borderBottomWidth")
	return borderBottomWidth
}

func (*CSSStyleDeclaration) SetBorderBottomWidth(borderBottomWidth string) {
	macro.Rewrite("$_.borderBottomWidth = $1", borderBottomWidth)
}

func (*CSSStyleDeclaration) BorderCollapse() (borderCollapse string) {
	macro.Rewrite("$_.borderCollapse")
	return borderCollapse
}

func (*CSSStyleDeclaration) SetBorderCollapse(borderCollapse string) {
	macro.Rewrite("$_.borderCollapse = $1", borderCollapse)
}

func (*CSSStyleDeclaration) BorderColor() (borderColor string) {
	macro.Rewrite("$_.borderColor")
	return borderColor
}

func (*CSSStyleDeclaration) SetBorderColor(borderColor string) {
	macro.Rewrite("$_.borderColor = $1", borderColor)
}

func (*CSSStyleDeclaration) BorderImage() (borderImage string) {
	macro.Rewrite("$_.borderImage")
	return borderImage
}

func (*CSSStyleDeclaration) SetBorderImage(borderImage string) {
	macro.Rewrite("$_.borderImage = $1", borderImage)
}

func (*CSSStyleDeclaration) BorderImageOutset() (borderImageOutset string) {
	macro.Rewrite("$_.borderImageOutset")
	return borderImageOutset
}

func (*CSSStyleDeclaration) SetBorderImageOutset(borderImageOutset string) {
	macro.Rewrite("$_.borderImageOutset = $1", borderImageOutset)
}

func (*CSSStyleDeclaration) BorderImageRepeat() (borderImageRepeat string) {
	macro.Rewrite("$_.borderImageRepeat")
	return borderImageRepeat
}

func (*CSSStyleDeclaration) SetBorderImageRepeat(borderImageRepeat string) {
	macro.Rewrite("$_.borderImageRepeat = $1", borderImageRepeat)
}

func (*CSSStyleDeclaration) BorderImageSlice() (borderImageSlice string) {
	macro.Rewrite("$_.borderImageSlice")
	return borderImageSlice
}

func (*CSSStyleDeclaration) SetBorderImageSlice(borderImageSlice string) {
	macro.Rewrite("$_.borderImageSlice = $1", borderImageSlice)
}

func (*CSSStyleDeclaration) BorderImageSource() (borderImageSource string) {
	macro.Rewrite("$_.borderImageSource")
	return borderImageSource
}

func (*CSSStyleDeclaration) SetBorderImageSource(borderImageSource string) {
	macro.Rewrite("$_.borderImageSource = $1", borderImageSource)
}

func (*CSSStyleDeclaration) BorderImageWidth() (borderImageWidth string) {
	macro.Rewrite("$_.borderImageWidth")
	return borderImageWidth
}

func (*CSSStyleDeclaration) SetBorderImageWidth(borderImageWidth string) {
	macro.Rewrite("$_.borderImageWidth = $1", borderImageWidth)
}

func (*CSSStyleDeclaration) BorderLeft() (borderLeft string) {
	macro.Rewrite("$_.borderLeft")
	return borderLeft
}

func (*CSSStyleDeclaration) SetBorderLeft(borderLeft string) {
	macro.Rewrite("$_.borderLeft = $1", borderLeft)
}

func (*CSSStyleDeclaration) BorderLeftColor() (borderLeftColor string) {
	macro.Rewrite("$_.borderLeftColor")
	return borderLeftColor
}

func (*CSSStyleDeclaration) SetBorderLeftColor(borderLeftColor string) {
	macro.Rewrite("$_.borderLeftColor = $1", borderLeftColor)
}

func (*CSSStyleDeclaration) BorderLeftStyle() (borderLeftStyle string) {
	macro.Rewrite("$_.borderLeftStyle")
	return borderLeftStyle
}

func (*CSSStyleDeclaration) SetBorderLeftStyle(borderLeftStyle string) {
	macro.Rewrite("$_.borderLeftStyle = $1", borderLeftStyle)
}

func (*CSSStyleDeclaration) BorderLeftWidth() (borderLeftWidth string) {
	macro.Rewrite("$_.borderLeftWidth")
	return borderLeftWidth
}

func (*CSSStyleDeclaration) SetBorderLeftWidth(borderLeftWidth string) {
	macro.Rewrite("$_.borderLeftWidth = $1", borderLeftWidth)
}

func (*CSSStyleDeclaration) BorderRadius() (borderRadius string) {
	macro.Rewrite("$_.borderRadius")
	return borderRadius
}

func (*CSSStyleDeclaration) SetBorderRadius(borderRadius string) {
	macro.Rewrite("$_.borderRadius = $1", borderRadius)
}

func (*CSSStyleDeclaration) BorderRight() (borderRight string) {
	macro.Rewrite("$_.borderRight")
	return borderRight
}

func (*CSSStyleDeclaration) SetBorderRight(borderRight string) {
	macro.Rewrite("$_.borderRight = $1", borderRight)
}

func (*CSSStyleDeclaration) BorderRightColor() (borderRightColor string) {
	macro.Rewrite("$_.borderRightColor")
	return borderRightColor
}

func (*CSSStyleDeclaration) SetBorderRightColor(borderRightColor string) {
	macro.Rewrite("$_.borderRightColor = $1", borderRightColor)
}

func (*CSSStyleDeclaration) BorderRightStyle() (borderRightStyle string) {
	macro.Rewrite("$_.borderRightStyle")
	return borderRightStyle
}

func (*CSSStyleDeclaration) SetBorderRightStyle(borderRightStyle string) {
	macro.Rewrite("$_.borderRightStyle = $1", borderRightStyle)
}

func (*CSSStyleDeclaration) BorderRightWidth() (borderRightWidth string) {
	macro.Rewrite("$_.borderRightWidth")
	return borderRightWidth
}

func (*CSSStyleDeclaration) SetBorderRightWidth(borderRightWidth string) {
	macro.Rewrite("$_.borderRightWidth = $1", borderRightWidth)
}

func (*CSSStyleDeclaration) BorderSpacing() (borderSpacing string) {
	macro.Rewrite("$_.borderSpacing")
	return borderSpacing
}

func (*CSSStyleDeclaration) SetBorderSpacing(borderSpacing string) {
	macro.Rewrite("$_.borderSpacing = $1", borderSpacing)
}

func (*CSSStyleDeclaration) BorderStyle() (borderStyle string) {
	macro.Rewrite("$_.borderStyle")
	return borderStyle
}

func (*CSSStyleDeclaration) SetBorderStyle(borderStyle string) {
	macro.Rewrite("$_.borderStyle = $1", borderStyle)
}

func (*CSSStyleDeclaration) BorderTop() (borderTop string) {
	macro.Rewrite("$_.borderTop")
	return borderTop
}

func (*CSSStyleDeclaration) SetBorderTop(borderTop string) {
	macro.Rewrite("$_.borderTop = $1", borderTop)
}

func (*CSSStyleDeclaration) BorderTopColor() (borderTopColor string) {
	macro.Rewrite("$_.borderTopColor")
	return borderTopColor
}

func (*CSSStyleDeclaration) SetBorderTopColor(borderTopColor string) {
	macro.Rewrite("$_.borderTopColor = $1", borderTopColor)
}

func (*CSSStyleDeclaration) BorderTopLeftRadius() (borderTopLeftRadius string) {
	macro.Rewrite("$_.borderTopLeftRadius")
	return borderTopLeftRadius
}

func (*CSSStyleDeclaration) SetBorderTopLeftRadius(borderTopLeftRadius string) {
	macro.Rewrite("$_.borderTopLeftRadius = $1", borderTopLeftRadius)
}

func (*CSSStyleDeclaration) BorderTopRightRadius() (borderTopRightRadius string) {
	macro.Rewrite("$_.borderTopRightRadius")
	return borderTopRightRadius
}

func (*CSSStyleDeclaration) SetBorderTopRightRadius(borderTopRightRadius string) {
	macro.Rewrite("$_.borderTopRightRadius = $1", borderTopRightRadius)
}

func (*CSSStyleDeclaration) BorderTopStyle() (borderTopStyle string) {
	macro.Rewrite("$_.borderTopStyle")
	return borderTopStyle
}

func (*CSSStyleDeclaration) SetBorderTopStyle(borderTopStyle string) {
	macro.Rewrite("$_.borderTopStyle = $1", borderTopStyle)
}

func (*CSSStyleDeclaration) BorderTopWidth() (borderTopWidth string) {
	macro.Rewrite("$_.borderTopWidth")
	return borderTopWidth
}

func (*CSSStyleDeclaration) SetBorderTopWidth(borderTopWidth string) {
	macro.Rewrite("$_.borderTopWidth = $1", borderTopWidth)
}

func (*CSSStyleDeclaration) BorderWidth() (borderWidth string) {
	macro.Rewrite("$_.borderWidth")
	return borderWidth
}

func (*CSSStyleDeclaration) SetBorderWidth(borderWidth string) {
	macro.Rewrite("$_.borderWidth = $1", borderWidth)
}

func (*CSSStyleDeclaration) Bottom() (bottom string) {
	macro.Rewrite("$_.bottom")
	return bottom
}

func (*CSSStyleDeclaration) SetBottom(bottom string) {
	macro.Rewrite("$_.bottom = $1", bottom)
}

func (*CSSStyleDeclaration) BoxShadow() (boxShadow string) {
	macro.Rewrite("$_.boxShadow")
	return boxShadow
}

func (*CSSStyleDeclaration) SetBoxShadow(boxShadow string) {
	macro.Rewrite("$_.boxShadow = $1", boxShadow)
}

func (*CSSStyleDeclaration) BoxSizing() (boxSizing string) {
	macro.Rewrite("$_.boxSizing")
	return boxSizing
}

func (*CSSStyleDeclaration) SetBoxSizing(boxSizing string) {
	macro.Rewrite("$_.boxSizing = $1", boxSizing)
}

func (*CSSStyleDeclaration) BreakAfter() (breakAfter string) {
	macro.Rewrite("$_.breakAfter")
	return breakAfter
}

func (*CSSStyleDeclaration) SetBreakAfter(breakAfter string) {
	macro.Rewrite("$_.breakAfter = $1", breakAfter)
}

func (*CSSStyleDeclaration) BreakBefore() (breakBefore string) {
	macro.Rewrite("$_.breakBefore")
	return breakBefore
}

func (*CSSStyleDeclaration) SetBreakBefore(breakBefore string) {
	macro.Rewrite("$_.breakBefore = $1", breakBefore)
}

func (*CSSStyleDeclaration) BreakInside() (breakInside string) {
	macro.Rewrite("$_.breakInside")
	return breakInside
}

func (*CSSStyleDeclaration) SetBreakInside(breakInside string) {
	macro.Rewrite("$_.breakInside = $1", breakInside)
}

func (*CSSStyleDeclaration) CaptionSide() (captionSide string) {
	macro.Rewrite("$_.captionSide")
	return captionSide
}

func (*CSSStyleDeclaration) SetCaptionSide(captionSide string) {
	macro.Rewrite("$_.captionSide = $1", captionSide)
}

func (*CSSStyleDeclaration) Clear() (clear string) {
	macro.Rewrite("$_.clear")
	return clear
}

func (*CSSStyleDeclaration) SetClear(clear string) {
	macro.Rewrite("$_.clear = $1", clear)
}

func (*CSSStyleDeclaration) Clip() (clip string) {
	macro.Rewrite("$_.clip")
	return clip
}

func (*CSSStyleDeclaration) SetClip(clip string) {
	macro.Rewrite("$_.clip = $1", clip)
}

func (*CSSStyleDeclaration) ClipPath() (clipPath string) {
	macro.Rewrite("$_.clipPath")
	return clipPath
}

func (*CSSStyleDeclaration) SetClipPath(clipPath string) {
	macro.Rewrite("$_.clipPath = $1", clipPath)
}

func (*CSSStyleDeclaration) ClipRule() (clipRule string) {
	macro.Rewrite("$_.clipRule")
	return clipRule
}

func (*CSSStyleDeclaration) SetClipRule(clipRule string) {
	macro.Rewrite("$_.clipRule = $1", clipRule)
}

func (*CSSStyleDeclaration) Color() (color string) {
	macro.Rewrite("$_.color")
	return color
}

func (*CSSStyleDeclaration) SetColor(color string) {
	macro.Rewrite("$_.color = $1", color)
}

func (*CSSStyleDeclaration) ColorInterpolationFilters() (colorInterpolationFilters string) {
	macro.Rewrite("$_.colorInterpolationFilters")
	return colorInterpolationFilters
}

func (*CSSStyleDeclaration) SetColorInterpolationFilters(colorInterpolationFilters string) {
	macro.Rewrite("$_.colorInterpolationFilters = $1", colorInterpolationFilters)
}

func (*CSSStyleDeclaration) ColumnCount() (columnCount interface{}) {
	macro.Rewrite("$_.columnCount")
	return columnCount
}

func (*CSSStyleDeclaration) SetColumnCount(columnCount interface{}) {
	macro.Rewrite("$_.columnCount = $1", columnCount)
}

func (*CSSStyleDeclaration) ColumnFill() (columnFill string) {
	macro.Rewrite("$_.columnFill")
	return columnFill
}

func (*CSSStyleDeclaration) SetColumnFill(columnFill string) {
	macro.Rewrite("$_.columnFill = $1", columnFill)
}

func (*CSSStyleDeclaration) ColumnGap() (columnGap interface{}) {
	macro.Rewrite("$_.columnGap")
	return columnGap
}

func (*CSSStyleDeclaration) SetColumnGap(columnGap interface{}) {
	macro.Rewrite("$_.columnGap = $1", columnGap)
}

func (*CSSStyleDeclaration) ColumnRule() (columnRule string) {
	macro.Rewrite("$_.columnRule")
	return columnRule
}

func (*CSSStyleDeclaration) SetColumnRule(columnRule string) {
	macro.Rewrite("$_.columnRule = $1", columnRule)
}

func (*CSSStyleDeclaration) ColumnRuleColor() (columnRuleColor interface{}) {
	macro.Rewrite("$_.columnRuleColor")
	return columnRuleColor
}

func (*CSSStyleDeclaration) SetColumnRuleColor(columnRuleColor interface{}) {
	macro.Rewrite("$_.columnRuleColor = $1", columnRuleColor)
}

func (*CSSStyleDeclaration) ColumnRuleStyle() (columnRuleStyle string) {
	macro.Rewrite("$_.columnRuleStyle")
	return columnRuleStyle
}

func (*CSSStyleDeclaration) SetColumnRuleStyle(columnRuleStyle string) {
	macro.Rewrite("$_.columnRuleStyle = $1", columnRuleStyle)
}

func (*CSSStyleDeclaration) ColumnRuleWidth() (columnRuleWidth interface{}) {
	macro.Rewrite("$_.columnRuleWidth")
	return columnRuleWidth
}

func (*CSSStyleDeclaration) SetColumnRuleWidth(columnRuleWidth interface{}) {
	macro.Rewrite("$_.columnRuleWidth = $1", columnRuleWidth)
}

func (*CSSStyleDeclaration) Columns() (columns string) {
	macro.Rewrite("$_.columns")
	return columns
}

func (*CSSStyleDeclaration) SetColumns(columns string) {
	macro.Rewrite("$_.columns = $1", columns)
}

func (*CSSStyleDeclaration) ColumnSpan() (columnSpan string) {
	macro.Rewrite("$_.columnSpan")
	return columnSpan
}

func (*CSSStyleDeclaration) SetColumnSpan(columnSpan string) {
	macro.Rewrite("$_.columnSpan = $1", columnSpan)
}

func (*CSSStyleDeclaration) ColumnWidth() (columnWidth interface{}) {
	macro.Rewrite("$_.columnWidth")
	return columnWidth
}

func (*CSSStyleDeclaration) SetColumnWidth(columnWidth interface{}) {
	macro.Rewrite("$_.columnWidth = $1", columnWidth)
}

func (*CSSStyleDeclaration) Content() (content string) {
	macro.Rewrite("$_.content")
	return content
}

func (*CSSStyleDeclaration) SetContent(content string) {
	macro.Rewrite("$_.content = $1", content)
}

func (*CSSStyleDeclaration) CounterIncrement() (counterIncrement string) {
	macro.Rewrite("$_.counterIncrement")
	return counterIncrement
}

func (*CSSStyleDeclaration) SetCounterIncrement(counterIncrement string) {
	macro.Rewrite("$_.counterIncrement = $1", counterIncrement)
}

func (*CSSStyleDeclaration) CounterReset() (counterReset string) {
	macro.Rewrite("$_.counterReset")
	return counterReset
}

func (*CSSStyleDeclaration) SetCounterReset(counterReset string) {
	macro.Rewrite("$_.counterReset = $1", counterReset)
}

func (*CSSStyleDeclaration) CSSFloat() (cssFloat string) {
	macro.Rewrite("$_.cssFloat")
	return cssFloat
}

func (*CSSStyleDeclaration) SetCSSFloat(cssFloat string) {
	macro.Rewrite("$_.cssFloat = $1", cssFloat)
}

func (*CSSStyleDeclaration) CSSText() (cssText string) {
	macro.Rewrite("$_.cssText")
	return cssText
}

func (*CSSStyleDeclaration) SetCSSText(cssText string) {
	macro.Rewrite("$_.cssText = $1", cssText)
}

func (*CSSStyleDeclaration) Cursor() (cursor string) {
	macro.Rewrite("$_.cursor")
	return cursor
}

func (*CSSStyleDeclaration) SetCursor(cursor string) {
	macro.Rewrite("$_.cursor = $1", cursor)
}

func (*CSSStyleDeclaration) Direction() (direction string) {
	macro.Rewrite("$_.direction")
	return direction
}

func (*CSSStyleDeclaration) SetDirection(direction string) {
	macro.Rewrite("$_.direction = $1", direction)
}

func (*CSSStyleDeclaration) Display() (display string) {
	macro.Rewrite("$_.display")
	return display
}

func (*CSSStyleDeclaration) SetDisplay(display string) {
	macro.Rewrite("$_.display = $1", display)
}

func (*CSSStyleDeclaration) DominantBaseline() (dominantBaseline string) {
	macro.Rewrite("$_.dominantBaseline")
	return dominantBaseline
}

func (*CSSStyleDeclaration) SetDominantBaseline(dominantBaseline string) {
	macro.Rewrite("$_.dominantBaseline = $1", dominantBaseline)
}

func (*CSSStyleDeclaration) EmptyCells() (emptyCells string) {
	macro.Rewrite("$_.emptyCells")
	return emptyCells
}

func (*CSSStyleDeclaration) SetEmptyCells(emptyCells string) {
	macro.Rewrite("$_.emptyCells = $1", emptyCells)
}

func (*CSSStyleDeclaration) EnableBackground() (enableBackground string) {
	macro.Rewrite("$_.enableBackground")
	return enableBackground
}

func (*CSSStyleDeclaration) SetEnableBackground(enableBackground string) {
	macro.Rewrite("$_.enableBackground = $1", enableBackground)
}

func (*CSSStyleDeclaration) Fill() (fill string) {
	macro.Rewrite("$_.fill")
	return fill
}

func (*CSSStyleDeclaration) SetFill(fill string) {
	macro.Rewrite("$_.fill = $1", fill)
}

func (*CSSStyleDeclaration) FillOpacity() (fillOpacity string) {
	macro.Rewrite("$_.fillOpacity")
	return fillOpacity
}

func (*CSSStyleDeclaration) SetFillOpacity(fillOpacity string) {
	macro.Rewrite("$_.fillOpacity = $1", fillOpacity)
}

func (*CSSStyleDeclaration) FillRule() (fillRule string) {
	macro.Rewrite("$_.fillRule")
	return fillRule
}

func (*CSSStyleDeclaration) SetFillRule(fillRule string) {
	macro.Rewrite("$_.fillRule = $1", fillRule)
}

func (*CSSStyleDeclaration) Filter() (filter string) {
	macro.Rewrite("$_.filter")
	return filter
}

func (*CSSStyleDeclaration) SetFilter(filter string) {
	macro.Rewrite("$_.filter = $1", filter)
}

func (*CSSStyleDeclaration) Flex() (flex string) {
	macro.Rewrite("$_.flex")
	return flex
}

func (*CSSStyleDeclaration) SetFlex(flex string) {
	macro.Rewrite("$_.flex = $1", flex)
}

func (*CSSStyleDeclaration) FlexBasis() (flexBasis string) {
	macro.Rewrite("$_.flexBasis")
	return flexBasis
}

func (*CSSStyleDeclaration) SetFlexBasis(flexBasis string) {
	macro.Rewrite("$_.flexBasis = $1", flexBasis)
}

func (*CSSStyleDeclaration) FlexDirection() (flexDirection string) {
	macro.Rewrite("$_.flexDirection")
	return flexDirection
}

func (*CSSStyleDeclaration) SetFlexDirection(flexDirection string) {
	macro.Rewrite("$_.flexDirection = $1", flexDirection)
}

func (*CSSStyleDeclaration) FlexFlow() (flexFlow string) {
	macro.Rewrite("$_.flexFlow")
	return flexFlow
}

func (*CSSStyleDeclaration) SetFlexFlow(flexFlow string) {
	macro.Rewrite("$_.flexFlow = $1", flexFlow)
}

func (*CSSStyleDeclaration) FlexGrow() (flexGrow string) {
	macro.Rewrite("$_.flexGrow")
	return flexGrow
}

func (*CSSStyleDeclaration) SetFlexGrow(flexGrow string) {
	macro.Rewrite("$_.flexGrow = $1", flexGrow)
}

func (*CSSStyleDeclaration) FlexShrink() (flexShrink string) {
	macro.Rewrite("$_.flexShrink")
	return flexShrink
}

func (*CSSStyleDeclaration) SetFlexShrink(flexShrink string) {
	macro.Rewrite("$_.flexShrink = $1", flexShrink)
}

func (*CSSStyleDeclaration) FlexWrap() (flexWrap string) {
	macro.Rewrite("$_.flexWrap")
	return flexWrap
}

func (*CSSStyleDeclaration) SetFlexWrap(flexWrap string) {
	macro.Rewrite("$_.flexWrap = $1", flexWrap)
}

func (*CSSStyleDeclaration) FloodColor() (floodColor string) {
	macro.Rewrite("$_.floodColor")
	return floodColor
}

func (*CSSStyleDeclaration) SetFloodColor(floodColor string) {
	macro.Rewrite("$_.floodColor = $1", floodColor)
}

func (*CSSStyleDeclaration) FloodOpacity() (floodOpacity string) {
	macro.Rewrite("$_.floodOpacity")
	return floodOpacity
}

func (*CSSStyleDeclaration) SetFloodOpacity(floodOpacity string) {
	macro.Rewrite("$_.floodOpacity = $1", floodOpacity)
}

func (*CSSStyleDeclaration) Font() (font string) {
	macro.Rewrite("$_.font")
	return font
}

func (*CSSStyleDeclaration) SetFont(font string) {
	macro.Rewrite("$_.font = $1", font)
}

func (*CSSStyleDeclaration) FontFamily() (fontFamily string) {
	macro.Rewrite("$_.fontFamily")
	return fontFamily
}

func (*CSSStyleDeclaration) SetFontFamily(fontFamily string) {
	macro.Rewrite("$_.fontFamily = $1", fontFamily)
}

func (*CSSStyleDeclaration) FontFeatureSettings() (fontFeatureSettings string) {
	macro.Rewrite("$_.fontFeatureSettings")
	return fontFeatureSettings
}

func (*CSSStyleDeclaration) SetFontFeatureSettings(fontFeatureSettings string) {
	macro.Rewrite("$_.fontFeatureSettings = $1", fontFeatureSettings)
}

func (*CSSStyleDeclaration) FontSize() (fontSize string) {
	macro.Rewrite("$_.fontSize")
	return fontSize
}

func (*CSSStyleDeclaration) SetFontSize(fontSize string) {
	macro.Rewrite("$_.fontSize = $1", fontSize)
}

func (*CSSStyleDeclaration) FontSizeAdjust() (fontSizeAdjust string) {
	macro.Rewrite("$_.fontSizeAdjust")
	return fontSizeAdjust
}

func (*CSSStyleDeclaration) SetFontSizeAdjust(fontSizeAdjust string) {
	macro.Rewrite("$_.fontSizeAdjust = $1", fontSizeAdjust)
}

func (*CSSStyleDeclaration) FontStretch() (fontStretch string) {
	macro.Rewrite("$_.fontStretch")
	return fontStretch
}

func (*CSSStyleDeclaration) SetFontStretch(fontStretch string) {
	macro.Rewrite("$_.fontStretch = $1", fontStretch)
}

func (*CSSStyleDeclaration) FontStyle() (fontStyle string) {
	macro.Rewrite("$_.fontStyle")
	return fontStyle
}

func (*CSSStyleDeclaration) SetFontStyle(fontStyle string) {
	macro.Rewrite("$_.fontStyle = $1", fontStyle)
}

func (*CSSStyleDeclaration) FontVariant() (fontVariant string) {
	macro.Rewrite("$_.fontVariant")
	return fontVariant
}

func (*CSSStyleDeclaration) SetFontVariant(fontVariant string) {
	macro.Rewrite("$_.fontVariant = $1", fontVariant)
}

func (*CSSStyleDeclaration) FontWeight() (fontWeight string) {
	macro.Rewrite("$_.fontWeight")
	return fontWeight
}

func (*CSSStyleDeclaration) SetFontWeight(fontWeight string) {
	macro.Rewrite("$_.fontWeight = $1", fontWeight)
}

func (*CSSStyleDeclaration) GlyphOrientationHorizontal() (glyphOrientationHorizontal string) {
	macro.Rewrite("$_.glyphOrientationHorizontal")
	return glyphOrientationHorizontal
}

func (*CSSStyleDeclaration) SetGlyphOrientationHorizontal(glyphOrientationHorizontal string) {
	macro.Rewrite("$_.glyphOrientationHorizontal = $1", glyphOrientationHorizontal)
}

func (*CSSStyleDeclaration) GlyphOrientationVertical() (glyphOrientationVertical string) {
	macro.Rewrite("$_.glyphOrientationVertical")
	return glyphOrientationVertical
}

func (*CSSStyleDeclaration) SetGlyphOrientationVertical(glyphOrientationVertical string) {
	macro.Rewrite("$_.glyphOrientationVertical = $1", glyphOrientationVertical)
}

func (*CSSStyleDeclaration) Height() (height string) {
	macro.Rewrite("$_.height")
	return height
}

func (*CSSStyleDeclaration) SetHeight(height string) {
	macro.Rewrite("$_.height = $1", height)
}

func (*CSSStyleDeclaration) ImeMode() (imeMode string) {
	macro.Rewrite("$_.imeMode")
	return imeMode
}

func (*CSSStyleDeclaration) SetImeMode(imeMode string) {
	macro.Rewrite("$_.imeMode = $1", imeMode)
}

func (*CSSStyleDeclaration) JustifyContent() (justifyContent string) {
	macro.Rewrite("$_.justifyContent")
	return justifyContent
}

func (*CSSStyleDeclaration) SetJustifyContent(justifyContent string) {
	macro.Rewrite("$_.justifyContent = $1", justifyContent)
}

func (*CSSStyleDeclaration) Kerning() (kerning string) {
	macro.Rewrite("$_.kerning")
	return kerning
}

func (*CSSStyleDeclaration) SetKerning(kerning string) {
	macro.Rewrite("$_.kerning = $1", kerning)
}

func (*CSSStyleDeclaration) LayoutGrid() (layoutGrid string) {
	macro.Rewrite("$_.layoutGrid")
	return layoutGrid
}

func (*CSSStyleDeclaration) SetLayoutGrid(layoutGrid string) {
	macro.Rewrite("$_.layoutGrid = $1", layoutGrid)
}

func (*CSSStyleDeclaration) LayoutGridChar() (layoutGridChar string) {
	macro.Rewrite("$_.layoutGridChar")
	return layoutGridChar
}

func (*CSSStyleDeclaration) SetLayoutGridChar(layoutGridChar string) {
	macro.Rewrite("$_.layoutGridChar = $1", layoutGridChar)
}

func (*CSSStyleDeclaration) LayoutGridLine() (layoutGridLine string) {
	macro.Rewrite("$_.layoutGridLine")
	return layoutGridLine
}

func (*CSSStyleDeclaration) SetLayoutGridLine(layoutGridLine string) {
	macro.Rewrite("$_.layoutGridLine = $1", layoutGridLine)
}

func (*CSSStyleDeclaration) LayoutGridMode() (layoutGridMode string) {
	macro.Rewrite("$_.layoutGridMode")
	return layoutGridMode
}

func (*CSSStyleDeclaration) SetLayoutGridMode(layoutGridMode string) {
	macro.Rewrite("$_.layoutGridMode = $1", layoutGridMode)
}

func (*CSSStyleDeclaration) LayoutGridType() (layoutGridType string) {
	macro.Rewrite("$_.layoutGridType")
	return layoutGridType
}

func (*CSSStyleDeclaration) SetLayoutGridType(layoutGridType string) {
	macro.Rewrite("$_.layoutGridType = $1", layoutGridType)
}

func (*CSSStyleDeclaration) Left() (left string) {
	macro.Rewrite("$_.left")
	return left
}

func (*CSSStyleDeclaration) SetLeft(left string) {
	macro.Rewrite("$_.left = $1", left)
}

func (*CSSStyleDeclaration) Length() (length uint) {
	macro.Rewrite("$_.length")
	return length
}

func (*CSSStyleDeclaration) LetterSpacing() (letterSpacing string) {
	macro.Rewrite("$_.letterSpacing")
	return letterSpacing
}

func (*CSSStyleDeclaration) SetLetterSpacing(letterSpacing string) {
	macro.Rewrite("$_.letterSpacing = $1", letterSpacing)
}

func (*CSSStyleDeclaration) LightingColor() (lightingColor string) {
	macro.Rewrite("$_.lightingColor")
	return lightingColor
}

func (*CSSStyleDeclaration) SetLightingColor(lightingColor string) {
	macro.Rewrite("$_.lightingColor = $1", lightingColor)
}

func (*CSSStyleDeclaration) LineBreak() (lineBreak string) {
	macro.Rewrite("$_.lineBreak")
	return lineBreak
}

func (*CSSStyleDeclaration) SetLineBreak(lineBreak string) {
	macro.Rewrite("$_.lineBreak = $1", lineBreak)
}

func (*CSSStyleDeclaration) LineHeight() (lineHeight string) {
	macro.Rewrite("$_.lineHeight")
	return lineHeight
}

func (*CSSStyleDeclaration) SetLineHeight(lineHeight string) {
	macro.Rewrite("$_.lineHeight = $1", lineHeight)
}

func (*CSSStyleDeclaration) ListStyle() (listStyle string) {
	macro.Rewrite("$_.listStyle")
	return listStyle
}

func (*CSSStyleDeclaration) SetListStyle(listStyle string) {
	macro.Rewrite("$_.listStyle = $1", listStyle)
}

func (*CSSStyleDeclaration) ListStyleImage() (listStyleImage string) {
	macro.Rewrite("$_.listStyleImage")
	return listStyleImage
}

func (*CSSStyleDeclaration) SetListStyleImage(listStyleImage string) {
	macro.Rewrite("$_.listStyleImage = $1", listStyleImage)
}

func (*CSSStyleDeclaration) ListStylePosition() (listStylePosition string) {
	macro.Rewrite("$_.listStylePosition")
	return listStylePosition
}

func (*CSSStyleDeclaration) SetListStylePosition(listStylePosition string) {
	macro.Rewrite("$_.listStylePosition = $1", listStylePosition)
}

func (*CSSStyleDeclaration) ListStyleType() (listStyleType string) {
	macro.Rewrite("$_.listStyleType")
	return listStyleType
}

func (*CSSStyleDeclaration) SetListStyleType(listStyleType string) {
	macro.Rewrite("$_.listStyleType = $1", listStyleType)
}

func (*CSSStyleDeclaration) Margin() (margin string) {
	macro.Rewrite("$_.margin")
	return margin
}

func (*CSSStyleDeclaration) SetMargin(margin string) {
	macro.Rewrite("$_.margin = $1", margin)
}

func (*CSSStyleDeclaration) MarginBottom() (marginBottom string) {
	macro.Rewrite("$_.marginBottom")
	return marginBottom
}

func (*CSSStyleDeclaration) SetMarginBottom(marginBottom string) {
	macro.Rewrite("$_.marginBottom = $1", marginBottom)
}

func (*CSSStyleDeclaration) MarginLeft() (marginLeft string) {
	macro.Rewrite("$_.marginLeft")
	return marginLeft
}

func (*CSSStyleDeclaration) SetMarginLeft(marginLeft string) {
	macro.Rewrite("$_.marginLeft = $1", marginLeft)
}

func (*CSSStyleDeclaration) MarginRight() (marginRight string) {
	macro.Rewrite("$_.marginRight")
	return marginRight
}

func (*CSSStyleDeclaration) SetMarginRight(marginRight string) {
	macro.Rewrite("$_.marginRight = $1", marginRight)
}

func (*CSSStyleDeclaration) MarginTop() (marginTop string) {
	macro.Rewrite("$_.marginTop")
	return marginTop
}

func (*CSSStyleDeclaration) SetMarginTop(marginTop string) {
	macro.Rewrite("$_.marginTop = $1", marginTop)
}

func (*CSSStyleDeclaration) Marker() (marker string) {
	macro.Rewrite("$_.marker")
	return marker
}

func (*CSSStyleDeclaration) SetMarker(marker string) {
	macro.Rewrite("$_.marker = $1", marker)
}

func (*CSSStyleDeclaration) MarkerEnd() (markerEnd string) {
	macro.Rewrite("$_.markerEnd")
	return markerEnd
}

func (*CSSStyleDeclaration) SetMarkerEnd(markerEnd string) {
	macro.Rewrite("$_.markerEnd = $1", markerEnd)
}

func (*CSSStyleDeclaration) MarkerMid() (markerMid string) {
	macro.Rewrite("$_.markerMid")
	return markerMid
}

func (*CSSStyleDeclaration) SetMarkerMid(markerMid string) {
	macro.Rewrite("$_.markerMid = $1", markerMid)
}

func (*CSSStyleDeclaration) MarkerStart() (markerStart string) {
	macro.Rewrite("$_.markerStart")
	return markerStart
}

func (*CSSStyleDeclaration) SetMarkerStart(markerStart string) {
	macro.Rewrite("$_.markerStart = $1", markerStart)
}

func (*CSSStyleDeclaration) Mask() (mask string) {
	macro.Rewrite("$_.mask")
	return mask
}

func (*CSSStyleDeclaration) SetMask(mask string) {
	macro.Rewrite("$_.mask = $1", mask)
}

func (*CSSStyleDeclaration) MaxHeight() (maxHeight string) {
	macro.Rewrite("$_.maxHeight")
	return maxHeight
}

func (*CSSStyleDeclaration) SetMaxHeight(maxHeight string) {
	macro.Rewrite("$_.maxHeight = $1", maxHeight)
}

func (*CSSStyleDeclaration) MaxWidth() (maxWidth string) {
	macro.Rewrite("$_.maxWidth")
	return maxWidth
}

func (*CSSStyleDeclaration) SetMaxWidth(maxWidth string) {
	macro.Rewrite("$_.maxWidth = $1", maxWidth)
}

func (*CSSStyleDeclaration) MinHeight() (minHeight string) {
	macro.Rewrite("$_.minHeight")
	return minHeight
}

func (*CSSStyleDeclaration) SetMinHeight(minHeight string) {
	macro.Rewrite("$_.minHeight = $1", minHeight)
}

func (*CSSStyleDeclaration) MinWidth() (minWidth string) {
	macro.Rewrite("$_.minWidth")
	return minWidth
}

func (*CSSStyleDeclaration) SetMinWidth(minWidth string) {
	macro.Rewrite("$_.minWidth = $1", minWidth)
}

func (*CSSStyleDeclaration) MsContentZoomChaining() (msContentZoomChaining string) {
	macro.Rewrite("$_.msContentZoomChaining")
	return msContentZoomChaining
}

func (*CSSStyleDeclaration) SetMsContentZoomChaining(msContentZoomChaining string) {
	macro.Rewrite("$_.msContentZoomChaining = $1", msContentZoomChaining)
}

func (*CSSStyleDeclaration) MsContentZooming() (msContentZooming string) {
	macro.Rewrite("$_.msContentZooming")
	return msContentZooming
}

func (*CSSStyleDeclaration) SetMsContentZooming(msContentZooming string) {
	macro.Rewrite("$_.msContentZooming = $1", msContentZooming)
}

func (*CSSStyleDeclaration) MsContentZoomLimit() (msContentZoomLimit string) {
	macro.Rewrite("$_.msContentZoomLimit")
	return msContentZoomLimit
}

func (*CSSStyleDeclaration) SetMsContentZoomLimit(msContentZoomLimit string) {
	macro.Rewrite("$_.msContentZoomLimit = $1", msContentZoomLimit)
}

func (*CSSStyleDeclaration) MsContentZoomLimitMax() (msContentZoomLimitMax interface{}) {
	macro.Rewrite("$_.msContentZoomLimitMax")
	return msContentZoomLimitMax
}

func (*CSSStyleDeclaration) SetMsContentZoomLimitMax(msContentZoomLimitMax interface{}) {
	macro.Rewrite("$_.msContentZoomLimitMax = $1", msContentZoomLimitMax)
}

func (*CSSStyleDeclaration) MsContentZoomLimitMin() (msContentZoomLimitMin interface{}) {
	macro.Rewrite("$_.msContentZoomLimitMin")
	return msContentZoomLimitMin
}

func (*CSSStyleDeclaration) SetMsContentZoomLimitMin(msContentZoomLimitMin interface{}) {
	macro.Rewrite("$_.msContentZoomLimitMin = $1", msContentZoomLimitMin)
}

func (*CSSStyleDeclaration) MsContentZoomSnap() (msContentZoomSnap string) {
	macro.Rewrite("$_.msContentZoomSnap")
	return msContentZoomSnap
}

func (*CSSStyleDeclaration) SetMsContentZoomSnap(msContentZoomSnap string) {
	macro.Rewrite("$_.msContentZoomSnap = $1", msContentZoomSnap)
}

func (*CSSStyleDeclaration) MsContentZoomSnapPoints() (msContentZoomSnapPoints string) {
	macro.Rewrite("$_.msContentZoomSnapPoints")
	return msContentZoomSnapPoints
}

func (*CSSStyleDeclaration) SetMsContentZoomSnapPoints(msContentZoomSnapPoints string) {
	macro.Rewrite("$_.msContentZoomSnapPoints = $1", msContentZoomSnapPoints)
}

func (*CSSStyleDeclaration) MsContentZoomSnapType() (msContentZoomSnapType string) {
	macro.Rewrite("$_.msContentZoomSnapType")
	return msContentZoomSnapType
}

func (*CSSStyleDeclaration) SetMsContentZoomSnapType(msContentZoomSnapType string) {
	macro.Rewrite("$_.msContentZoomSnapType = $1", msContentZoomSnapType)
}

func (*CSSStyleDeclaration) MsFlowFrom() (msFlowFrom string) {
	macro.Rewrite("$_.msFlowFrom")
	return msFlowFrom
}

func (*CSSStyleDeclaration) SetMsFlowFrom(msFlowFrom string) {
	macro.Rewrite("$_.msFlowFrom = $1", msFlowFrom)
}

func (*CSSStyleDeclaration) MsFlowInto() (msFlowInto string) {
	macro.Rewrite("$_.msFlowInto")
	return msFlowInto
}

func (*CSSStyleDeclaration) SetMsFlowInto(msFlowInto string) {
	macro.Rewrite("$_.msFlowInto = $1", msFlowInto)
}

func (*CSSStyleDeclaration) MsFontFeatureSettings() (msFontFeatureSettings string) {
	macro.Rewrite("$_.msFontFeatureSettings")
	return msFontFeatureSettings
}

func (*CSSStyleDeclaration) SetMsFontFeatureSettings(msFontFeatureSettings string) {
	macro.Rewrite("$_.msFontFeatureSettings = $1", msFontFeatureSettings)
}

func (*CSSStyleDeclaration) MsGridColumn() (msGridColumn interface{}) {
	macro.Rewrite("$_.msGridColumn")
	return msGridColumn
}

func (*CSSStyleDeclaration) SetMsGridColumn(msGridColumn interface{}) {
	macro.Rewrite("$_.msGridColumn = $1", msGridColumn)
}

func (*CSSStyleDeclaration) MsGridColumnAlign() (msGridColumnAlign string) {
	macro.Rewrite("$_.msGridColumnAlign")
	return msGridColumnAlign
}

func (*CSSStyleDeclaration) SetMsGridColumnAlign(msGridColumnAlign string) {
	macro.Rewrite("$_.msGridColumnAlign = $1", msGridColumnAlign)
}

func (*CSSStyleDeclaration) MsGridColumns() (msGridColumns string) {
	macro.Rewrite("$_.msGridColumns")
	return msGridColumns
}

func (*CSSStyleDeclaration) SetMsGridColumns(msGridColumns string) {
	macro.Rewrite("$_.msGridColumns = $1", msGridColumns)
}

func (*CSSStyleDeclaration) MsGridColumnSpan() (msGridColumnSpan interface{}) {
	macro.Rewrite("$_.msGridColumnSpan")
	return msGridColumnSpan
}

func (*CSSStyleDeclaration) SetMsGridColumnSpan(msGridColumnSpan interface{}) {
	macro.Rewrite("$_.msGridColumnSpan = $1", msGridColumnSpan)
}

func (*CSSStyleDeclaration) MsGridRow() (msGridRow interface{}) {
	macro.Rewrite("$_.msGridRow")
	return msGridRow
}

func (*CSSStyleDeclaration) SetMsGridRow(msGridRow interface{}) {
	macro.Rewrite("$_.msGridRow = $1", msGridRow)
}

func (*CSSStyleDeclaration) MsGridRowAlign() (msGridRowAlign string) {
	macro.Rewrite("$_.msGridRowAlign")
	return msGridRowAlign
}

func (*CSSStyleDeclaration) SetMsGridRowAlign(msGridRowAlign string) {
	macro.Rewrite("$_.msGridRowAlign = $1", msGridRowAlign)
}

func (*CSSStyleDeclaration) MsGridRows() (msGridRows string) {
	macro.Rewrite("$_.msGridRows")
	return msGridRows
}

func (*CSSStyleDeclaration) SetMsGridRows(msGridRows string) {
	macro.Rewrite("$_.msGridRows = $1", msGridRows)
}

func (*CSSStyleDeclaration) MsGridRowSpan() (msGridRowSpan interface{}) {
	macro.Rewrite("$_.msGridRowSpan")
	return msGridRowSpan
}

func (*CSSStyleDeclaration) SetMsGridRowSpan(msGridRowSpan interface{}) {
	macro.Rewrite("$_.msGridRowSpan = $1", msGridRowSpan)
}

func (*CSSStyleDeclaration) MsHighContrastAdjust() (msHighContrastAdjust string) {
	macro.Rewrite("$_.msHighContrastAdjust")
	return msHighContrastAdjust
}

func (*CSSStyleDeclaration) SetMsHighContrastAdjust(msHighContrastAdjust string) {
	macro.Rewrite("$_.msHighContrastAdjust = $1", msHighContrastAdjust)
}

func (*CSSStyleDeclaration) MsHyphenateLimitChars() (msHyphenateLimitChars string) {
	macro.Rewrite("$_.msHyphenateLimitChars")
	return msHyphenateLimitChars
}

func (*CSSStyleDeclaration) SetMsHyphenateLimitChars(msHyphenateLimitChars string) {
	macro.Rewrite("$_.msHyphenateLimitChars = $1", msHyphenateLimitChars)
}

func (*CSSStyleDeclaration) MsHyphenateLimitLines() (msHyphenateLimitLines interface{}) {
	macro.Rewrite("$_.msHyphenateLimitLines")
	return msHyphenateLimitLines
}

func (*CSSStyleDeclaration) SetMsHyphenateLimitLines(msHyphenateLimitLines interface{}) {
	macro.Rewrite("$_.msHyphenateLimitLines = $1", msHyphenateLimitLines)
}

func (*CSSStyleDeclaration) MsHyphenateLimitZone() (msHyphenateLimitZone interface{}) {
	macro.Rewrite("$_.msHyphenateLimitZone")
	return msHyphenateLimitZone
}

func (*CSSStyleDeclaration) SetMsHyphenateLimitZone(msHyphenateLimitZone interface{}) {
	macro.Rewrite("$_.msHyphenateLimitZone = $1", msHyphenateLimitZone)
}

func (*CSSStyleDeclaration) MsHyphens() (msHyphens string) {
	macro.Rewrite("$_.msHyphens")
	return msHyphens
}

func (*CSSStyleDeclaration) SetMsHyphens(msHyphens string) {
	macro.Rewrite("$_.msHyphens = $1", msHyphens)
}

func (*CSSStyleDeclaration) MsImeAlign() (msImeAlign string) {
	macro.Rewrite("$_.msImeAlign")
	return msImeAlign
}

func (*CSSStyleDeclaration) SetMsImeAlign(msImeAlign string) {
	macro.Rewrite("$_.msImeAlign = $1", msImeAlign)
}

func (*CSSStyleDeclaration) MsOverflowStyle() (msOverflowStyle string) {
	macro.Rewrite("$_.msOverflowStyle")
	return msOverflowStyle
}

func (*CSSStyleDeclaration) SetMsOverflowStyle(msOverflowStyle string) {
	macro.Rewrite("$_.msOverflowStyle = $1", msOverflowStyle)
}

func (*CSSStyleDeclaration) MsScrollChaining() (msScrollChaining string) {
	macro.Rewrite("$_.msScrollChaining")
	return msScrollChaining
}

func (*CSSStyleDeclaration) SetMsScrollChaining(msScrollChaining string) {
	macro.Rewrite("$_.msScrollChaining = $1", msScrollChaining)
}

func (*CSSStyleDeclaration) MsScrollLimit() (msScrollLimit string) {
	macro.Rewrite("$_.msScrollLimit")
	return msScrollLimit
}

func (*CSSStyleDeclaration) SetMsScrollLimit(msScrollLimit string) {
	macro.Rewrite("$_.msScrollLimit = $1", msScrollLimit)
}

func (*CSSStyleDeclaration) MsScrollLimitXMax() (msScrollLimitXMax interface{}) {
	macro.Rewrite("$_.msScrollLimitXMax")
	return msScrollLimitXMax
}

func (*CSSStyleDeclaration) SetMsScrollLimitXMax(msScrollLimitXMax interface{}) {
	macro.Rewrite("$_.msScrollLimitXMax = $1", msScrollLimitXMax)
}

func (*CSSStyleDeclaration) MsScrollLimitXMin() (msScrollLimitXMin interface{}) {
	macro.Rewrite("$_.msScrollLimitXMin")
	return msScrollLimitXMin
}

func (*CSSStyleDeclaration) SetMsScrollLimitXMin(msScrollLimitXMin interface{}) {
	macro.Rewrite("$_.msScrollLimitXMin = $1", msScrollLimitXMin)
}

func (*CSSStyleDeclaration) MsScrollLimitYMax() (msScrollLimitYMax interface{}) {
	macro.Rewrite("$_.msScrollLimitYMax")
	return msScrollLimitYMax
}

func (*CSSStyleDeclaration) SetMsScrollLimitYMax(msScrollLimitYMax interface{}) {
	macro.Rewrite("$_.msScrollLimitYMax = $1", msScrollLimitYMax)
}

func (*CSSStyleDeclaration) MsScrollLimitYMin() (msScrollLimitYMin interface{}) {
	macro.Rewrite("$_.msScrollLimitYMin")
	return msScrollLimitYMin
}

func (*CSSStyleDeclaration) SetMsScrollLimitYMin(msScrollLimitYMin interface{}) {
	macro.Rewrite("$_.msScrollLimitYMin = $1", msScrollLimitYMin)
}

func (*CSSStyleDeclaration) MsScrollRails() (msScrollRails string) {
	macro.Rewrite("$_.msScrollRails")
	return msScrollRails
}

func (*CSSStyleDeclaration) SetMsScrollRails(msScrollRails string) {
	macro.Rewrite("$_.msScrollRails = $1", msScrollRails)
}

func (*CSSStyleDeclaration) MsScrollSnapPointsX() (msScrollSnapPointsX string) {
	macro.Rewrite("$_.msScrollSnapPointsX")
	return msScrollSnapPointsX
}

func (*CSSStyleDeclaration) SetMsScrollSnapPointsX(msScrollSnapPointsX string) {
	macro.Rewrite("$_.msScrollSnapPointsX = $1", msScrollSnapPointsX)
}

func (*CSSStyleDeclaration) MsScrollSnapPointsY() (msScrollSnapPointsY string) {
	macro.Rewrite("$_.msScrollSnapPointsY")
	return msScrollSnapPointsY
}

func (*CSSStyleDeclaration) SetMsScrollSnapPointsY(msScrollSnapPointsY string) {
	macro.Rewrite("$_.msScrollSnapPointsY = $1", msScrollSnapPointsY)
}

func (*CSSStyleDeclaration) MsScrollSnapType() (msScrollSnapType string) {
	macro.Rewrite("$_.msScrollSnapType")
	return msScrollSnapType
}

func (*CSSStyleDeclaration) SetMsScrollSnapType(msScrollSnapType string) {
	macro.Rewrite("$_.msScrollSnapType = $1", msScrollSnapType)
}

func (*CSSStyleDeclaration) MsScrollSnapX() (msScrollSnapX string) {
	macro.Rewrite("$_.msScrollSnapX")
	return msScrollSnapX
}

func (*CSSStyleDeclaration) SetMsScrollSnapX(msScrollSnapX string) {
	macro.Rewrite("$_.msScrollSnapX = $1", msScrollSnapX)
}

func (*CSSStyleDeclaration) MsScrollSnapY() (msScrollSnapY string) {
	macro.Rewrite("$_.msScrollSnapY")
	return msScrollSnapY
}

func (*CSSStyleDeclaration) SetMsScrollSnapY(msScrollSnapY string) {
	macro.Rewrite("$_.msScrollSnapY = $1", msScrollSnapY)
}

func (*CSSStyleDeclaration) MsScrollTranslation() (msScrollTranslation string) {
	macro.Rewrite("$_.msScrollTranslation")
	return msScrollTranslation
}

func (*CSSStyleDeclaration) SetMsScrollTranslation(msScrollTranslation string) {
	macro.Rewrite("$_.msScrollTranslation = $1", msScrollTranslation)
}

func (*CSSStyleDeclaration) MsTextCombineHorizontal() (msTextCombineHorizontal string) {
	macro.Rewrite("$_.msTextCombineHorizontal")
	return msTextCombineHorizontal
}

func (*CSSStyleDeclaration) SetMsTextCombineHorizontal(msTextCombineHorizontal string) {
	macro.Rewrite("$_.msTextCombineHorizontal = $1", msTextCombineHorizontal)
}

func (*CSSStyleDeclaration) MsTextSizeAdjust() (msTextSizeAdjust interface{}) {
	macro.Rewrite("$_.msTextSizeAdjust")
	return msTextSizeAdjust
}

func (*CSSStyleDeclaration) SetMsTextSizeAdjust(msTextSizeAdjust interface{}) {
	macro.Rewrite("$_.msTextSizeAdjust = $1", msTextSizeAdjust)
}

func (*CSSStyleDeclaration) MsTouchAction() (msTouchAction string) {
	macro.Rewrite("$_.msTouchAction")
	return msTouchAction
}

func (*CSSStyleDeclaration) SetMsTouchAction(msTouchAction string) {
	macro.Rewrite("$_.msTouchAction = $1", msTouchAction)
}

func (*CSSStyleDeclaration) MsTouchSelect() (msTouchSelect string) {
	macro.Rewrite("$_.msTouchSelect")
	return msTouchSelect
}

func (*CSSStyleDeclaration) SetMsTouchSelect(msTouchSelect string) {
	macro.Rewrite("$_.msTouchSelect = $1", msTouchSelect)
}

func (*CSSStyleDeclaration) MsUserSelect() (msUserSelect string) {
	macro.Rewrite("$_.msUserSelect")
	return msUserSelect
}

func (*CSSStyleDeclaration) SetMsUserSelect(msUserSelect string) {
	macro.Rewrite("$_.msUserSelect = $1", msUserSelect)
}

func (*CSSStyleDeclaration) MsWrapFlow() (msWrapFlow string) {
	macro.Rewrite("$_.msWrapFlow")
	return msWrapFlow
}

func (*CSSStyleDeclaration) SetMsWrapFlow(msWrapFlow string) {
	macro.Rewrite("$_.msWrapFlow = $1", msWrapFlow)
}

func (*CSSStyleDeclaration) MsWrapMargin() (msWrapMargin interface{}) {
	macro.Rewrite("$_.msWrapMargin")
	return msWrapMargin
}

func (*CSSStyleDeclaration) SetMsWrapMargin(msWrapMargin interface{}) {
	macro.Rewrite("$_.msWrapMargin = $1", msWrapMargin)
}

func (*CSSStyleDeclaration) MsWrapThrough() (msWrapThrough string) {
	macro.Rewrite("$_.msWrapThrough")
	return msWrapThrough
}

func (*CSSStyleDeclaration) SetMsWrapThrough(msWrapThrough string) {
	macro.Rewrite("$_.msWrapThrough = $1", msWrapThrough)
}

func (*CSSStyleDeclaration) Opacity() (opacity string) {
	macro.Rewrite("$_.opacity")
	return opacity
}

func (*CSSStyleDeclaration) SetOpacity(opacity string) {
	macro.Rewrite("$_.opacity = $1", opacity)
}

func (*CSSStyleDeclaration) Order() (order string) {
	macro.Rewrite("$_.order")
	return order
}

func (*CSSStyleDeclaration) SetOrder(order string) {
	macro.Rewrite("$_.order = $1", order)
}

func (*CSSStyleDeclaration) Orphans() (orphans string) {
	macro.Rewrite("$_.orphans")
	return orphans
}

func (*CSSStyleDeclaration) SetOrphans(orphans string) {
	macro.Rewrite("$_.orphans = $1", orphans)
}

func (*CSSStyleDeclaration) Outline() (outline string) {
	macro.Rewrite("$_.outline")
	return outline
}

func (*CSSStyleDeclaration) SetOutline(outline string) {
	macro.Rewrite("$_.outline = $1", outline)
}

func (*CSSStyleDeclaration) OutlineColor() (outlineColor string) {
	macro.Rewrite("$_.outlineColor")
	return outlineColor
}

func (*CSSStyleDeclaration) SetOutlineColor(outlineColor string) {
	macro.Rewrite("$_.outlineColor = $1", outlineColor)
}

func (*CSSStyleDeclaration) OutlineOffset() (outlineOffset string) {
	macro.Rewrite("$_.outlineOffset")
	return outlineOffset
}

func (*CSSStyleDeclaration) SetOutlineOffset(outlineOffset string) {
	macro.Rewrite("$_.outlineOffset = $1", outlineOffset)
}

func (*CSSStyleDeclaration) OutlineStyle() (outlineStyle string) {
	macro.Rewrite("$_.outlineStyle")
	return outlineStyle
}

func (*CSSStyleDeclaration) SetOutlineStyle(outlineStyle string) {
	macro.Rewrite("$_.outlineStyle = $1", outlineStyle)
}

func (*CSSStyleDeclaration) OutlineWidth() (outlineWidth string) {
	macro.Rewrite("$_.outlineWidth")
	return outlineWidth
}

func (*CSSStyleDeclaration) SetOutlineWidth(outlineWidth string) {
	macro.Rewrite("$_.outlineWidth = $1", outlineWidth)
}

func (*CSSStyleDeclaration) Overflow() (overflow string) {
	macro.Rewrite("$_.overflow")
	return overflow
}

func (*CSSStyleDeclaration) SetOverflow(overflow string) {
	macro.Rewrite("$_.overflow = $1", overflow)
}

func (*CSSStyleDeclaration) OverflowX() (overflowX string) {
	macro.Rewrite("$_.overflowX")
	return overflowX
}

func (*CSSStyleDeclaration) SetOverflowX(overflowX string) {
	macro.Rewrite("$_.overflowX = $1", overflowX)
}

func (*CSSStyleDeclaration) OverflowY() (overflowY string) {
	macro.Rewrite("$_.overflowY")
	return overflowY
}

func (*CSSStyleDeclaration) SetOverflowY(overflowY string) {
	macro.Rewrite("$_.overflowY = $1", overflowY)
}

func (*CSSStyleDeclaration) Padding() (padding string) {
	macro.Rewrite("$_.padding")
	return padding
}

func (*CSSStyleDeclaration) SetPadding(padding string) {
	macro.Rewrite("$_.padding = $1", padding)
}

func (*CSSStyleDeclaration) PaddingBottom() (paddingBottom string) {
	macro.Rewrite("$_.paddingBottom")
	return paddingBottom
}

func (*CSSStyleDeclaration) SetPaddingBottom(paddingBottom string) {
	macro.Rewrite("$_.paddingBottom = $1", paddingBottom)
}

func (*CSSStyleDeclaration) PaddingLeft() (paddingLeft string) {
	macro.Rewrite("$_.paddingLeft")
	return paddingLeft
}

func (*CSSStyleDeclaration) SetPaddingLeft(paddingLeft string) {
	macro.Rewrite("$_.paddingLeft = $1", paddingLeft)
}

func (*CSSStyleDeclaration) PaddingRight() (paddingRight string) {
	macro.Rewrite("$_.paddingRight")
	return paddingRight
}

func (*CSSStyleDeclaration) SetPaddingRight(paddingRight string) {
	macro.Rewrite("$_.paddingRight = $1", paddingRight)
}

func (*CSSStyleDeclaration) PaddingTop() (paddingTop string) {
	macro.Rewrite("$_.paddingTop")
	return paddingTop
}

func (*CSSStyleDeclaration) SetPaddingTop(paddingTop string) {
	macro.Rewrite("$_.paddingTop = $1", paddingTop)
}

func (*CSSStyleDeclaration) PageBreakAfter() (pageBreakAfter string) {
	macro.Rewrite("$_.pageBreakAfter")
	return pageBreakAfter
}

func (*CSSStyleDeclaration) SetPageBreakAfter(pageBreakAfter string) {
	macro.Rewrite("$_.pageBreakAfter = $1", pageBreakAfter)
}

func (*CSSStyleDeclaration) PageBreakBefore() (pageBreakBefore string) {
	macro.Rewrite("$_.pageBreakBefore")
	return pageBreakBefore
}

func (*CSSStyleDeclaration) SetPageBreakBefore(pageBreakBefore string) {
	macro.Rewrite("$_.pageBreakBefore = $1", pageBreakBefore)
}

func (*CSSStyleDeclaration) PageBreakInside() (pageBreakInside string) {
	macro.Rewrite("$_.pageBreakInside")
	return pageBreakInside
}

func (*CSSStyleDeclaration) SetPageBreakInside(pageBreakInside string) {
	macro.Rewrite("$_.pageBreakInside = $1", pageBreakInside)
}

func (*CSSStyleDeclaration) ParentRule() (parentRule CSSRule) {
	macro.Rewrite("$_.parentRule")
	return parentRule
}

func (*CSSStyleDeclaration) Perspective() (perspective string) {
	macro.Rewrite("$_.perspective")
	return perspective
}

func (*CSSStyleDeclaration) SetPerspective(perspective string) {
	macro.Rewrite("$_.perspective = $1", perspective)
}

func (*CSSStyleDeclaration) PerspectiveOrigin() (perspectiveOrigin string) {
	macro.Rewrite("$_.perspectiveOrigin")
	return perspectiveOrigin
}

func (*CSSStyleDeclaration) SetPerspectiveOrigin(perspectiveOrigin string) {
	macro.Rewrite("$_.perspectiveOrigin = $1", perspectiveOrigin)
}

func (*CSSStyleDeclaration) PointerEvents() (pointerEvents string) {
	macro.Rewrite("$_.pointerEvents")
	return pointerEvents
}

func (*CSSStyleDeclaration) SetPointerEvents(pointerEvents string) {
	macro.Rewrite("$_.pointerEvents = $1", pointerEvents)
}

func (*CSSStyleDeclaration) Position() (position string) {
	macro.Rewrite("$_.position")
	return position
}

func (*CSSStyleDeclaration) SetPosition(position string) {
	macro.Rewrite("$_.position = $1", position)
}

func (*CSSStyleDeclaration) Quotes() (quotes string) {
	macro.Rewrite("$_.quotes")
	return quotes
}

func (*CSSStyleDeclaration) SetQuotes(quotes string) {
	macro.Rewrite("$_.quotes = $1", quotes)
}

func (*CSSStyleDeclaration) Right() (right string) {
	macro.Rewrite("$_.right")
	return right
}

func (*CSSStyleDeclaration) SetRight(right string) {
	macro.Rewrite("$_.right = $1", right)
}

func (*CSSStyleDeclaration) Rotate() (rotate string) {
	macro.Rewrite("$_.rotate")
	return rotate
}

func (*CSSStyleDeclaration) SetRotate(rotate string) {
	macro.Rewrite("$_.rotate = $1", rotate)
}

func (*CSSStyleDeclaration) RubyAlign() (rubyAlign string) {
	macro.Rewrite("$_.rubyAlign")
	return rubyAlign
}

func (*CSSStyleDeclaration) SetRubyAlign(rubyAlign string) {
	macro.Rewrite("$_.rubyAlign = $1", rubyAlign)
}

func (*CSSStyleDeclaration) RubyOverhang() (rubyOverhang string) {
	macro.Rewrite("$_.rubyOverhang")
	return rubyOverhang
}

func (*CSSStyleDeclaration) SetRubyOverhang(rubyOverhang string) {
	macro.Rewrite("$_.rubyOverhang = $1", rubyOverhang)
}

func (*CSSStyleDeclaration) RubyPosition() (rubyPosition string) {
	macro.Rewrite("$_.rubyPosition")
	return rubyPosition
}

func (*CSSStyleDeclaration) SetRubyPosition(rubyPosition string) {
	macro.Rewrite("$_.rubyPosition = $1", rubyPosition)
}

func (*CSSStyleDeclaration) Scale() (scale string) {
	macro.Rewrite("$_.scale")
	return scale
}

func (*CSSStyleDeclaration) SetScale(scale string) {
	macro.Rewrite("$_.scale = $1", scale)
}

func (*CSSStyleDeclaration) StopColor() (stopColor string) {
	macro.Rewrite("$_.stopColor")
	return stopColor
}

func (*CSSStyleDeclaration) SetStopColor(stopColor string) {
	macro.Rewrite("$_.stopColor = $1", stopColor)
}

func (*CSSStyleDeclaration) StopOpacity() (stopOpacity string) {
	macro.Rewrite("$_.stopOpacity")
	return stopOpacity
}

func (*CSSStyleDeclaration) SetStopOpacity(stopOpacity string) {
	macro.Rewrite("$_.stopOpacity = $1", stopOpacity)
}

func (*CSSStyleDeclaration) Stroke() (stroke string) {
	macro.Rewrite("$_.stroke")
	return stroke
}

func (*CSSStyleDeclaration) SetStroke(stroke string) {
	macro.Rewrite("$_.stroke = $1", stroke)
}

func (*CSSStyleDeclaration) StrokeDasharray() (strokeDasharray string) {
	macro.Rewrite("$_.strokeDasharray")
	return strokeDasharray
}

func (*CSSStyleDeclaration) SetStrokeDasharray(strokeDasharray string) {
	macro.Rewrite("$_.strokeDasharray = $1", strokeDasharray)
}

func (*CSSStyleDeclaration) StrokeDashoffset() (strokeDashoffset string) {
	macro.Rewrite("$_.strokeDashoffset")
	return strokeDashoffset
}

func (*CSSStyleDeclaration) SetStrokeDashoffset(strokeDashoffset string) {
	macro.Rewrite("$_.strokeDashoffset = $1", strokeDashoffset)
}

func (*CSSStyleDeclaration) StrokeLinecap() (strokeLinecap string) {
	macro.Rewrite("$_.strokeLinecap")
	return strokeLinecap
}

func (*CSSStyleDeclaration) SetStrokeLinecap(strokeLinecap string) {
	macro.Rewrite("$_.strokeLinecap = $1", strokeLinecap)
}

func (*CSSStyleDeclaration) StrokeLinejoin() (strokeLinejoin string) {
	macro.Rewrite("$_.strokeLinejoin")
	return strokeLinejoin
}

func (*CSSStyleDeclaration) SetStrokeLinejoin(strokeLinejoin string) {
	macro.Rewrite("$_.strokeLinejoin = $1", strokeLinejoin)
}

func (*CSSStyleDeclaration) StrokeMiterlimit() (strokeMiterlimit string) {
	macro.Rewrite("$_.strokeMiterlimit")
	return strokeMiterlimit
}

func (*CSSStyleDeclaration) SetStrokeMiterlimit(strokeMiterlimit string) {
	macro.Rewrite("$_.strokeMiterlimit = $1", strokeMiterlimit)
}

func (*CSSStyleDeclaration) StrokeOpacity() (strokeOpacity string) {
	macro.Rewrite("$_.strokeOpacity")
	return strokeOpacity
}

func (*CSSStyleDeclaration) SetStrokeOpacity(strokeOpacity string) {
	macro.Rewrite("$_.strokeOpacity = $1", strokeOpacity)
}

func (*CSSStyleDeclaration) StrokeWidth() (strokeWidth string) {
	macro.Rewrite("$_.strokeWidth")
	return strokeWidth
}

func (*CSSStyleDeclaration) SetStrokeWidth(strokeWidth string) {
	macro.Rewrite("$_.strokeWidth = $1", strokeWidth)
}

func (*CSSStyleDeclaration) TableLayout() (tableLayout string) {
	macro.Rewrite("$_.tableLayout")
	return tableLayout
}

func (*CSSStyleDeclaration) SetTableLayout(tableLayout string) {
	macro.Rewrite("$_.tableLayout = $1", tableLayout)
}

func (*CSSStyleDeclaration) TextAlign() (textAlign string) {
	macro.Rewrite("$_.textAlign")
	return textAlign
}

func (*CSSStyleDeclaration) SetTextAlign(textAlign string) {
	macro.Rewrite("$_.textAlign = $1", textAlign)
}

func (*CSSStyleDeclaration) TextAlignLast() (textAlignLast string) {
	macro.Rewrite("$_.textAlignLast")
	return textAlignLast
}

func (*CSSStyleDeclaration) SetTextAlignLast(textAlignLast string) {
	macro.Rewrite("$_.textAlignLast = $1", textAlignLast)
}

func (*CSSStyleDeclaration) TextAnchor() (textAnchor string) {
	macro.Rewrite("$_.textAnchor")
	return textAnchor
}

func (*CSSStyleDeclaration) SetTextAnchor(textAnchor string) {
	macro.Rewrite("$_.textAnchor = $1", textAnchor)
}

func (*CSSStyleDeclaration) TextDecoration() (textDecoration string) {
	macro.Rewrite("$_.textDecoration")
	return textDecoration
}

func (*CSSStyleDeclaration) SetTextDecoration(textDecoration string) {
	macro.Rewrite("$_.textDecoration = $1", textDecoration)
}

func (*CSSStyleDeclaration) TextIndent() (textIndent string) {
	macro.Rewrite("$_.textIndent")
	return textIndent
}

func (*CSSStyleDeclaration) SetTextIndent(textIndent string) {
	macro.Rewrite("$_.textIndent = $1", textIndent)
}

func (*CSSStyleDeclaration) TextJustify() (textJustify string) {
	macro.Rewrite("$_.textJustify")
	return textJustify
}

func (*CSSStyleDeclaration) SetTextJustify(textJustify string) {
	macro.Rewrite("$_.textJustify = $1", textJustify)
}

func (*CSSStyleDeclaration) TextKashida() (textKashida string) {
	macro.Rewrite("$_.textKashida")
	return textKashida
}

func (*CSSStyleDeclaration) SetTextKashida(textKashida string) {
	macro.Rewrite("$_.textKashida = $1", textKashida)
}

func (*CSSStyleDeclaration) TextKashidaSpace() (textKashidaSpace string) {
	macro.Rewrite("$_.textKashidaSpace")
	return textKashidaSpace
}

func (*CSSStyleDeclaration) SetTextKashidaSpace(textKashidaSpace string) {
	macro.Rewrite("$_.textKashidaSpace = $1", textKashidaSpace)
}

func (*CSSStyleDeclaration) TextOverflow() (textOverflow string) {
	macro.Rewrite("$_.textOverflow")
	return textOverflow
}

func (*CSSStyleDeclaration) SetTextOverflow(textOverflow string) {
	macro.Rewrite("$_.textOverflow = $1", textOverflow)
}

func (*CSSStyleDeclaration) TextShadow() (textShadow string) {
	macro.Rewrite("$_.textShadow")
	return textShadow
}

func (*CSSStyleDeclaration) SetTextShadow(textShadow string) {
	macro.Rewrite("$_.textShadow = $1", textShadow)
}

func (*CSSStyleDeclaration) TextTransform() (textTransform string) {
	macro.Rewrite("$_.textTransform")
	return textTransform
}

func (*CSSStyleDeclaration) SetTextTransform(textTransform string) {
	macro.Rewrite("$_.textTransform = $1", textTransform)
}

func (*CSSStyleDeclaration) TextUnderlinePosition() (textUnderlinePosition string) {
	macro.Rewrite("$_.textUnderlinePosition")
	return textUnderlinePosition
}

func (*CSSStyleDeclaration) SetTextUnderlinePosition(textUnderlinePosition string) {
	macro.Rewrite("$_.textUnderlinePosition = $1", textUnderlinePosition)
}

func (*CSSStyleDeclaration) Top() (top string) {
	macro.Rewrite("$_.top")
	return top
}

func (*CSSStyleDeclaration) SetTop(top string) {
	macro.Rewrite("$_.top = $1", top)
}

func (*CSSStyleDeclaration) TouchAction() (touchAction string) {
	macro.Rewrite("$_.touchAction")
	return touchAction
}

func (*CSSStyleDeclaration) SetTouchAction(touchAction string) {
	macro.Rewrite("$_.touchAction = $1", touchAction)
}

func (*CSSStyleDeclaration) Transform() (transform string) {
	macro.Rewrite("$_.transform")
	return transform
}

func (*CSSStyleDeclaration) SetTransform(transform string) {
	macro.Rewrite("$_.transform = $1", transform)
}

func (*CSSStyleDeclaration) TransformOrigin() (transformOrigin string) {
	macro.Rewrite("$_.transformOrigin")
	return transformOrigin
}

func (*CSSStyleDeclaration) SetTransformOrigin(transformOrigin string) {
	macro.Rewrite("$_.transformOrigin = $1", transformOrigin)
}

func (*CSSStyleDeclaration) TransformStyle() (transformStyle string) {
	macro.Rewrite("$_.transformStyle")
	return transformStyle
}

func (*CSSStyleDeclaration) SetTransformStyle(transformStyle string) {
	macro.Rewrite("$_.transformStyle = $1", transformStyle)
}

func (*CSSStyleDeclaration) Transition() (transition string) {
	macro.Rewrite("$_.transition")
	return transition
}

func (*CSSStyleDeclaration) SetTransition(transition string) {
	macro.Rewrite("$_.transition = $1", transition)
}

func (*CSSStyleDeclaration) TransitionDelay() (transitionDelay string) {
	macro.Rewrite("$_.transitionDelay")
	return transitionDelay
}

func (*CSSStyleDeclaration) SetTransitionDelay(transitionDelay string) {
	macro.Rewrite("$_.transitionDelay = $1", transitionDelay)
}

func (*CSSStyleDeclaration) TransitionDuration() (transitionDuration string) {
	macro.Rewrite("$_.transitionDuration")
	return transitionDuration
}

func (*CSSStyleDeclaration) SetTransitionDuration(transitionDuration string) {
	macro.Rewrite("$_.transitionDuration = $1", transitionDuration)
}

func (*CSSStyleDeclaration) TransitionProperty() (transitionProperty string) {
	macro.Rewrite("$_.transitionProperty")
	return transitionProperty
}

func (*CSSStyleDeclaration) SetTransitionProperty(transitionProperty string) {
	macro.Rewrite("$_.transitionProperty = $1", transitionProperty)
}

func (*CSSStyleDeclaration) TransitionTimingFunction() (transitionTimingFunction string) {
	macro.Rewrite("$_.transitionTimingFunction")
	return transitionTimingFunction
}

func (*CSSStyleDeclaration) SetTransitionTimingFunction(transitionTimingFunction string) {
	macro.Rewrite("$_.transitionTimingFunction = $1", transitionTimingFunction)
}

func (*CSSStyleDeclaration) Translate() (translate string) {
	macro.Rewrite("$_.translate")
	return translate
}

func (*CSSStyleDeclaration) SetTranslate(translate string) {
	macro.Rewrite("$_.translate = $1", translate)
}

func (*CSSStyleDeclaration) UnicodeBidi() (unicodeBidi string) {
	macro.Rewrite("$_.unicodeBidi")
	return unicodeBidi
}

func (*CSSStyleDeclaration) SetUnicodeBidi(unicodeBidi string) {
	macro.Rewrite("$_.unicodeBidi = $1", unicodeBidi)
}

func (*CSSStyleDeclaration) VerticalAlign() (verticalAlign string) {
	macro.Rewrite("$_.verticalAlign")
	return verticalAlign
}

func (*CSSStyleDeclaration) SetVerticalAlign(verticalAlign string) {
	macro.Rewrite("$_.verticalAlign = $1", verticalAlign)
}

func (*CSSStyleDeclaration) Visibility() (visibility string) {
	macro.Rewrite("$_.visibility")
	return visibility
}

func (*CSSStyleDeclaration) SetVisibility(visibility string) {
	macro.Rewrite("$_.visibility = $1", visibility)
}

func (*CSSStyleDeclaration) WebkitAlignContent() (webkitAlignContent string) {
	macro.Rewrite("$_.webkitAlignContent")
	return webkitAlignContent
}

func (*CSSStyleDeclaration) SetWebkitAlignContent(webkitAlignContent string) {
	macro.Rewrite("$_.webkitAlignContent = $1", webkitAlignContent)
}

func (*CSSStyleDeclaration) WebkitAlignItems() (webkitAlignItems string) {
	macro.Rewrite("$_.webkitAlignItems")
	return webkitAlignItems
}

func (*CSSStyleDeclaration) SetWebkitAlignItems(webkitAlignItems string) {
	macro.Rewrite("$_.webkitAlignItems = $1", webkitAlignItems)
}

func (*CSSStyleDeclaration) WebkitAlignSelf() (webkitAlignSelf string) {
	macro.Rewrite("$_.webkitAlignSelf")
	return webkitAlignSelf
}

func (*CSSStyleDeclaration) SetWebkitAlignSelf(webkitAlignSelf string) {
	macro.Rewrite("$_.webkitAlignSelf = $1", webkitAlignSelf)
}

func (*CSSStyleDeclaration) WebkitAnimation() (webkitAnimation string) {
	macro.Rewrite("$_.webkitAnimation")
	return webkitAnimation
}

func (*CSSStyleDeclaration) SetWebkitAnimation(webkitAnimation string) {
	macro.Rewrite("$_.webkitAnimation = $1", webkitAnimation)
}

func (*CSSStyleDeclaration) WebkitAnimationDelay() (webkitAnimationDelay string) {
	macro.Rewrite("$_.webkitAnimationDelay")
	return webkitAnimationDelay
}

func (*CSSStyleDeclaration) SetWebkitAnimationDelay(webkitAnimationDelay string) {
	macro.Rewrite("$_.webkitAnimationDelay = $1", webkitAnimationDelay)
}

func (*CSSStyleDeclaration) WebkitAnimationDirection() (webkitAnimationDirection string) {
	macro.Rewrite("$_.webkitAnimationDirection")
	return webkitAnimationDirection
}

func (*CSSStyleDeclaration) SetWebkitAnimationDirection(webkitAnimationDirection string) {
	macro.Rewrite("$_.webkitAnimationDirection = $1", webkitAnimationDirection)
}

func (*CSSStyleDeclaration) WebkitAnimationDuration() (webkitAnimationDuration string) {
	macro.Rewrite("$_.webkitAnimationDuration")
	return webkitAnimationDuration
}

func (*CSSStyleDeclaration) SetWebkitAnimationDuration(webkitAnimationDuration string) {
	macro.Rewrite("$_.webkitAnimationDuration = $1", webkitAnimationDuration)
}

func (*CSSStyleDeclaration) WebkitAnimationFillMode() (webkitAnimationFillMode string) {
	macro.Rewrite("$_.webkitAnimationFillMode")
	return webkitAnimationFillMode
}

func (*CSSStyleDeclaration) SetWebkitAnimationFillMode(webkitAnimationFillMode string) {
	macro.Rewrite("$_.webkitAnimationFillMode = $1", webkitAnimationFillMode)
}

func (*CSSStyleDeclaration) WebkitAnimationIterationCount() (webkitAnimationIterationCount string) {
	macro.Rewrite("$_.webkitAnimationIterationCount")
	return webkitAnimationIterationCount
}

func (*CSSStyleDeclaration) SetWebkitAnimationIterationCount(webkitAnimationIterationCount string) {
	macro.Rewrite("$_.webkitAnimationIterationCount = $1", webkitAnimationIterationCount)
}

func (*CSSStyleDeclaration) WebkitAnimationName() (webkitAnimationName string) {
	macro.Rewrite("$_.webkitAnimationName")
	return webkitAnimationName
}

func (*CSSStyleDeclaration) SetWebkitAnimationName(webkitAnimationName string) {
	macro.Rewrite("$_.webkitAnimationName = $1", webkitAnimationName)
}

func (*CSSStyleDeclaration) WebkitAnimationPlayState() (webkitAnimationPlayState string) {
	macro.Rewrite("$_.webkitAnimationPlayState")
	return webkitAnimationPlayState
}

func (*CSSStyleDeclaration) SetWebkitAnimationPlayState(webkitAnimationPlayState string) {
	macro.Rewrite("$_.webkitAnimationPlayState = $1", webkitAnimationPlayState)
}

func (*CSSStyleDeclaration) WebkitAnimationTimingFunction() (webkitAnimationTimingFunction string) {
	macro.Rewrite("$_.webkitAnimationTimingFunction")
	return webkitAnimationTimingFunction
}

func (*CSSStyleDeclaration) SetWebkitAnimationTimingFunction(webkitAnimationTimingFunction string) {
	macro.Rewrite("$_.webkitAnimationTimingFunction = $1", webkitAnimationTimingFunction)
}

func (*CSSStyleDeclaration) WebkitAppearance() (webkitAppearance string) {
	macro.Rewrite("$_.webkitAppearance")
	return webkitAppearance
}

func (*CSSStyleDeclaration) SetWebkitAppearance(webkitAppearance string) {
	macro.Rewrite("$_.webkitAppearance = $1", webkitAppearance)
}

func (*CSSStyleDeclaration) WebkitBackfaceVisibility() (webkitBackfaceVisibility string) {
	macro.Rewrite("$_.webkitBackfaceVisibility")
	return webkitBackfaceVisibility
}

func (*CSSStyleDeclaration) SetWebkitBackfaceVisibility(webkitBackfaceVisibility string) {
	macro.Rewrite("$_.webkitBackfaceVisibility = $1", webkitBackfaceVisibility)
}

func (*CSSStyleDeclaration) WebkitBackgroundClip() (webkitBackgroundClip string) {
	macro.Rewrite("$_.webkitBackgroundClip")
	return webkitBackgroundClip
}

func (*CSSStyleDeclaration) SetWebkitBackgroundClip(webkitBackgroundClip string) {
	macro.Rewrite("$_.webkitBackgroundClip = $1", webkitBackgroundClip)
}

func (*CSSStyleDeclaration) WebkitBackgroundOrigin() (webkitBackgroundOrigin string) {
	macro.Rewrite("$_.webkitBackgroundOrigin")
	return webkitBackgroundOrigin
}

func (*CSSStyleDeclaration) SetWebkitBackgroundOrigin(webkitBackgroundOrigin string) {
	macro.Rewrite("$_.webkitBackgroundOrigin = $1", webkitBackgroundOrigin)
}

func (*CSSStyleDeclaration) WebkitBackgroundSize() (webkitBackgroundSize string) {
	macro.Rewrite("$_.webkitBackgroundSize")
	return webkitBackgroundSize
}

func (*CSSStyleDeclaration) SetWebkitBackgroundSize(webkitBackgroundSize string) {
	macro.Rewrite("$_.webkitBackgroundSize = $1", webkitBackgroundSize)
}

func (*CSSStyleDeclaration) WebkitBorderBottomLeftRadius() (webkitBorderBottomLeftRadius string) {
	macro.Rewrite("$_.webkitBorderBottomLeftRadius")
	return webkitBorderBottomLeftRadius
}

func (*CSSStyleDeclaration) SetWebkitBorderBottomLeftRadius(webkitBorderBottomLeftRadius string) {
	macro.Rewrite("$_.webkitBorderBottomLeftRadius = $1", webkitBorderBottomLeftRadius)
}

func (*CSSStyleDeclaration) WebkitBorderBottomRightRadius() (webkitBorderBottomRightRadius string) {
	macro.Rewrite("$_.webkitBorderBottomRightRadius")
	return webkitBorderBottomRightRadius
}

func (*CSSStyleDeclaration) SetWebkitBorderBottomRightRadius(webkitBorderBottomRightRadius string) {
	macro.Rewrite("$_.webkitBorderBottomRightRadius = $1", webkitBorderBottomRightRadius)
}

func (*CSSStyleDeclaration) WebkitBorderImage() (webkitBorderImage string) {
	macro.Rewrite("$_.webkitBorderImage")
	return webkitBorderImage
}

func (*CSSStyleDeclaration) SetWebkitBorderImage(webkitBorderImage string) {
	macro.Rewrite("$_.webkitBorderImage = $1", webkitBorderImage)
}

func (*CSSStyleDeclaration) WebkitBorderRadius() (webkitBorderRadius string) {
	macro.Rewrite("$_.webkitBorderRadius")
	return webkitBorderRadius
}

func (*CSSStyleDeclaration) SetWebkitBorderRadius(webkitBorderRadius string) {
	macro.Rewrite("$_.webkitBorderRadius = $1", webkitBorderRadius)
}

func (*CSSStyleDeclaration) WebkitBorderTopLeftRadius() (webkitBorderTopLeftRadius string) {
	macro.Rewrite("$_.webkitBorderTopLeftRadius")
	return webkitBorderTopLeftRadius
}

func (*CSSStyleDeclaration) SetWebkitBorderTopLeftRadius(webkitBorderTopLeftRadius string) {
	macro.Rewrite("$_.webkitBorderTopLeftRadius = $1", webkitBorderTopLeftRadius)
}

func (*CSSStyleDeclaration) WebkitBorderTopRightRadius() (webkitBorderTopRightRadius string) {
	macro.Rewrite("$_.webkitBorderTopRightRadius")
	return webkitBorderTopRightRadius
}

func (*CSSStyleDeclaration) SetWebkitBorderTopRightRadius(webkitBorderTopRightRadius string) {
	macro.Rewrite("$_.webkitBorderTopRightRadius = $1", webkitBorderTopRightRadius)
}

func (*CSSStyleDeclaration) WebkitBoxAlign() (webkitBoxAlign string) {
	macro.Rewrite("$_.webkitBoxAlign")
	return webkitBoxAlign
}

func (*CSSStyleDeclaration) SetWebkitBoxAlign(webkitBoxAlign string) {
	macro.Rewrite("$_.webkitBoxAlign = $1", webkitBoxAlign)
}

func (*CSSStyleDeclaration) WebkitBoxDirection() (webkitBoxDirection string) {
	macro.Rewrite("$_.webkitBoxDirection")
	return webkitBoxDirection
}

func (*CSSStyleDeclaration) SetWebkitBoxDirection(webkitBoxDirection string) {
	macro.Rewrite("$_.webkitBoxDirection = $1", webkitBoxDirection)
}

func (*CSSStyleDeclaration) WebkitBoxFlex() (webkitBoxFlex string) {
	macro.Rewrite("$_.webkitBoxFlex")
	return webkitBoxFlex
}

func (*CSSStyleDeclaration) SetWebkitBoxFlex(webkitBoxFlex string) {
	macro.Rewrite("$_.webkitBoxFlex = $1", webkitBoxFlex)
}

func (*CSSStyleDeclaration) WebkitBoxOrdinalGroup() (webkitBoxOrdinalGroup string) {
	macro.Rewrite("$_.webkitBoxOrdinalGroup")
	return webkitBoxOrdinalGroup
}

func (*CSSStyleDeclaration) SetWebkitBoxOrdinalGroup(webkitBoxOrdinalGroup string) {
	macro.Rewrite("$_.webkitBoxOrdinalGroup = $1", webkitBoxOrdinalGroup)
}

func (*CSSStyleDeclaration) WebkitBoxOrient() (webkitBoxOrient string) {
	macro.Rewrite("$_.webkitBoxOrient")
	return webkitBoxOrient
}

func (*CSSStyleDeclaration) SetWebkitBoxOrient(webkitBoxOrient string) {
	macro.Rewrite("$_.webkitBoxOrient = $1", webkitBoxOrient)
}

func (*CSSStyleDeclaration) WebkitBoxPack() (webkitBoxPack string) {
	macro.Rewrite("$_.webkitBoxPack")
	return webkitBoxPack
}

func (*CSSStyleDeclaration) SetWebkitBoxPack(webkitBoxPack string) {
	macro.Rewrite("$_.webkitBoxPack = $1", webkitBoxPack)
}

func (*CSSStyleDeclaration) WebkitBoxSizing() (webkitBoxSizing string) {
	macro.Rewrite("$_.webkitBoxSizing")
	return webkitBoxSizing
}

func (*CSSStyleDeclaration) SetWebkitBoxSizing(webkitBoxSizing string) {
	macro.Rewrite("$_.webkitBoxSizing = $1", webkitBoxSizing)
}

func (*CSSStyleDeclaration) WebkitColumnBreakAfter() (webkitColumnBreakAfter string) {
	macro.Rewrite("$_.webkitColumnBreakAfter")
	return webkitColumnBreakAfter
}

func (*CSSStyleDeclaration) SetWebkitColumnBreakAfter(webkitColumnBreakAfter string) {
	macro.Rewrite("$_.webkitColumnBreakAfter = $1", webkitColumnBreakAfter)
}

func (*CSSStyleDeclaration) WebkitColumnBreakBefore() (webkitColumnBreakBefore string) {
	macro.Rewrite("$_.webkitColumnBreakBefore")
	return webkitColumnBreakBefore
}

func (*CSSStyleDeclaration) SetWebkitColumnBreakBefore(webkitColumnBreakBefore string) {
	macro.Rewrite("$_.webkitColumnBreakBefore = $1", webkitColumnBreakBefore)
}

func (*CSSStyleDeclaration) WebkitColumnBreakInside() (webkitColumnBreakInside string) {
	macro.Rewrite("$_.webkitColumnBreakInside")
	return webkitColumnBreakInside
}

func (*CSSStyleDeclaration) SetWebkitColumnBreakInside(webkitColumnBreakInside string) {
	macro.Rewrite("$_.webkitColumnBreakInside = $1", webkitColumnBreakInside)
}

func (*CSSStyleDeclaration) WebkitColumnCount() (webkitColumnCount interface{}) {
	macro.Rewrite("$_.webkitColumnCount")
	return webkitColumnCount
}

func (*CSSStyleDeclaration) SetWebkitColumnCount(webkitColumnCount interface{}) {
	macro.Rewrite("$_.webkitColumnCount = $1", webkitColumnCount)
}

func (*CSSStyleDeclaration) WebkitColumnGap() (webkitColumnGap interface{}) {
	macro.Rewrite("$_.webkitColumnGap")
	return webkitColumnGap
}

func (*CSSStyleDeclaration) SetWebkitColumnGap(webkitColumnGap interface{}) {
	macro.Rewrite("$_.webkitColumnGap = $1", webkitColumnGap)
}

func (*CSSStyleDeclaration) WebkitColumnRule() (webkitColumnRule string) {
	macro.Rewrite("$_.webkitColumnRule")
	return webkitColumnRule
}

func (*CSSStyleDeclaration) SetWebkitColumnRule(webkitColumnRule string) {
	macro.Rewrite("$_.webkitColumnRule = $1", webkitColumnRule)
}

func (*CSSStyleDeclaration) WebkitColumnRuleColor() (webkitColumnRuleColor interface{}) {
	macro.Rewrite("$_.webkitColumnRuleColor")
	return webkitColumnRuleColor
}

func (*CSSStyleDeclaration) SetWebkitColumnRuleColor(webkitColumnRuleColor interface{}) {
	macro.Rewrite("$_.webkitColumnRuleColor = $1", webkitColumnRuleColor)
}

func (*CSSStyleDeclaration) WebkitColumnRuleStyle() (webkitColumnRuleStyle string) {
	macro.Rewrite("$_.webkitColumnRuleStyle")
	return webkitColumnRuleStyle
}

func (*CSSStyleDeclaration) SetWebkitColumnRuleStyle(webkitColumnRuleStyle string) {
	macro.Rewrite("$_.webkitColumnRuleStyle = $1", webkitColumnRuleStyle)
}

func (*CSSStyleDeclaration) WebkitColumnRuleWidth() (webkitColumnRuleWidth interface{}) {
	macro.Rewrite("$_.webkitColumnRuleWidth")
	return webkitColumnRuleWidth
}

func (*CSSStyleDeclaration) SetWebkitColumnRuleWidth(webkitColumnRuleWidth interface{}) {
	macro.Rewrite("$_.webkitColumnRuleWidth = $1", webkitColumnRuleWidth)
}

func (*CSSStyleDeclaration) WebkitColumns() (webkitColumns string) {
	macro.Rewrite("$_.webkitColumns")
	return webkitColumns
}

func (*CSSStyleDeclaration) SetWebkitColumns(webkitColumns string) {
	macro.Rewrite("$_.webkitColumns = $1", webkitColumns)
}

func (*CSSStyleDeclaration) WebkitColumnSpan() (webkitColumnSpan string) {
	macro.Rewrite("$_.webkitColumnSpan")
	return webkitColumnSpan
}

func (*CSSStyleDeclaration) SetWebkitColumnSpan(webkitColumnSpan string) {
	macro.Rewrite("$_.webkitColumnSpan = $1", webkitColumnSpan)
}

func (*CSSStyleDeclaration) WebkitColumnWidth() (webkitColumnWidth interface{}) {
	macro.Rewrite("$_.webkitColumnWidth")
	return webkitColumnWidth
}

func (*CSSStyleDeclaration) SetWebkitColumnWidth(webkitColumnWidth interface{}) {
	macro.Rewrite("$_.webkitColumnWidth = $1", webkitColumnWidth)
}

func (*CSSStyleDeclaration) WebkitFilter() (webkitFilter string) {
	macro.Rewrite("$_.webkitFilter")
	return webkitFilter
}

func (*CSSStyleDeclaration) SetWebkitFilter(webkitFilter string) {
	macro.Rewrite("$_.webkitFilter = $1", webkitFilter)
}

func (*CSSStyleDeclaration) WebkitFlex() (webkitFlex string) {
	macro.Rewrite("$_.webkitFlex")
	return webkitFlex
}

func (*CSSStyleDeclaration) SetWebkitFlex(webkitFlex string) {
	macro.Rewrite("$_.webkitFlex = $1", webkitFlex)
}

func (*CSSStyleDeclaration) WebkitFlexBasis() (webkitFlexBasis string) {
	macro.Rewrite("$_.webkitFlexBasis")
	return webkitFlexBasis
}

func (*CSSStyleDeclaration) SetWebkitFlexBasis(webkitFlexBasis string) {
	macro.Rewrite("$_.webkitFlexBasis = $1", webkitFlexBasis)
}

func (*CSSStyleDeclaration) WebkitFlexDirection() (webkitFlexDirection string) {
	macro.Rewrite("$_.webkitFlexDirection")
	return webkitFlexDirection
}

func (*CSSStyleDeclaration) SetWebkitFlexDirection(webkitFlexDirection string) {
	macro.Rewrite("$_.webkitFlexDirection = $1", webkitFlexDirection)
}

func (*CSSStyleDeclaration) WebkitFlexFlow() (webkitFlexFlow string) {
	macro.Rewrite("$_.webkitFlexFlow")
	return webkitFlexFlow
}

func (*CSSStyleDeclaration) SetWebkitFlexFlow(webkitFlexFlow string) {
	macro.Rewrite("$_.webkitFlexFlow = $1", webkitFlexFlow)
}

func (*CSSStyleDeclaration) WebkitFlexGrow() (webkitFlexGrow string) {
	macro.Rewrite("$_.webkitFlexGrow")
	return webkitFlexGrow
}

func (*CSSStyleDeclaration) SetWebkitFlexGrow(webkitFlexGrow string) {
	macro.Rewrite("$_.webkitFlexGrow = $1", webkitFlexGrow)
}

func (*CSSStyleDeclaration) WebkitFlexShrink() (webkitFlexShrink string) {
	macro.Rewrite("$_.webkitFlexShrink")
	return webkitFlexShrink
}

func (*CSSStyleDeclaration) SetWebkitFlexShrink(webkitFlexShrink string) {
	macro.Rewrite("$_.webkitFlexShrink = $1", webkitFlexShrink)
}

func (*CSSStyleDeclaration) WebkitFlexWrap() (webkitFlexWrap string) {
	macro.Rewrite("$_.webkitFlexWrap")
	return webkitFlexWrap
}

func (*CSSStyleDeclaration) SetWebkitFlexWrap(webkitFlexWrap string) {
	macro.Rewrite("$_.webkitFlexWrap = $1", webkitFlexWrap)
}

func (*CSSStyleDeclaration) WebkitJustifyContent() (webkitJustifyContent string) {
	macro.Rewrite("$_.webkitJustifyContent")
	return webkitJustifyContent
}

func (*CSSStyleDeclaration) SetWebkitJustifyContent(webkitJustifyContent string) {
	macro.Rewrite("$_.webkitJustifyContent = $1", webkitJustifyContent)
}

func (*CSSStyleDeclaration) WebkitOrder() (webkitOrder string) {
	macro.Rewrite("$_.webkitOrder")
	return webkitOrder
}

func (*CSSStyleDeclaration) SetWebkitOrder(webkitOrder string) {
	macro.Rewrite("$_.webkitOrder = $1", webkitOrder)
}

func (*CSSStyleDeclaration) WebkitPerspective() (webkitPerspective string) {
	macro.Rewrite("$_.webkitPerspective")
	return webkitPerspective
}

func (*CSSStyleDeclaration) SetWebkitPerspective(webkitPerspective string) {
	macro.Rewrite("$_.webkitPerspective = $1", webkitPerspective)
}

func (*CSSStyleDeclaration) WebkitPerspectiveOrigin() (webkitPerspectiveOrigin string) {
	macro.Rewrite("$_.webkitPerspectiveOrigin")
	return webkitPerspectiveOrigin
}

func (*CSSStyleDeclaration) SetWebkitPerspectiveOrigin(webkitPerspectiveOrigin string) {
	macro.Rewrite("$_.webkitPerspectiveOrigin = $1", webkitPerspectiveOrigin)
}

func (*CSSStyleDeclaration) WebkitTapHighlightColor() (webkitTapHighlightColor string) {
	macro.Rewrite("$_.webkitTapHighlightColor")
	return webkitTapHighlightColor
}

func (*CSSStyleDeclaration) SetWebkitTapHighlightColor(webkitTapHighlightColor string) {
	macro.Rewrite("$_.webkitTapHighlightColor = $1", webkitTapHighlightColor)
}

func (*CSSStyleDeclaration) WebkitTextFillColor() (webkitTextFillColor string) {
	macro.Rewrite("$_.webkitTextFillColor")
	return webkitTextFillColor
}

func (*CSSStyleDeclaration) SetWebkitTextFillColor(webkitTextFillColor string) {
	macro.Rewrite("$_.webkitTextFillColor = $1", webkitTextFillColor)
}

func (*CSSStyleDeclaration) WebkitTextSizeAdjust() (webkitTextSizeAdjust interface{}) {
	macro.Rewrite("$_.webkitTextSizeAdjust")
	return webkitTextSizeAdjust
}

func (*CSSStyleDeclaration) SetWebkitTextSizeAdjust(webkitTextSizeAdjust interface{}) {
	macro.Rewrite("$_.webkitTextSizeAdjust = $1", webkitTextSizeAdjust)
}

func (*CSSStyleDeclaration) WebkitTextStroke() (webkitTextStroke string) {
	macro.Rewrite("$_.webkitTextStroke")
	return webkitTextStroke
}

func (*CSSStyleDeclaration) SetWebkitTextStroke(webkitTextStroke string) {
	macro.Rewrite("$_.webkitTextStroke = $1", webkitTextStroke)
}

func (*CSSStyleDeclaration) WebkitTextStrokeColor() (webkitTextStrokeColor string) {
	macro.Rewrite("$_.webkitTextStrokeColor")
	return webkitTextStrokeColor
}

func (*CSSStyleDeclaration) SetWebkitTextStrokeColor(webkitTextStrokeColor string) {
	macro.Rewrite("$_.webkitTextStrokeColor = $1", webkitTextStrokeColor)
}

func (*CSSStyleDeclaration) WebkitTextStrokeWidth() (webkitTextStrokeWidth string) {
	macro.Rewrite("$_.webkitTextStrokeWidth")
	return webkitTextStrokeWidth
}

func (*CSSStyleDeclaration) SetWebkitTextStrokeWidth(webkitTextStrokeWidth string) {
	macro.Rewrite("$_.webkitTextStrokeWidth = $1", webkitTextStrokeWidth)
}

func (*CSSStyleDeclaration) WebkitTransform() (webkitTransform string) {
	macro.Rewrite("$_.webkitTransform")
	return webkitTransform
}

func (*CSSStyleDeclaration) SetWebkitTransform(webkitTransform string) {
	macro.Rewrite("$_.webkitTransform = $1", webkitTransform)
}

func (*CSSStyleDeclaration) WebkitTransformOrigin() (webkitTransformOrigin string) {
	macro.Rewrite("$_.webkitTransformOrigin")
	return webkitTransformOrigin
}

func (*CSSStyleDeclaration) SetWebkitTransformOrigin(webkitTransformOrigin string) {
	macro.Rewrite("$_.webkitTransformOrigin = $1", webkitTransformOrigin)
}

func (*CSSStyleDeclaration) WebkitTransformStyle() (webkitTransformStyle string) {
	macro.Rewrite("$_.webkitTransformStyle")
	return webkitTransformStyle
}

func (*CSSStyleDeclaration) SetWebkitTransformStyle(webkitTransformStyle string) {
	macro.Rewrite("$_.webkitTransformStyle = $1", webkitTransformStyle)
}

func (*CSSStyleDeclaration) WebkitTransition() (webkitTransition string) {
	macro.Rewrite("$_.webkitTransition")
	return webkitTransition
}

func (*CSSStyleDeclaration) SetWebkitTransition(webkitTransition string) {
	macro.Rewrite("$_.webkitTransition = $1", webkitTransition)
}

func (*CSSStyleDeclaration) WebkitTransitionDelay() (webkitTransitionDelay string) {
	macro.Rewrite("$_.webkitTransitionDelay")
	return webkitTransitionDelay
}

func (*CSSStyleDeclaration) SetWebkitTransitionDelay(webkitTransitionDelay string) {
	macro.Rewrite("$_.webkitTransitionDelay = $1", webkitTransitionDelay)
}

func (*CSSStyleDeclaration) WebkitTransitionDuration() (webkitTransitionDuration string) {
	macro.Rewrite("$_.webkitTransitionDuration")
	return webkitTransitionDuration
}

func (*CSSStyleDeclaration) SetWebkitTransitionDuration(webkitTransitionDuration string) {
	macro.Rewrite("$_.webkitTransitionDuration = $1", webkitTransitionDuration)
}

func (*CSSStyleDeclaration) WebkitTransitionProperty() (webkitTransitionProperty string) {
	macro.Rewrite("$_.webkitTransitionProperty")
	return webkitTransitionProperty
}

func (*CSSStyleDeclaration) SetWebkitTransitionProperty(webkitTransitionProperty string) {
	macro.Rewrite("$_.webkitTransitionProperty = $1", webkitTransitionProperty)
}

func (*CSSStyleDeclaration) WebkitTransitionTimingFunction() (webkitTransitionTimingFunction string) {
	macro.Rewrite("$_.webkitTransitionTimingFunction")
	return webkitTransitionTimingFunction
}

func (*CSSStyleDeclaration) SetWebkitTransitionTimingFunction(webkitTransitionTimingFunction string) {
	macro.Rewrite("$_.webkitTransitionTimingFunction = $1", webkitTransitionTimingFunction)
}

func (*CSSStyleDeclaration) WebkitUserModify() (webkitUserModify string) {
	macro.Rewrite("$_.webkitUserModify")
	return webkitUserModify
}

func (*CSSStyleDeclaration) SetWebkitUserModify(webkitUserModify string) {
	macro.Rewrite("$_.webkitUserModify = $1", webkitUserModify)
}

func (*CSSStyleDeclaration) WebkitUserSelect() (webkitUserSelect string) {
	macro.Rewrite("$_.webkitUserSelect")
	return webkitUserSelect
}

func (*CSSStyleDeclaration) SetWebkitUserSelect(webkitUserSelect string) {
	macro.Rewrite("$_.webkitUserSelect = $1", webkitUserSelect)
}

func (*CSSStyleDeclaration) WebkitWritingMode() (webkitWritingMode string) {
	macro.Rewrite("$_.webkitWritingMode")
	return webkitWritingMode
}

func (*CSSStyleDeclaration) SetWebkitWritingMode(webkitWritingMode string) {
	macro.Rewrite("$_.webkitWritingMode = $1", webkitWritingMode)
}

func (*CSSStyleDeclaration) WhiteSpace() (whiteSpace string) {
	macro.Rewrite("$_.whiteSpace")
	return whiteSpace
}

func (*CSSStyleDeclaration) SetWhiteSpace(whiteSpace string) {
	macro.Rewrite("$_.whiteSpace = $1", whiteSpace)
}

func (*CSSStyleDeclaration) Widows() (widows string) {
	macro.Rewrite("$_.widows")
	return widows
}

func (*CSSStyleDeclaration) SetWidows(widows string) {
	macro.Rewrite("$_.widows = $1", widows)
}

func (*CSSStyleDeclaration) Width() (width string) {
	macro.Rewrite("$_.width")
	return width
}

func (*CSSStyleDeclaration) SetWidth(width string) {
	macro.Rewrite("$_.width = $1", width)
}

func (*CSSStyleDeclaration) WordBreak() (wordBreak string) {
	macro.Rewrite("$_.wordBreak")
	return wordBreak
}

func (*CSSStyleDeclaration) SetWordBreak(wordBreak string) {
	macro.Rewrite("$_.wordBreak = $1", wordBreak)
}

func (*CSSStyleDeclaration) WordSpacing() (wordSpacing string) {
	macro.Rewrite("$_.wordSpacing")
	return wordSpacing
}

func (*CSSStyleDeclaration) SetWordSpacing(wordSpacing string) {
	macro.Rewrite("$_.wordSpacing = $1", wordSpacing)
}

func (*CSSStyleDeclaration) WordWrap() (wordWrap string) {
	macro.Rewrite("$_.wordWrap")
	return wordWrap
}

func (*CSSStyleDeclaration) SetWordWrap(wordWrap string) {
	macro.Rewrite("$_.wordWrap = $1", wordWrap)
}

func (*CSSStyleDeclaration) WritingMode() (writingMode string) {
	macro.Rewrite("$_.writingMode")
	return writingMode
}

func (*CSSStyleDeclaration) SetWritingMode(writingMode string) {
	macro.Rewrite("$_.writingMode = $1", writingMode)
}

func (*CSSStyleDeclaration) ZIndex() (zIndex string) {
	macro.Rewrite("$_.zIndex")
	return zIndex
}

func (*CSSStyleDeclaration) SetZIndex(zIndex string) {
	macro.Rewrite("$_.zIndex = $1", zIndex)
}

func (*CSSStyleDeclaration) Zoom() (zoom string) {
	macro.Rewrite("$_.zoom")
	return zoom
}

func (*CSSStyleDeclaration) SetZoom(zoom string) {
	macro.Rewrite("$_.zoom = $1", zoom)
}
