package generatestructure

type File struct {
	Name string // File name
	Confirmed bool // Check if the file has been confirmed or should be checked for circular errors
	Merge string // If set the file will be merged with a different file (basically appended at the bottom
}