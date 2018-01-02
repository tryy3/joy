package url

import (
	"github.com/matthewmueller/joy/macro"
)

func New(url string, base *string) *URL {
	macro.Rewrite("new URL($1, $2)", url, base)
	return &URL{}
}

type URL struct {
}

func (*URL) CreateObjectURL(object interface{}, options *ObjectURLOptions) (s string) {
	macro.Rewrite("$_.createObjectURL($1, $2)", object, options)
	return s
}

func (*URL) RevokeObjectURL(url string) {
	macro.Rewrite("$_.revokeObjectURL($1)", url)
}

func (*URL) ToString() (s string) {
	macro.Rewrite("$_.toString()")
	return s
}

func (*URL) Hash() (hash string) {
	macro.Rewrite("$_.hash")
	return hash
}

func (*URL) SetHash(hash string) {
	macro.Rewrite("$_.hash = $1", hash)
}

func (*URL) Host() (host string) {
	macro.Rewrite("$_.host")
	return host
}

func (*URL) SetHost(host string) {
	macro.Rewrite("$_.host = $1", host)
}

func (*URL) Hostname() (hostname string) {
	macro.Rewrite("$_.hostname")
	return hostname
}

func (*URL) SetHostname(hostname string) {
	macro.Rewrite("$_.hostname = $1", hostname)
}

func (*URL) Href() (href string) {
	macro.Rewrite("$_.href")
	return href
}

func (*URL) SetHref(href string) {
	macro.Rewrite("$_.href = $1", href)
}

func (*URL) Origin() (origin string) {
	macro.Rewrite("$_.origin")
	return origin
}

func (*URL) Password() (password string) {
	macro.Rewrite("$_.password")
	return password
}

func (*URL) SetPassword(password string) {
	macro.Rewrite("$_.password = $1", password)
}

func (*URL) Pathname() (pathname string) {
	macro.Rewrite("$_.pathname")
	return pathname
}

func (*URL) SetPathname(pathname string) {
	macro.Rewrite("$_.pathname = $1", pathname)
}

func (*URL) Port() (port string) {
	macro.Rewrite("$_.port")
	return port
}

func (*URL) SetPort(port string) {
	macro.Rewrite("$_.port = $1", port)
}

func (*URL) Protocol() (protocol string) {
	macro.Rewrite("$_.protocol")
	return protocol
}

func (*URL) SetProtocol(protocol string) {
	macro.Rewrite("$_.protocol = $1", protocol)
}

func (*URL) Search() (search string) {
	macro.Rewrite("$_.search")
	return search
}

func (*URL) SetSearch(search string) {
	macro.Rewrite("$_.search = $1", search)
}

func (*URL) Username() (username string) {
	macro.Rewrite("$_.username")
	return username
}

func (*URL) SetUsername(username string) {
	macro.Rewrite("$_.username = $1", username)
}
