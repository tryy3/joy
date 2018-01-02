package dom

import "github.com/matthewmueller/joy/macro"

type SVGAngle struct {
}

func (*SVGAngle) ConvertToSpecifiedUnits(unitType uint8) {
	macro.Rewrite("$_.convertToSpecifiedUnits($1)", unitType)
}

func (*SVGAngle) NewValueSpecifiedUnits(unitType uint8, valueInSpecifiedUnits float32) {
	macro.Rewrite("$_.newValueSpecifiedUnits($1, $2)", unitType, valueInSpecifiedUnits)
}

func (*SVGAngle) UnitType() (unitType uint8) {
	macro.Rewrite("$_.unitType")
	return unitType
}

func (*SVGAngle) Value() (value float32) {
	macro.Rewrite("$_.value")
	return value
}

func (*SVGAngle) SetValue(value float32) {
	macro.Rewrite("$_.value = $1", value)
}

func (*SVGAngle) ValueAsString() (valueAsString string) {
	macro.Rewrite("$_.valueAsString")
	return valueAsString
}

func (*SVGAngle) SetValueAsString(valueAsString string) {
	macro.Rewrite("$_.valueAsString = $1", valueAsString)
}

func (*SVGAngle) ValueInSpecifiedUnits() (valueInSpecifiedUnits float32) {
	macro.Rewrite("$_.valueInSpecifiedUnits")
	return valueInSpecifiedUnits
}

func (*SVGAngle) SetValueInSpecifiedUnits(valueInSpecifiedUnits float32) {
	macro.Rewrite("$_.valueInSpecifiedUnits = $1", valueInSpecifiedUnits)
}
