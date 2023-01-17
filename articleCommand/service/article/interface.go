package usecase

import "github.com/michaelwongycn/article/articleCommand/entity"

type Repository interface {
	CreateArticle(article *entity.Article) (int, error)
}

type UseCase interface {
	CreateArticle(author, title, body string) (*entity.Article, error)
}
