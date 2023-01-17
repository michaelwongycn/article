package repository

import (
	"database/sql"

	"github.com/michaelwongycn/article/articleCommand/entity"
)

type ArticleMySQL struct {
	db *sql.DB
}

func NewArticleMySQL(db *sql.DB) *ArticleMySQL {
	return &ArticleMySQL{
		db: db,
	}
}

func (r *ArticleMySQL) CreateArticle(article *entity.Article) (int, error) {
	query := `INSERT INTO article(author, title, body, created) VALUES (?, ?, ?, ?)`

	stmt, err := r.db.Prepare(query)

	if err != nil {
		return 0, err
	}

	result, err := stmt.Exec(
		article.Author, article.Title, article.Body, article.Created,
	)

	if err != nil {
		return 0, err
	}

	err = stmt.Close()

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return int(id), nil
}
