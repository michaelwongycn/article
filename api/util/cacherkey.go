package util

import (
	"fmt"

	"github.com/michaelwongycn/article/entity"
	"github.com/patrickmn/go-cache"
)

func SetArticleCacheByKey(author, keyword string, data []*entity.Article) {
	fmt.Println("Write Cache")
	Cache.Set(author+"~"+keyword, data, cache.NoExpiration)
}

func GetArticleCacheByKey(author, keyword string) ([]*entity.Article, bool) {
	var articles []*entity.Article
	data, found := Cache.Get(author + "~" + keyword)

	if found {
		fmt.Println("Cache Found")
		articles = data.([]*entity.Article)
	}
	return articles, found
}
