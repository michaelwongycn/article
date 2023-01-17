package usecase

import (
	"github.com/michaelwongycn/article/articleCommand/entity"
)

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) CreateArticle(author, title, body string) (*entity.Article, error) {
	article := entity.NewArticle(author, title, body)

	ID, err := s.repo.CreateArticle(article)

	if err != nil {
		return nil, err
	}

	article.ID = ID

	return article, nil
}
