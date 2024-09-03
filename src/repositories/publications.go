package repositories

import (
	"api/src/models"
	"database/sql"
)

type Publications struct {
	db *sql.DB
}

func NewRepositoryForPublications(db *sql.DB) *Publications {
	return &Publications{db}
}

func (repository Publications) Create(publication models.Publication) (uint64, error) {
	statament, erro := repository.db.Prepare("insert into publications (title, content, author_id) values (?, ?, ?)")
	if erro != nil {
		return 0, erro
	}
	defer statament.Close()

	result, erro := statament.Exec(publication.Title, publication.Content, publication.AuthorID)
	if erro != nil {
		return 0, erro
	}

	lastIdInsert, erro := result.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(lastIdInsert), nil
}
