package document

import (
	"io/fs"
	"path/filepath"
)

type Manager struct {
	documents     map[string]*Document
	documentOrder map[int]*Document
}

func NewManager(fileSystem fs.FS) (Manager, error) {
	dm := Manager{
		documents:     make(map[string]*Document),
		documentOrder: make(map[int]*Document),
	}

	parser := NewParser()

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

		doc, err := parser.Parse(path, source)
		if err != nil {
			return err
		}

		dm.documents[doc.Path()] = doc
		dm.documentOrder[doc.Metadata().Position] = doc
		return nil
	}); err != nil {
		return dm, err
	}

	return dm, nil
}

func (m Manager) AllPaths() []string {
	paths := make([]string, 0, len(m.documents))
	for path := range m.documents {
		paths = append(paths, path)
	}
	return paths
}

func (m Manager) GetByPosition(position int) *Document {
	document, ok := m.documentOrder[position]
	if !ok {
		return nil
	}
	return document
}

func (m Manager) GetByPath(path string) *Document {
	document, ok := m.documents[path]
	if !ok {
		return nil
	}
	return document
}
