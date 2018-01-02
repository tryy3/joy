package ms

import "github.com/matthewmueller/joy/dom/event"

type MSBaseReader interface {
	Abort()

	Onabort() (onabort func(event.Event))

	SetOnabort(onabort func(event.Event))

	Onerror() (onerror func(event.Event))

	SetOnerror(onerror func(event.Event))

	Onload() (onload func(event.Event))

	SetOnload(onload func(event.Event))

	Onloadend() (onloadend func(event.Event))

	SetOnloadend(onloadend func(event.Event))

	Onloadstart() (onloadstart func(event.Event))

	SetOnloadstart(onloadstart func(event.Event))

	Onprogress() (onprogress func(event.Event))

	SetOnprogress(onprogress func(event.Event))

	ReadyState() (readyState uint8)

	Result() (result interface{})
}
