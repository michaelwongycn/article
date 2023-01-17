package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/michaelwongycn/article/api/presenter"
	"github.com/michaelwongycn/article/api/util"
	"github.com/michaelwongycn/article/entity"
	article "github.com/michaelwongycn/article/service/article"
)

func CreateArticle(service article.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var input struct {
			Author string `json:"author"`
			Title  string `json:"title"`
			Body   string `json:"body"`
		}
		errorMessage := "Please contact system admin"
		err := json.NewDecoder(r.Body).Decode(&input)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(errorMessage))
			return
		}

		var articles []*entity.Article

		article, err := service.CreateArticle(input.Author, input.Title, input.Body)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		articles = append(articles, article)

		result := &presenter.Article{
			Data: articles,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		err = json.NewEncoder(w).Encode(result)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
		}
	})
}

func ReadArticles(service article.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Please contact system admin"
		var data []*entity.Article
		var err error

		keyword := r.URL.Query().Get("query")
		author := r.URL.Query().Get("author")

		data, found := util.GetArticleCacheByKey(author, keyword)

		if !found {
			data, err = service.ReadArticles(author, keyword)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(errorMessage))
				return
			}

			util.SetArticleCacheByKey(author, keyword, data)
		}

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
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
			log.Println("[ERROR] Error encoding result|Message:" + err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
		}
	})
}

func ReadArticlesFC(service article.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Please contact system admin"
		var data []*entity.Article
		var err error

		keyword := r.URL.Query().Get("query")
		author := r.URL.Query().Get("author")

		data, found := util.GetArticleCache()

		if !found {
			data, err = service.ReadArticlesFC()

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(errorMessage))
				return
			}

			util.SetArticleCache(data)
		}

		var articles []*entity.Article

		for i := range data {
			if (strings.Contains(data[i].Author, author)) && ((strings.Contains(data[i].Title, keyword)) || (strings.Contains(data[i].Body, keyword))) {
				articles = append(articles, data[i])
			}
		}

		if articles == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(errorMessage))
			return
		}

		result := &presenter.Article{
			Data: articles,
		}

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(result)

		if err != nil {
			log.Println("[ERROR] Error encoding result|Message:" + err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
		}
	})
}

func MakeMakeArticleHandlers(r *mux.Router, service article.UseCase) {
	r.Handle("", CreateArticle(service)).Methods(http.MethodPost)
	// r.Handle("", ReadArticles(service)).Methods(http.MethodGet)
	r.Handle("", ReadArticlesFC(service)).Methods(http.MethodGet)
}
