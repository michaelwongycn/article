package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/michaelwongycn/article/articleQuery/api/presenter"
	"github.com/michaelwongycn/article/articleQuery/api/util"
	"github.com/michaelwongycn/article/articleQuery/entity"
	article "github.com/michaelwongycn/article/articleQuery/service/article"
)

func ReadArticles(service article.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Please contact system admin"
		var data []*entity.Article
		var err error

		keyword := r.URL.Query().Get("query")
		author := r.URL.Query().Get("author")

		data, found := util.GetArticleCache(author, keyword)

		if !found {
			data, err = service.ReadArticles(author, keyword)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(errorMessage))
				return
			}

			util.SetArticleCache(author, keyword, data)
		}

		if data == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(errorMessage))
			return
		}

		var articles []*entity.Article
		articles = append(articles, data...)

		result := &presenter.Article{
			Data: articles,
		}

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(result)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
		}
	})
}

func MakeMakeArticleHandlers(r *mux.Router, service article.UseCase) {
	r.Handle("", ReadArticles(service)).Methods(http.MethodGet)
}
