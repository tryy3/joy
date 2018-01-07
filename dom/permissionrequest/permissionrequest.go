package permissionrequest

import (
	"github.com/matthewmueller/joy/dom/ms"
	"github.com/matthewmueller/joy/dom/permissions"
	"github.com/matthewmueller/joy/macro"
)

var _ permissions.DeferredPermissionRequest = (*PermissionRequest)(nil)

// PermissionRequest struct
// js:"PermissionRequest,omit"
type PermissionRequest struct {
}

// Defer fn
// js:"defer"
func (*PermissionRequest) Defer() {
	macro.Rewrite("$_.defer()")
}

// Allow fn
// js:"allow"
func (*PermissionRequest) Allow() {
	macro.Rewrite("$_.allow()")
}

// Deny fn
// js:"deny"
func (*PermissionRequest) Deny() {
	macro.Rewrite("$_.deny()")
}

// State prop
// js:"state"
func (*PermissionRequest) State() (state *ms.MSWebViewPermissionState) {
	macro.Rewrite("$_.state")
	return state
}

// ID prop
// js:"id"
func (*PermissionRequest) ID() (id uint) {
	macro.Rewrite("$_.id")
	return id
}

// Type prop
// js:"type"
func (*PermissionRequest) Type() (kind *ms.MSWebViewPermissionType) {
	macro.Rewrite("$_.type")
	return kind
}

// URI prop
// js:"uri"
func (*PermissionRequest) URI() (uri string) {
	macro.Rewrite("$_.uri")
	return uri
}
