package window

import "github.com/matthewmueller/joy/dom/storage"

type WindowSessionStorage interface {
	SessionStorage() (sessionStorage *storage.Storage)
}
