package document

type Document struct {
	path     string
	html     string
	metadata DocumentMetadata
	node     *Node
}

func (d Document) Path() string {
	return d.path
}

func (d Document) Html() string {
	return d.html
}

func (d Document) Metadata() DocumentMetadata {
	return d.metadata
}

type DocumentMetadata struct {
	Label    string `json:"label"`
	Position int    `json:"position"`
}

type Node struct {
	id       string
	children []*Node
}
