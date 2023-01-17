package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/michaelwongycn/article/articleCommand/api/presenter"
	"github.com/michaelwongycn/article/articleCommand/entity"
	article "github.com/michaelwongycn/article/articleCommand/service/article"
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

func MakeMakeArticleHandlers(r *mux.Router, service article.UseCase) {
	r.Handle("", CreateArticle(service)).Methods(http.MethodPost)
}
