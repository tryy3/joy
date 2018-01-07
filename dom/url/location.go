package url

import "github.com/matthewmueller/joy/macro"

// Location struct
// js:"Location,omit"
type Location struct {
}

// Assign fn
// js:"assign"
func (*Location) Assign(url string) {
	macro.Rewrite("$_.assign($1)", url)
}

// Reload fn
// js:"reload"
func (*Location) Reload(forcedReload *bool) {
	macro.Rewrite("$_.reload($1)", forcedReload)
}

// Replace fn
// js:"replace"
func (*Location) Replace(url string) {
	macro.Rewrite("$_.replace($1)", url)
}

// ToString fn
// js:"toString"
func (*Location) ToString() (s string) {
	macro.Rewrite("$_.toString()")
	return s
}

// Hash prop
// js:"hash"
func (*Location) Hash() (hash string) {
	macro.Rewrite("$_.hash")
	return hash
}

// SetHash prop
// js:"hash"
func (*Location) SetHash(hash string) {
	macro.Rewrite("$_.hash = $1", hash)
}

// Host prop
// js:"host"
func (*Location) Host() (host string) {
	macro.Rewrite("$_.host")
	return host
}

// SetHost prop
// js:"host"
func (*Location) SetHost(host string) {
	macro.Rewrite("$_.host = $1", host)
}

// Hostname prop
// js:"hostname"
func (*Location) Hostname() (hostname string) {
	macro.Rewrite("$_.hostname")
	return hostname
}

// SetHostname prop
// js:"hostname"
func (*Location) SetHostname(hostname string) {
	macro.Rewrite("$_.hostname = $1", hostname)
}

// Href prop
// js:"href"
func (*Location) Href() (href string) {
	macro.Rewrite("$_.href")
	return href
}

// SetHref prop
// js:"href"
func (*Location) SetHref(href string) {
	macro.Rewrite("$_.href = $1", href)
}

// Origin prop
// js:"origin"
func (*Location) Origin() (origin string) {
	macro.Rewrite("$_.origin")
	return origin
}

// Pathname prop
// js:"pathname"
func (*Location) Pathname() (pathname string) {
	macro.Rewrite("$_.pathname")
	return pathname
}

// SetPathname prop
// js:"pathname"
func (*Location) SetPathname(pathname string) {
	macro.Rewrite("$_.pathname = $1", pathname)
}

// Port prop
// js:"port"
func (*Location) Port() (port string) {
	macro.Rewrite("$_.port")
	return port
}

// SetPort prop
// js:"port"
func (*Location) SetPort(port string) {
	macro.Rewrite("$_.port = $1", port)
}

// Protocol prop
// js:"protocol"
func (*Location) Protocol() (protocol string) {
	macro.Rewrite("$_.protocol")
	return protocol
}

// SetProtocol prop
// js:"protocol"
func (*Location) SetProtocol(protocol string) {
	macro.Rewrite("$_.protocol = $1", protocol)
}

// Search prop
// js:"search"
func (*Location) Search() (search string) {
	macro.Rewrite("$_.search")
	return search
}

// SetSearch prop
// js:"search"
func (*Location) SetSearch(search string) {
	macro.Rewrite("$_.search = $1", search)
}
