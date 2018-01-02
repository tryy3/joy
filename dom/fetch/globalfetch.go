package fetch

type GlobalFetch interface {
	Fetch(input *Request, init *RequestInit) (r *Response)
}
