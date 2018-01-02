package url

import "github.com/matthewmueller/joy/macro"

type Location struct {
}

func (*Location) Assign(url string) {
	macro.Rewrite("$_.assign($1)", url)
}

func (*Location) Reload(forcedReload *bool) {
	macro.Rewrite("$_.reload($1)", forcedReload)
}

func (*Location) Replace(url string) {
	macro.Rewrite("$_.replace($1)", url)
}

func (*Location) ToString() (s string) {
	macro.Rewrite("$_.toString()")
	return s
}

func (*Location) Hash() (hash string) {
	macro.Rewrite("$_.hash")
	return hash
}

func (*Location) SetHash(hash string) {
	macro.Rewrite("$_.hash = $1", hash)
}

func (*Location) Host() (host string) {
	macro.Rewrite("$_.host")
	return host
}

func (*Location) SetHost(host string) {
	macro.Rewrite("$_.host = $1", host)
}

func (*Location) Hostname() (hostname string) {
	macro.Rewrite("$_.hostname")
	return hostname
}

func (*Location) SetHostname(hostname string) {
	macro.Rewrite("$_.hostname = $1", hostname)
}

func (*Location) Href() (href string) {
	macro.Rewrite("$_.href")
	return href
}

func (*Location) SetHref(href string) {
	macro.Rewrite("$_.href = $1", href)
}

func (*Location) Origin() (origin string) {
	macro.Rewrite("$_.origin")
	return origin
}

func (*Location) Pathname() (pathname string) {
	macro.Rewrite("$_.pathname")
	return pathname
}

func (*Location) SetPathname(pathname string) {
	macro.Rewrite("$_.pathname = $1", pathname)
}

func (*Location) Port() (port string) {
	macro.Rewrite("$_.port")
	return port
}

func (*Location) SetPort(port string) {
	macro.Rewrite("$_.port = $1", port)
}

func (*Location) Protocol() (protocol string) {
	macro.Rewrite("$_.protocol")
	return protocol
}

func (*Location) SetProtocol(protocol string) {
	macro.Rewrite("$_.protocol = $1", protocol)
}

func (*Location) Search() (search string) {
	macro.Rewrite("$_.search")
	return search
}

func (*Location) SetSearch(search string) {
	macro.Rewrite("$_.search = $1", search)
}
