package fetch

type RequestInit struct {
	body		*interface{}
	cache		*RequestCache
	credentials	*RequestCredentials
	headers		*interface{}
	integrity	*string
	keepalive	*bool
	method		*string
	mode		*RequestMode
	redirect	*RequestRedirect
	referrer	*string
	referrerPolicy	*ReferrerPolicy
	window		*interface{}
}
