package repository

import "github.com/michaelwongycn/article/entity"

type ArticleRepository interface {
	CreateArticle(article *entity.Article) (int, error)
	ReadArticles(author, keyword string) ([]*entity.Article, error)
	ReadArticlesFC() ([]*entity.Article, error)
}
