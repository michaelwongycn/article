package entity

import (
	"time"
)

type Article struct {
	ID      int
	Author  string
	Title   string
	Body    string
	Created time.Time
}

func NewArticle(author string, title string, body string) *Article {
	article := &Article{
		Author:  author,
		Title:   title,
		Body:    body,
		Created: time.Now(),
	}
	return article
}
