package dom


import "github.com/matthewmueller/joy/macro"

// SVGAngle struct
// js:"SVGAngle,omit"
type SVGAngle struct {
}

// ConvertToSpecifiedUnits fn
// js:"convertToSpecifiedUnits"
func (*SVGAngle) ConvertToSpecifiedUnits(unitType uint8) {
	macro.Rewrite("$_.convertToSpecifiedUnits($1)", unitType)
}

// NewValueSpecifiedUnits fn
// js:"newValueSpecifiedUnits"
func (*SVGAngle) NewValueSpecifiedUnits(unitType uint8, valueInSpecifiedUnits float32) {
	macro.Rewrite("$_.newValueSpecifiedUnits($1, $2)", unitType, valueInSpecifiedUnits)
}

// UnitType prop
// js:"unitType"
func (*SVGAngle) UnitType() (unitType uint8) {
	macro.Rewrite("$_.unitType")
	return unitType
}

// Value prop
// js:"value"
func (*SVGAngle) Value() (value float32) {
	macro.Rewrite("$_.value")
	return value
}

// SetValue prop
// js:"value"
func (*SVGAngle) SetValue(value float32) {
	macro.Rewrite("$_.value = $1", value)
}

// ValueAsString prop
// js:"valueAsString"
func (*SVGAngle) ValueAsString() (valueAsString string) {
	macro.Rewrite("$_.valueAsString")
	return valueAsString
}

// SetValueAsString prop
// js:"valueAsString"
func (*SVGAngle) SetValueAsString(valueAsString string) {
	macro.Rewrite("$_.valueAsString = $1", valueAsString)
}

// ValueInSpecifiedUnits prop
// js:"valueInSpecifiedUnits"
func (*SVGAngle) ValueInSpecifiedUnits() (valueInSpecifiedUnits float32) {
	macro.Rewrite("$_.valueInSpecifiedUnits")
	return valueInSpecifiedUnits
}

// SetValueInSpecifiedUnits prop
// js:"valueInSpecifiedUnits"
func (*SVGAngle) SetValueInSpecifiedUnits(valueInSpecifiedUnits float32) {
	macro.Rewrite("$_.valueInSpecifiedUnits = $1", valueInSpecifiedUnits)
}
