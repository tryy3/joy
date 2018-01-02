package utils

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/element"
)

type Console struct {
}

func (*Console) Assert(test *bool, message *string, optionalParams interface{}) {
	macro.Rewrite("$_.assert($1, $2, $3)", test, message, optionalParams)
}

func (*Console) Clear() {
	macro.Rewrite("$_.clear()")
}

func (*Console) Count(countTitle *string) {
	macro.Rewrite("$_.count($1)", countTitle)
}

func (*Console) Debug(message *string, optionalParams interface{}) {
	macro.Rewrite("$_.debug($1, $2)", message, optionalParams)
}

func (*Console) Dir(value *interface{}, optionalParams interface{}) {
	macro.Rewrite("$_.dir($1, $2)", value, optionalParams)
}

func (*Console) Dirxml(value interface{}) {
	macro.Rewrite("$_.dirxml($1)", value)
}

func (*Console) Error(message *string, optionalParams interface{}) {
	macro.Rewrite("$_.error($1, $2)", message, optionalParams)
}

func (*Console) Exception(message *string, optionalParams interface{}) {
	macro.Rewrite("$_.exception($1, $2)", message, optionalParams)
}

func (*Console) Group(groupTitle *string) {
	macro.Rewrite("$_.group($1)", groupTitle)
}

func (*Console) GroupCollapsed(groupTitle *string) {
	macro.Rewrite("$_.groupCollapsed($1)", groupTitle)
}

func (*Console) GroupEnd() {
	macro.Rewrite("$_.groupEnd()")
}

func (*Console) Info(message *string, optionalParams interface{}) {
	macro.Rewrite("$_.info($1, $2)", message, optionalParams)
}

func (*Console) Log(message *string, optionalParams interface{}) {
	macro.Rewrite("$_.log($1, $2)", message, optionalParams)
}

func (*Console) MsIsIndependentlyComposed(element element.Element) (b bool) {
	macro.Rewrite("$_.msIsIndependentlyComposed($1)", element)
	return b
}

func (*Console) Profile(reportName *string) {
	macro.Rewrite("$_.profile($1)", reportName)
}

func (*Console) ProfileEnd() {
	macro.Rewrite("$_.profileEnd()")
}

func (*Console) Select(element element.Element) {
	macro.Rewrite("$_.select($1)", element)
}

func (*Console) Table(data interface{}) {
	macro.Rewrite("$_.table($1)", data)
}

func (*Console) Time(timerName *string) {
	macro.Rewrite("$_.time($1)", timerName)
}

func (*Console) TimeEnd(timerName *string) {
	macro.Rewrite("$_.timeEnd($1)", timerName)
}

func (*Console) Trace() {
	macro.Rewrite("$_.trace()")
}

func (*Console) Warn(message *string, optionalParams interface{}) {
	macro.Rewrite("$_.warn($1, $2)", message, optionalParams)
}
