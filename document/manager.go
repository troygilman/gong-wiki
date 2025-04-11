package document

import (
	"errors"
	"io/fs"
	"path/filepath"
)

type DocumentManager struct {
	documents     map[string]*Document
	documentOrder map[int]*Document
	headers       map[string]string
}

func NewDocumentManager(fileSystem fs.FS) (DocumentManager, error) {
	dm := DocumentManager{
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

func (dm DocumentManager) GetByPosition(position int) (*Document, error) {
	document, ok := dm.documentOrder[position]
	if !ok {
		return nil, errors.New("document does not exist")
	}
	return document, nil
}

func (dm DocumentManager) GetByPath(path string) (*Document, error) {
	document, ok := dm.documents[path]
	if !ok {
		return nil, errors.New("document does not exist")
	}
	return document, nil
}
