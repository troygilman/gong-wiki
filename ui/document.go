package ui

import (
	"bytes"
	"encoding/json"
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
	documents     map[string]document
	documentOrder map[int]string
}

func NewDocumentManager(fileSystem fs.FS) (DocumentManager, error) {
	dm := DocumentManager{
		documents:     make(map[string]document),
		documentOrder: make(map[int]string),
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

		document := document{
			path:     strings.TrimSuffix(path, filepath.Ext(path)),
			html:     html,
			metadata: metadata,
		}
		dm.documents[document.path] = document
		dm.documentOrder[metadata.Position] = document.path
		return nil
	}); err != nil {
		return dm, err
	}

	return dm, nil
}

type documentMetadata struct {
	Position int `json:"position"`
}

func compileMarkdownToHTML(source []byte) (string, error) {
	var buf bytes.Buffer
	if err := goldmark.Convert(source, &buf); err != nil {
		return "", err
	}
	return buf.String(), nil

}
