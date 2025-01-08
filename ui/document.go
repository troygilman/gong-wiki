package ui

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/yuin/goldmark"
)

type document struct {
	path     string
	html     string
	metadata documentMetadata
}

type DocumentManager struct {
	documents     map[string]*document
	documentOrder map[int]*document
}

func NewDocumentManager(fileSystem fs.FS) (DocumentManager, error) {
	dm := DocumentManager{
		documents:     make(map[string]*document),
		documentOrder: make(map[int]*document),
	}

	if err := fs.WalkDir(fileSystem, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		if filepath.Ext(path) != ".md" {
			return nil
		}

		source, err := fs.ReadFile(fileSystem, path)
		if err != nil {
			return err
		}

		sourceSplit := strings.Split(string(source), "---")

		var metadata documentMetadata
		if err := json.Unmarshal([]byte(sourceSplit[1]), &metadata); err != nil {
			return err
		}

		html, err := compileMarkdownToHTML([]byte(sourceSplit[2]))
		if err != nil {
			return err
		}

		document := &document{
			path:     strings.TrimSuffix(path, filepath.Ext(path)),
			html:     html,
			metadata: metadata,
		}
		dm.documents[document.path] = document
		dm.documentOrder[metadata.Position] = document
		return nil
	}); err != nil {
		return dm, err
	}

	return dm, nil
}

func (dm DocumentManager) getByPosition(position int) (*document, error) {
	document, ok := dm.documentOrder[position]
	if !ok {
		return nil, errors.New("document does not exist")
	}
	return document, nil
}

func (dm DocumentManager) getByPath(path string) (*document, error) {
	document, ok := dm.documents[path]
	if !ok {
		return nil, errors.New("document does not exist")
	}
	return document, nil
}

type documentMetadata struct {
	Label    string `json:"label"`
	Position int    `json:"position"`
}

func compileMarkdownToHTML(source []byte) (string, error) {
	var buf bytes.Buffer
	if err := goldmark.Convert(source, &buf); err != nil {
		return "", err
	}
	return buf.String(), nil

}
