package repository

import "github.com/michaelwongycn/article/articleCommand/entity"

type ArticleRepository interface {
	CreateArticle(article *entity.Article) (int, error)
}
