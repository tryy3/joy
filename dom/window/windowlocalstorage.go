package window

import "github.com/matthewmueller/joy/dom/storage"

type WindowLocalStorage interface {
	LocalStorage() (localStorage *storage.Storage)
}
