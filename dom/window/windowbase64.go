package window

type WindowBase64 interface {
	Atob(encodedString string) (s string)

	Btoa(rawString string) (s string)
}
