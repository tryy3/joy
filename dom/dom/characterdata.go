package dom

type CharacterData interface {
	Node

	AppendData(arg string)

	DeleteData(offset uint, count uint)

	InsertData(offset uint, arg string)

	ReplaceData(offset uint, count uint, arg string)

	SubstringData(offset uint, count uint) (s string)

	Data() (data string)

	SetData(data string)

	Length() (length uint)
}
