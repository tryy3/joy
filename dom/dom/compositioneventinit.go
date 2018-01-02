package dom

import "github.com/matthewmueller/joy/dom/ui"

type CompositionEventInit struct {
	*ui.UIEventInit

	data	*string
}
