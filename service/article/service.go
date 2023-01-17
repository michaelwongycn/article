package usecase

import (
	"github.com/michaelwongycn/article/entity"
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

func (s *Service) ReadArticles(author, keyword string) ([]*entity.Article, error) {
	return s.repo.ReadArticles(author, keyword)
}

func (s *Service) ReadArticlesFC() ([]*entity.Article, error) {
	return s.repo.ReadArticlesFC()
}
