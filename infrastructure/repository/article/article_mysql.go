package repository

import (
	"database/sql"

	"github.com/michaelwongycn/article/entity"
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

func (r *ArticleMySQL) ReadArticles(author, keyword string) ([]*entity.Article, error) {
	query := `SELECT * FROM article WHERE author LIKE ? AND ((title LIKE ?) OR (body LIKE ?)) ORDER BY created DESC`

	author = "%" + author + "%"
	keyword = "%" + keyword + "%"

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return nil, err
	}

	var articles []*entity.Article

	rows, err := stmt.Query(author, keyword, keyword)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var article entity.Article

		err = rows.Scan(&article.ID, &article.Author, &article.Title, &article.Body, &article.Created)

		if err != nil {
			return nil, err
		}
		articles = append(articles, &article)
	}

	err = stmt.Close()

	if err != nil {
		return nil, err
	}

	return articles, nil
}

func (r *ArticleMySQL) ReadArticlesFC() ([]*entity.Article, error) {
	query := `SELECT * FROM article ORDER BY created DESC`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return nil, err
	}

	var articles []*entity.Article

	rows, err := stmt.Query()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var article entity.Article

		err = rows.Scan(&article.ID, &article.Author, &article.Title, &article.Body, &article.Created)

		if err != nil {
			return nil, err
		}
		articles = append(articles, &article)
	}

	err = stmt.Close()

	if err != nil {
		return nil, err
	}

	return articles, nil
}
