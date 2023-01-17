package repository

import "github.com/michaelwongycn/article/articleQuery/entity"

type ArticleRepository interface {
	ReadArticles(author, keyword string) ([]*entity.Article, error)
}
