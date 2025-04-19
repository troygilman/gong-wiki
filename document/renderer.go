package document

import (
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/util"
)

type NodeRenderer struct {
	renderer *html.Renderer
}

func NewRenderer() renderer.Renderer {
	nodeRenderer := &NodeRenderer{
		renderer: &html.Renderer{
			Config: html.NewConfig(),
		},
	}
	return renderer.NewRenderer(renderer.WithNodeRenderers(util.Prioritized(nodeRenderer, 100)))
}

func (r *NodeRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	r.renderer.RegisterFuncs(reg)
	reg.Register(ast.KindHeading, r.renderHeading)
	reg.Register(ast.KindFencedCodeBlock, r.renderFencedCodeBlock)
}

func (r *NodeRenderer) writeLines(w util.BufWriter, source []byte, n ast.Node) {
	l := n.Lines().Len()
	for i := 0; i < l; i++ {
		line := n.Lines().At(i)
		r.renderer.Writer.RawWrite(w, line.Value(source))
	}
}

func (r *NodeRenderer) renderHeading(
	w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	n := node.(*ast.Heading)
	if entering {
		_, _ = w.WriteString("<h")
		_ = w.WriteByte("0123456"[n.Level])
		if n.Attributes() != nil {
			html.RenderAttributes(w, node, html.HeadingAttributeFilter)
		}
		_ = w.WriteByte('>')
		idBytes, ok := node.AttributeString("id")
		if ok && n.Level != 1 {
			id := string(idBytes.([]byte))
			_, _ = w.WriteString("<a href=\"#")
			_, _ = w.WriteString(id)
			_, _ = w.WriteString("\" class=\"header-link\">#</a>")
		}
	} else {
		_, _ = w.WriteString("</h")
		_ = w.WriteByte("0123456"[n.Level])
		_, _ = w.WriteString(">\n")
	}
	return ast.WalkContinue, nil
}

func (r *NodeRenderer) renderFencedCodeBlock(
	w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	n := node.(*ast.FencedCodeBlock)
	if entering {
		_, _ = w.WriteString("<div class=\"code-container\"><pre><code")
		language := n.Language(source)
		if language != nil {
			_, _ = w.WriteString(" class=\"language-")
			r.renderer.Writer.Write(w, language)
			_, _ = w.WriteString("\"")
		}
		_ = w.WriteByte('>')
		r.writeLines(w, source, n)
	} else {
		_, _ = w.WriteString("</code></pre></div>\n")
	}
	return ast.WalkContinue, nil
}
