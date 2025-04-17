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

func (d Document) Node() *Node {
	return d.node
}

type DocumentMetadata struct {
	Label    string `json:"label"`
	Position int    `json:"position"`
}

type Node struct {
	title    string
	id       string
	level    int
	parent   *Node
	children []*Node
}

func (node Node) Title() string {
	return node.title
}

func (node Node) ID() string {
	return node.id
}

func (node Node) Level() int {
	return node.level
}

func (node Node) Children() []*Node {
	children := make([]*Node, len(node.children))
	for i, child := range node.children {
		children[i] = child
	}
	return children
}
