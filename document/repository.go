package document

import (
	"database/sql"
)

type Repository struct {
	db *sql.DB
}

func NewRepository() Repository {
	db, err := sql.Open("sqlite3", "file:application.db?mode=memory")
	// db, err := sql.Open("sqlite3", "file:./tmp/application.db")
	if err != nil {
		panic(err)
	}

	repository := Repository{
		db: db,
	}

	if err := repository.Migrate(); err != nil {
		panic(err)
	}

	return repository
}

func (repository Repository) Migrate() error {
	_, err := repository.db.Exec("CREATE VIRTUAL TABLE IF NOT EXISTS document_chunk USING fts5(id, name, document, content)")
	return err
}

func (repository Repository) AddDocument(doc *Document) error {
	tx, err := repository.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("INSERT INTO document_chunk (id, name, document, content) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	if err := repository.addNode(stmt, doc.node, doc); err != nil {
		return err
	}

	return tx.Commit()
}

func (repository Repository) addNode(stmt *sql.Stmt, node *Node, doc *Document) error {
	for _, child := range node.children {
		if err := repository.addNode(stmt, child, doc); err != nil {
			return err
		}
	}

	_, err := stmt.Exec(doc.Path()+"#"+node.ID(), node.title, doc.metadata.Label, node.content)
	return err
}

func (repository Repository) SearchDocumentChunk(query string) (chunks []DocumentChunk, err error) {
	rows, err := repository.db.Query("SELECT id, name, document FROM document_chunk WHERE content MATCH ? ORDER BY rank", query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var chunk DocumentChunk
	for rows.Next() {
		if err := rows.Scan(&chunk.ID, &chunk.Name, &chunk.Document); err != nil {
			return nil, err
		}
		chunks = append(chunks, chunk)
	}

	return chunks, nil
}
