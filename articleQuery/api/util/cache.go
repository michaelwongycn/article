package util

import (
	"fmt"
	"time"

	"github.com/michaelwongycn/article/articleQuery/entity"
	"github.com/patrickmn/go-cache"
)

var Cache = cache.New(5*time.Minute, 5*time.Minute)

func SetArticleCache(author, keyword string, data []*entity.Article) {
	fmt.Println("Write Cache")
	Cache.Set(author+"~~~~"+keyword, data, cache.NoExpiration)
}

func GetArticleCache(author, keyword string) ([]*entity.Article, bool) {
	var articles []*entity.Article
	data, found := Cache.Get(author + "~~~~" + keyword)

	if found {
		fmt.Println("Cache Found")
		articles = data.([]*entity.Article)
	}
	return articles, found
}
