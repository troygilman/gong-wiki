package document

import (
	"io"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/renderer"
)

type Renderer struct {
	renderer.Renderer
	node *Node
}

func NewRenderer() Renderer {
	return Renderer{Renderer: goldmark.DefaultRenderer()}
}

func (renderer Renderer) Node() *Node {
	return renderer.node
}

func (renderer Renderer) Render(w io.Writer, source []byte, astNode ast.Node) error {
	renderer.node = &Node{}
	renderer.searchNode(astNode.FirstChild(), renderer.node)
	return renderer.Renderer.Render(w, source, astNode)
}

func (renderer Renderer) searchNode(astNode ast.Node, parent *Node) {
	if astNode == nil {
		return
	}

	if id, ok := astNode.AttributeString("id"); ok {
		node := &Node{}
		node.id = string(id.([]byte))
		parent.children = append(parent.children, node)
		renderer.searchNode(astNode.FirstChild(), node)
	}

	renderer.searchNode(astNode.NextSibling(), parent)
}
