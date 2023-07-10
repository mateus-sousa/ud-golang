package repository

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

// Struct que recebe um ponteiro da conexao com o banco de dados
type Users struct {
	db *sql.DB
}

// Função que emula um metodo construtor
func NewUsersRepository(db *sql.DB) *Users {
	// Retorna por referencia um struct de Users com a conexao com o banco recebido
	return &Users{db}
}

func (u Users) Create(user models.User) (uint64, error) {
	statement, err := u.db.Prepare(
		"insert into users (name, nick, email, password) values (?, ?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}

	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastInsertID), nil
}

func (u Users) Get(nameOrNick string) ([]models.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick) // %nameOrNick%

	rows, err := u.db.Query(
		"select id, name, nick, email, created_at from users where name LIKE ? or nick LIKE ?",
		nameOrNick, nameOrNick)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User

		if err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (u Users) GetByID(ID uint64) (models.User, error) {
	rows, err := u.db.Query(
		"select id, name, nick, email, created_at from users where id = ?", ID,
	)
	if err != nil {
		return models.User{}, err
	}

	defer rows.Close()

	if rows.Next() {
		var user models.User
		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return models.User{}, err
		}

		return user, nil
	}

	return models.User{}, err
}

func (u Users) Update(ID uint64, user models.User) error {
	statement, err := u.db.Prepare("update users set name = ?, nick = ?, email = ? where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(user.Name, user.Nick, user.Email, ID); err != nil {
		return err
	}

	return nil
}

func (u Users) Delete(ID uint64) error {
	statement, err := u.db.Prepare("delete from users where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(ID); err != nil {
		return err
	}

	return nil
}

func (u Users) GetByEmail(email string) (models.User, error) {
	row, err := u.db.Query("select id, password from users where email = ?", email)
	if err != nil {
		return models.User{}, err
	}
	defer row.Close()
	var user models.User
	if row.Next() {
		if err := row.Scan(&user.ID, &user.Password); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

func (u Users) FollowUser(userID, followerID uint64) error {
	statement, err := u.db.Prepare(
		"insert ignore into followers (follower_id, user_id) values (?,?)",
	)
	if err != nil {
		return err
	}
	defer statement.Close()
	_, err = statement.Exec(followerID, userID)
	if err != nil {
		return err
	}

	return nil
}

func (u Users) StopFollowUser(userID, followerID uint64) error {
	statement, err := u.db.Prepare("delete from followers where user_id = ? and follower_id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()
	_, err = statement.Exec(userID, followerID)
	if err != nil {
		return err
	}

	return nil
}

func (u Users) GetFollowers(userID uint64) ([]models.User, error) {
	rows, err := u.db.Query(`
		select u.id, u.name, u.nick, u.email, u.created_at from users u
		inner join followers f on u.id = f.follower_id where f.user_id = ? 
		`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []models.User
	for rows.Next() {
		var user models.User
		err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (u Users) GetFollowing(userID uint64) ([]models.User, error) {
	rows, err := u.db.Query(`
	select u.id, u.name, u.nick, u.email, u.created_at from users u
	inner join followers f on u.id = f.user_id where f.follower_id = ?
	`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []models.User
	for rows.Next() {
		var user models.User
		err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (u Users) GetPassword(userID uint64) (string, error) {
	row, err := u.db.Query("select password from users where id = ?", userID)
	if err != nil {
		return "", nil
	}
	defer row.Close()
	var user models.User
	if row.Next() {
		err := row.Scan(&user.Password)
		if err != nil {
			return "", nil
		}
	}

	return user.Password, nil
}

func (u Users) UpdatePassword(userID uint64, passwordHash string) error {
	statement, err := u.db.Prepare("update users set password = ? where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(passwordHash, userID)
	if err != nil {
		return err
	}

	return nil
}
