package ms

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/file"
)

type MSApp struct {
}

func (*MSApp) ClearTemporaryWebDataAsync() (m *MSAppAsyncOperation) {
	macro.Rewrite("$_.clearTemporaryWebDataAsync()")
	return m
}

func (*MSApp) CreateBlobFromRandomAccessStream(kind string, seeker interface{}) (b file.Blob) {
	macro.Rewrite("$_.createBlobFromRandomAccessStream($1, $2)", kind, seeker)
	return b
}

func (*MSApp) CreateDataPackage(object interface{}) (i interface{}) {
	macro.Rewrite("$_.createDataPackage($1)", object)
	return i
}

func (*MSApp) CreateDataPackageFromSelection() (i interface{}) {
	macro.Rewrite("$_.createDataPackageFromSelection()")
	return i
}

func (*MSApp) CreateFileFromStorageFile(storageFile interface{}) (f *file.File) {
	macro.Rewrite("$_.createFileFromStorageFile($1)", storageFile)
	return f
}

func (*MSApp) CreateStreamFromInputStream(kind string, inputStream interface{}) (m *MSStream) {
	macro.Rewrite("$_.createStreamFromInputStream($1, $2)", kind, inputStream)
	return m
}

func (*MSApp) ExecAsyncAtPriority(asynchronousCallback func(args interface{}) (MSExecAtPriorityFunctionCallback interface{}), priority string, args interface{}) {
	macro.Rewrite("$_.execAsyncAtPriority($1, $2, $3)", asynchronousCallback, priority, args)
}

func (*MSApp) ExecAtPriority(synchronousCallback func(args interface{}) (MSExecAtPriorityFunctionCallback interface{}), priority string, args interface{}) (i interface{}) {
	macro.Rewrite("$_.execAtPriority($1, $2, $3)", synchronousCallback, priority, args)
	return i
}

func (*MSApp) GetCurrentPriority() (s string) {
	macro.Rewrite("$_.getCurrentPriority()")
	return s
}

func (*MSApp) GetHTMLPrintDocumentSourceAsync(htmlDoc interface{}) (i interface{}) {
	macro.Rewrite("await $_.getHtmlPrintDocumentSourceAsync($1)", htmlDoc)
	return i
}

func (*MSApp) GetViewID(view interface{}) (i interface{}) {
	macro.Rewrite("$_.getViewId($1)", view)
	return i
}

func (*MSApp) IsTaskScheduledAtPriorityOrHigher(priority string) (b bool) {
	macro.Rewrite("$_.isTaskScheduledAtPriorityOrHigher($1)", priority)
	return b
}

func (*MSApp) PageHandlesAllApplicationActivations(enabled bool) {
	macro.Rewrite("$_.pageHandlesAllApplicationActivations($1)", enabled)
}

func (*MSApp) SuppressSubdownloadCredentialPrompts(suppress bool) {
	macro.Rewrite("$_.suppressSubdownloadCredentialPrompts($1)", suppress)
}

func (*MSApp) TerminateApp(exceptionObject interface{}) {
	macro.Rewrite("$_.terminateApp($1)", exceptionObject)
}
