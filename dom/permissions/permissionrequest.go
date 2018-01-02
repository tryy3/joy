package permissions

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/ms"
)

var _ DeferredPermissionRequest = (*PermissionRequest)(nil)

type PermissionRequest struct {
}

func (*PermissionRequest) Defer() {
	macro.Rewrite("$_.defer()")
}

func (*PermissionRequest) Allow() {
	macro.Rewrite("$_.allow()")
}

func (*PermissionRequest) Deny() {
	macro.Rewrite("$_.deny()")
}

func (*PermissionRequest) State() (state *ms.MSWebViewPermissionState) {
	macro.Rewrite("$_.state")
	return state
}

func (*PermissionRequest) ID() (id uint) {
	macro.Rewrite("$_.id")
	return id
}

func (*PermissionRequest) Type() (kind *ms.MSWebViewPermissionType) {
	macro.Rewrite("$_.type")
	return kind
}

func (*PermissionRequest) URI() (uri string) {
	macro.Rewrite("$_.uri")
	return uri
}
