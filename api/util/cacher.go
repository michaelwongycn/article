package util

import (
	"fmt"
	"time"

	"github.com/michaelwongycn/article/entity"
	"github.com/patrickmn/go-cache"
)

var Cache = cache.New(5*time.Minute, 5*time.Minute)
var key = "article"

func SetArticleCache(data []*entity.Article) {
	fmt.Println("Write Cache")
	Cache.Set(key, data, cache.NoExpiration)
}

func GetArticleCache() ([]*entity.Article, bool) {
	var articles []*entity.Article
	data, found := Cache.Get(key)

	if found {
		fmt.Println("Cache Found")
		articles = data.([]*entity.Article)
	}
	return articles, found
}
