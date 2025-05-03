package document

import (
	"io/fs"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

type Manager struct {
	documents     map[string]*Document
	documentOrder map[int]*Document
	Repository    Repository
}

func NewManager(fileSystem fs.FS) (Manager, error) {
	dm := Manager{
		documents:     make(map[string]*Document),
		documentOrder: make(map[int]*Document),
		Repository:    NewRepository(),
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

	for _, doc := range dm.documents {
		if err := dm.Repository.AddDocument(doc); err != nil {
			return dm, err
		}
	}

	return dm, nil
}

func (m Manager) AllPaths() []string {
	paths := make([]string, len(m.documents))
	for order, document := range m.documentOrder {
		paths[order-1] = document.path
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
