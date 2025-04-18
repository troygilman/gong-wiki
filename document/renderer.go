package document

import (
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/util"
)

type NodeRenderer struct {
	*html.Renderer
}

func NewRenderer() renderer.Renderer {
	nodeRenderer := &NodeRenderer{
		Renderer: &html.Renderer{
			Config: html.NewConfig(),
		},
	}
	return renderer.NewRenderer(renderer.WithNodeRenderers(util.Prioritized(nodeRenderer, 1000)))
}

func (r *NodeRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	r.Renderer.RegisterFuncs(reg)
	reg.Register(ast.KindFencedCodeBlock, r.renderFencedCodeBlock)
}

func (r *NodeRenderer) writeLines(w util.BufWriter, source []byte, n ast.Node) {
	l := n.Lines().Len()
	for i := 0; i < l; i++ {
		line := n.Lines().At(i)
		r.Writer.RawWrite(w, line.Value(source))
	}
}

func (r *NodeRenderer) renderFencedCodeBlock(
	w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	n := node.(*ast.FencedCodeBlock)
	if entering {
		_, _ = w.WriteString("<div class=\"code-container\"><pre><code")
		language := n.Language(source)
		if language != nil {
			_, _ = w.WriteString(" class=\"language-")
			r.Writer.Write(w, language)
			_, _ = w.WriteString("\"")
		}
		_ = w.WriteByte('>')
		r.writeLines(w, source, n)
	} else {
		_, _ = w.WriteString("</code></pre></div>\n")
	}
	return ast.WalkContinue, nil
}
