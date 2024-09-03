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

func (repository Publications) SearchById(publicationId uint64) (models.Publication, error) {
	row, erro := repository.db.Query("select p.*, u.nick from publications p inner join users u on u.id = p.author_id where p.id = ?", publicationId)
	if erro != nil {
		return models.Publication{}, erro
	}
	defer row.Close()

	var publication models.Publication

	if row.Next() {
		if erro = row.Scan(
			&publication.ID,
			&publication.Title,
			&publication.Content,
			&publication.AuthorID,
			&publication.Likes,
			&publication.CreatedAt,
			&publication.AuthorNick,
		); erro != nil {
			return models.Publication{}, erro
		}
	}

	return publication, nil
}

func (repository Publications) Search(userId uint64) ([]models.Publication, error) {
	rows, erro := repository.db.Query(`
			select distinct p.*, u.nick from publications p 
			inner join users u on u.id = p.author_id 
			inner join followers f on p.author_id = f.user_id 
			where u.id = ? or f.follow_id = ?
			order by 1 desc`,
		userId, userId,
	)
	if erro != nil {
		return nil, erro
	}
	defer rows.Close()

	var publications []models.Publication

	for rows.Next() {
		var publication models.Publication

		if erro = rows.Scan(
			&publication.ID,
			&publication.Title,
			&publication.Content,
			&publication.AuthorID,
			&publication.Likes,
			&publication.CreatedAt,
			&publication.AuthorNick,
		); erro != nil {
			return nil, erro
		}

		publications = append(publications, publication)
	}

	return publications, nil
}

func (repository Publications) Update(publicationId uint64, publication models.Publication) error {
	statament, erro := repository.db.Prepare("update publications set title = ?, content = ? where id = ?")
	if erro != nil {
		return erro
	}
	defer statament.Close()

	if _, erro = statament.Exec(publication.Title, publication.Content, publicationId); erro != nil {
		return erro
	}

	return nil
}

func (repository Publications) Delete(publicationId uint64) error {
	statament, erro := repository.db.Prepare("delete from publications where id = ?")
	if erro != nil {
		return erro
	}
	defer statament.Close()

	if _, erro = statament.Exec(publicationId); erro != nil {
		return erro
	}

	return nil
}

func (repository Publications) SearchByUser(userId uint64) ([]models.Publication, error) {
	rows, erro := repository.db.Query("select p.*, u.nick from publications p join users u on u.id = p.author_id where p.author_id = ?", userId)
	if erro != nil {
		return nil, erro
	}
	defer rows.Close()

	var publications []models.Publication

	for rows.Next() {
		var publication models.Publication

		if erro = rows.Scan(
			&publication.ID,
			&publication.Title,
			&publication.Content,
			&publication.AuthorID,
			&publication.Likes,
			&publication.CreatedAt,
			&publication.AuthorNick,
		); erro != nil {
			return nil, erro
		}

		publications = append(publications, publication)
	}

	return publications, nil
}

func (repository Publications) Like(publicationID uint64) error {
	statament, erro := repository.db.Prepare("update publications set likes = likes + 1 where id = ?")
	if erro != nil {
		return erro
	}
	defer statament.Close()

	if _, erro = statament.Exec(publicationID); erro != nil {
		return erro
	}

	return nil
}

func (repository Publications) Unlike(publicationID uint64) error {
	statament, erro := repository.db.Prepare("update publications set likes = CASE WHEN likes > 0 THEN likes - 1 ELSE likes END where id = ?")
	if erro != nil {
		return erro
	}
	defer statament.Close()

	if _, erro = statament.Exec(publicationID); erro != nil {
		return erro
	}

	return nil
}
