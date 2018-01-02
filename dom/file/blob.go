package file

type Blob interface {
	MsClose()

	MsDetachStream() (i interface{})

	Slice(start *int64, end *int64, contentType *string) (b Blob)

	Size() (size uint64)

	Type() (kind string)
}
