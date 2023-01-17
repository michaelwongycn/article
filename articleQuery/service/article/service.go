package usecase

import (
	"github.com/michaelwongycn/article/articleQuery/entity"
)

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) ReadArticles(author, keyword string) ([]*entity.Article, error) {
	return s.repo.ReadArticles(author, keyword)
}
