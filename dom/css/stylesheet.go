package css

import "github.com/matthewmueller/joy/dom/dom"

type StyleSheet interface {
	Disabled() (disabled bool)

	SetDisabled(disabled bool)

	Href() (href string)

	Media() (media *MediaList)

	OwnerNode() (ownerNode dom.Node)

	ParentStyleSheet() (parentStyleSheet StyleSheet)

	Title() (title string)

	Type() (kind string)
}
