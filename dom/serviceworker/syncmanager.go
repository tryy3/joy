package serviceworker

import "github.com/matthewmueller/joy/macro"

type SyncManager struct {
}

func (*SyncManager) GetTags() (s []string) {
	macro.Rewrite("await $_.getTags()")
	return s
}

func (*SyncManager) Register(tag string) {
	macro.Rewrite("await $_.register($1)", tag)
}
