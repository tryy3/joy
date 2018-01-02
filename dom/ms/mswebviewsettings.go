package ms

import "github.com/matthewmueller/joy/macro"

type MSWebViewSettings struct {
}

func (*MSWebViewSettings) IsIndexedDBEnabled() (isIndexedDBEnabled bool) {
	macro.Rewrite("$_.isIndexedDBEnabled")
	return isIndexedDBEnabled
}

func (*MSWebViewSettings) SetIsIndexedDBEnabled(isIndexedDBEnabled bool) {
	macro.Rewrite("$_.isIndexedDBEnabled = $1", isIndexedDBEnabled)
}

func (*MSWebViewSettings) IsJavaScriptEnabled() (isJavaScriptEnabled bool) {
	macro.Rewrite("$_.isJavaScriptEnabled")
	return isJavaScriptEnabled
}

func (*MSWebViewSettings) SetIsJavaScriptEnabled(isJavaScriptEnabled bool) {
	macro.Rewrite("$_.isJavaScriptEnabled = $1", isJavaScriptEnabled)
}
