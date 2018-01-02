package fetch

import "github.com/matthewmueller/joy/dom/file"

type Body interface {
	ArrayBuffer() (b []byte)

	Blob() (b file.Blob)

	JSON() (i interface{})

	Text() (s string)

	BodyUsed() (bodyUsed bool)
}
