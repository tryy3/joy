package dom

import "github.com/matthewmueller/joy/macro"

type SVGLength struct {
}

func (*SVGLength) ConvertToSpecifiedUnits(unitType uint8) {
	macro.Rewrite("$_.convertToSpecifiedUnits($1)", unitType)
}

func (*SVGLength) NewValueSpecifiedUnits(unitType uint8, valueInSpecifiedUnits float32) {
	macro.Rewrite("$_.newValueSpecifiedUnits($1, $2)", unitType, valueInSpecifiedUnits)
}

func (*SVGLength) UnitType() (unitType uint8) {
	macro.Rewrite("$_.unitType")
	return unitType
}

func (*SVGLength) Value() (value float32) {
	macro.Rewrite("$_.value")
	return value
}

func (*SVGLength) SetValue(value float32) {
	macro.Rewrite("$_.value = $1", value)
}

func (*SVGLength) ValueAsString() (valueAsString string) {
	macro.Rewrite("$_.valueAsString")
	return valueAsString
}

func (*SVGLength) SetValueAsString(valueAsString string) {
	macro.Rewrite("$_.valueAsString = $1", valueAsString)
}

func (*SVGLength) ValueInSpecifiedUnits() (valueInSpecifiedUnits float32) {
	macro.Rewrite("$_.valueInSpecifiedUnits")
	return valueInSpecifiedUnits
}

func (*SVGLength) SetValueInSpecifiedUnits(valueInSpecifiedUnits float32) {
	macro.Rewrite("$_.valueInSpecifiedUnits = $1", valueInSpecifiedUnits)
}
