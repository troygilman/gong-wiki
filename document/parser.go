package document

import (
	"bytes"
	"encoding/json"
	"errors"
	"path/filepath"
	"strings"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/text"
)

type Parser struct {
	parser   parser.Parser
	renderer renderer.Renderer
}

func NewParser() Parser {
	md := goldmark.New(
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRenderer(NewRenderer()),
	)
	return Parser{
		renderer: md.Renderer(),
		parser:   md.Parser(),
	}
}

func (p Parser) Parse(path string, source []byte) (*Document, error) {
	sourceSplit := strings.Split(string(source), "---")

	var metadata DocumentMetadata
	if err := json.Unmarshal([]byte(sourceSplit[1]), &metadata); err != nil {
		return nil, err
	}

	html, node, err := p.parseMarkdown([]byte(sourceSplit[2]))
	if err != nil {
		return nil, err
	}

	doc := &Document{
		path:     "/" + strings.TrimSuffix(path, filepath.Ext(path)),
		html:     html,
		metadata: metadata,
		node:     node,
	}
	return doc, nil
}

func (p Parser) parseMarkdown(source []byte) (html string, node *Node, err error) {
	reader := text.NewReader(source)
	astNode := p.parser.Parse(reader)

	root := &Node{}
	currentNode := root
	ast.Walk(astNode, func(astNode ast.Node, entering bool) (ast.WalkStatus, error) {
		if !entering {
			return ast.WalkContinue, nil
		}

		// Check if the node is a heading
		if heading, ok := astNode.(*ast.Heading); ok {
			headingLevel := heading.Level

			// Get heading text
			var buf bytes.Buffer
			for c := heading.FirstChild(); c != nil; c = c.NextSibling() {
				if text, ok := c.(*ast.Text); ok {
					buf.Write(text.Segment.Value(reader.Source()))
				}
			}
			headingText := buf.String()

			// Get heading id
			idAttr, ok := astNode.AttributeString("id")
			if !ok {
				return ast.WalkStop, errors.New("header is missing id")
			}
			id := string(idAttr.([]byte))

			// Create a new header node
			newNode := &Node{
				level: headingLevel,
				title: headingText,
				id:    id,
			}

			// Find the right parent for this node based on its level
			for currentNode != root && currentNode.level >= headingLevel {
				currentNode = currentNode.parent
			}

			// Set parent-child relationship
			newNode.parent = currentNode
			currentNode.children = append(currentNode.children, newNode)

			// Make this node the current node
			currentNode = newNode
		}

		return ast.WalkContinue, nil
	})

	writer := bytes.NewBuffer([]byte{})
	if err := p.renderer.Render(writer, source, astNode); err != nil {
		return html, node, err
	}

	return writer.String(), root, nil
}
