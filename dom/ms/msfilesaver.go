package ms

type MSFileSaver interface {
	MsSaveBlob(blob interface{}, defaultName *string) (b bool)

	MsSaveOrOpenBlob(blob interface{}, defaultName *string) (b bool)
}
