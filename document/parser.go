package document

import (
	"bytes"
	"encoding/json"
	"path/filepath"
	"strings"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
)

type Parser struct {
	renderer Renderer
	markdown goldmark.Markdown
}

func NewParser() Parser {
	renderer := NewRenderer()
	markdown := goldmark.New(
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRenderer(renderer),
	)
	return Parser{
		renderer: renderer,
		markdown: markdown,
	}
}

func (p Parser) Parse(path string, source []byte) (doc Document, err error) {
	sourceSplit := strings.Split(string(source), "---")

	var metadata DocumentMetadata
	if err := json.Unmarshal([]byte(sourceSplit[1]), &metadata); err != nil {
		return doc, err
	}

	buf := bytes.NewBuffer([]byte{})
	if err := p.markdown.Convert([]byte(sourceSplit[2]), buf); err != nil {
		return doc, err
	}

	doc = Document{
		path:     strings.TrimSuffix(path, filepath.Ext(path)),
		html:     buf.String(),
		metadata: metadata,
		node:     p.renderer.Node(),
	}
	return doc, nil
}
