package ms

import "github.com/matthewmueller/joy/macro"

type ExtensionScriptApis struct {
}

func (*ExtensionScriptApis) ExtensionIDToShortID(extensionId string) (i int) {
	macro.Rewrite("$_.extensionIdToShortId($1)", extensionId)
	return i
}

func (*ExtensionScriptApis) FireExtensionAPITelemetry(functionName string, isSucceeded bool, isSupported bool) {
	macro.Rewrite("$_.fireExtensionApiTelemetry($1, $2, $3)", functionName, isSucceeded, isSupported)
}

func (*ExtensionScriptApis) GenericFunction(routerAddress interface{}, parameters *string, callbackId *int) {
	macro.Rewrite("$_.genericFunction($1, $2, $3)", routerAddress, parameters, callbackId)
}

func (*ExtensionScriptApis) GenericSynchronousFunction(functionId int, parameters *string) (s string) {
	macro.Rewrite("$_.genericSynchronousFunction($1, $2)", functionId, parameters)
	return s
}

func (*ExtensionScriptApis) GetExtensionID() (s string) {
	macro.Rewrite("$_.getExtensionId()")
	return s
}

func (*ExtensionScriptApis) RegisterGenericFunctionCallbackHandler(callbackHandler func()) {
	macro.Rewrite("$_.registerGenericFunctionCallbackHandler($1)", callbackHandler)
}

func (*ExtensionScriptApis) RegisterGenericPersistentCallbackHandler(callbackHandler func()) {
	macro.Rewrite("$_.registerGenericPersistentCallbackHandler($1)", callbackHandler)
}
