package permissions

import "github.com/matthewmueller/joy/dom/ms"

// DeferredPermissionRequest interface
// js:"DeferredPermissionRequest"
type DeferredPermissionRequest interface {

	// Allow
	// js:"allow"
	// jsrewrite:"$_.allow()"
	Allow()

	// Deny
	// js:"deny"
	// jsrewrite:"$_.deny()"
	Deny()

	// ID prop
	// js:"id"
	// jsrewrite:"$_.id"
	ID() (id uint)

	// Type prop
	// js:"type"
	// jsrewrite:"$_.type"
	Type() (kind *ms.MSWebViewPermissionType)

	// URI prop
	// js:"uri"
	// jsrewrite:"$_.uri"
	URI() (uri string)
}
