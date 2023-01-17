package usecase

import "github.com/michaelwongycn/article/entity"

type Repository interface {
	CreateArticle(article *entity.Article) (int, error)
	ReadArticles(author, keyword string) ([]*entity.Article, error)
	ReadArticlesFC() ([]*entity.Article, error)
}

type UseCase interface {
	CreateArticle(author, title, body string) (*entity.Article, error)
	ReadArticles(author, keyword string) ([]*entity.Article, error)
	ReadArticlesFC() ([]*entity.Article, error)
}
