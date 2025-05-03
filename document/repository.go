package document

import "database/sql"

type Repository struct {
	db *sql.DB
}

func NewRepository() Repository {
	db, err := sql.Open("sqlite3", ":memory:")
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
	_, err := repository.db.Exec(`create table document (name text not null, id text not null, content text not null, PRIMARY KEY(name, id));`)
	return err
}

func (repository Repository) AddDocument(doc *Document) error {
	tx, err := repository.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("insert into document(name, id, content) values(?, ?, ?)")
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

	_, err := stmt.Exec(doc.Path(), node.ID(), node.html)
	return err
}
