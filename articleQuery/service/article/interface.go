package usecase

import "github.com/michaelwongycn/article/articleQuery/entity"

type Repository interface {
	ReadArticles(author, keyword string) ([]*entity.Article, error)
}

type UseCase interface {
	ReadArticles(author, keyword string) ([]*entity.Article, error)
}
