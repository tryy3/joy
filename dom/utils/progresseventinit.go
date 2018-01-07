package utils

import "github.com/matthewmueller/joy/dom/event"

type ProgressEventInit struct {
	*event.EventInit

	lengthComputable *bool
	loaded           *uint64
	total            *uint64
}
