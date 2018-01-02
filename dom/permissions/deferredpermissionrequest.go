package permissions

import "github.com/matthewmueller/joy/dom/ms"

type DeferredPermissionRequest interface {
	Allow()

	Deny()

	ID() (id uint)

	Type() (kind *ms.MSWebViewPermissionType)

	URI() (uri string)
}
