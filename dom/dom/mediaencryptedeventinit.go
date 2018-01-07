package dom

import "github.com/matthewmueller/joy/dom/event"

type MediaEncryptedEventInit struct {
	*event.EventInit

	initData     *[]byte
	initDataType *string
}
