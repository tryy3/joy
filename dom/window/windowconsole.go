package window

import "github.com/matthewmueller/joy/dom/utils"

type WindowConsole interface {
	Console() (console *utils.Console)
}
