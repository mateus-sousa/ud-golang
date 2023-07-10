package repository

import (
	"api/src/models"
	"database/sql"
)

type Publishes struct {
	db *sql.DB
}

func NewPublishRepository(db *sql.DB) *Publishes {
	return &Publishes{db: db}
}

func (p *Publishes) Create(publish models.Publish) (uint64, error) {
	statement, err := p.db.Prepare("insert into publishes (title, content, author_id) values (?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(publish.Title, publish.Content, publish.AuthorID)
	if err != nil {
		return 0, err
	}
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return uint64(lastInsertId), nil
}

func (p *Publishes) GetPublish(publishId uint64) (models.Publish, error) {
	row, err := p.db.Query(
		`select p.*, u.nick from publishes p 
    			inner join users u on p.author_id = u.id 
				where p.id = ?`,
		publishId,
	)
	if err != nil {
		return models.Publish{}, err
	}
	defer row.Close()

	var publish models.Publish
	if row.Next() {
		if err = row.Scan(
			&publish.ID,
			&publish.Title,
			&publish.Content,
			&publish.AuthorID,
			&publish.Likes,
			&publish.CreatedAt,
			&publish.AuthorNick,
		); err != nil {
			return models.Publish{}, err
		}
	}

	return publish, nil
}

func (p *Publishes) GetPublishes(userId uint64) ([]models.Publish, error) {
	rows, err := p.db.Query(
		`select distinct p.*, u.nick from publishes p 
    			inner join users u on p.author_id = u.id 
				inner join followers f on p.author_id = f.user_id 
				where u.id = ? or f.follower_id = ?
				order by 1 desc`,
		userId, userId,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var publishes []models.Publish
	for rows.Next() {
		var publish models.Publish
		if err = rows.Scan(
			&publish.ID,
			&publish.Title,
			&publish.Content,
			&publish.AuthorID,
			&publish.Likes,
			&publish.CreatedAt,
			&publish.AuthorNick,
		); err != nil {
			return nil, err
		}
		publishes = append(publishes, publish)
	}

	return publishes, nil
}

func (p *Publishes) Update(publish models.Publish, userId uint64) error {
	statement, err := p.db.Prepare("update publishes set title = ?, content = ? where author_id = ?")
	if err != nil {
		return err
	}

	_, err = statement.Exec(publish.Title, publish.Content, userId)
	if err != nil {
		return err
	}

	return nil
}

func (p *Publishes) Delete(publishId uint64) error {
	statement, err := p.db.Prepare("delete from publishes where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()
	_, err = statement.Exec(publishId)
	if err != nil {
		return err
	}

	return nil
}

func (p *Publishes) GetPublishesByUser(userID uint64) ([]models.Publish, error) {
	rows, err := p.db.Query(
		"select p.*, u.nick from publishes p inner join users u on p.author_id = u.id where p.author_id = ?",
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var publishes []models.Publish
	for rows.Next() {
		var publish models.Publish
		if err = rows.Scan(
			&publish.ID,
			&publish.Title,
			&publish.Content,
			&publish.AuthorID,
			&publish.Likes,
			&publish.CreatedAt,
			&publish.AuthorNick,
		); err != nil {
			return nil, err
		}
		publishes = append(publishes, publish)
	}

	return publishes, nil
}
