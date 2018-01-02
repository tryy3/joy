package dom

type Text interface {
	CharacterData

	SplitText(offset uint) (t Text)

	WholeText() (wholeText string)
}
