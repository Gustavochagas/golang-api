package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

type users struct {
	db *sql.DB
}

func NewRepositoryForUser(db *sql.DB) *users {
	return &users{db}
}

func (u users) Create(user models.User) (uint64, error) {
	statement, erro := u.db.Prepare("insert into users (name, nick, email, password) values (?, ?, ?, ?)")
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	result, erro := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if erro != nil {
		return 0, erro
	}

	lastIDInsert, erro := result.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(lastIDInsert), nil
}

func (u users) Search(nameOrNick string) ([]models.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick)

	rows, erro := u.db.Query(
		"select id, name, nick, email, created_at from users where name LIKE ? or nick LIKE ?",
		nameOrNick, nameOrNick,
	)

	if erro != nil {
		return nil, erro
	}

	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User

		if erro = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); erro != nil {
			return nil, erro
		}

		users = append(users, user)
	}

	return users, nil
}

func (u users) SearchById(id uint64) (models.User, error) {
	rows, erro := u.db.Query(
		"select id, name, nick, email, created_at from users where id = ?", id,
	)
	if erro != nil {
		return models.User{}, erro
	}
	defer rows.Close()

	var user models.User

	if rows.Next() {
		if erro = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); erro != nil {
			return models.User{}, erro
		}
	}

	return user, nil
}

func (u users) Update(id uint64, user models.User) error {
	statement, erro := u.db.Prepare("update users set name = ?, nick = ?, email = ? where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(user.Name, user.Nick, user.Email, id); erro != nil {
		return erro
	}

	return nil
}

func (u users) Delete(id uint64) error {
	statement, erro := u.db.Prepare("delete from users where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro := statement.Exec(id); erro != nil {
		return erro
	}

	return nil
}

func (u users) SearchByEmail(email string) (models.User, error) {
	row, erro := u.db.Query("select id, password from users where email  = ?", email)
	if erro != nil {
		return models.User{}, erro
	}
	defer row.Close()

	var user models.User

	if row.Next() {
		if erro = row.Scan(&user.ID, &user.Password); erro != nil {
			return models.User{}, erro
		}
	}

	return user, nil
}

func (u users) Follow(userId, followerId uint64) error {
	statement, erro := u.db.Prepare("insert ignore into followers (user_id, follower_id) values (?, ?)")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(userId, followerId); erro != nil {
		return erro
	}

	return nil
}
