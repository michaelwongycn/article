package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/michaelwongycn/article/articleCommand/api/handler"
	repository "github.com/michaelwongycn/article/articleCommand/infrastructure/repository/article"
	usecase "github.com/michaelwongycn/article/articleCommand/service/article"
)

func loadDB() (*sql.DB, error) {
	err := godotenv.Load(".env")

	if err != nil {
		return nil, err
	}

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE"))
	db, err := sql.Open("mysql", dataSourceName)

	if err != nil {
		return nil, err
	}

	return db, nil
}

func main() {
	db, err := loadDB()

	if err != nil {
		fmt.Println("[ERROR] Error loading database|Message:" + err.Error())
	}

	articleRepository := repository.NewArticleMySQL(db)
	articleService := usecase.NewService(articleRepository)

	r := mux.NewRouter()
	articleRouter := r.PathPrefix("/article").Subrouter()

	handler.MakeMakeArticleHandlers(articleRouter, articleService)

	err = http.ListenAndServe("localhost:8001", r)
	fmt.Println("Runnning")
	if err != nil {
		fmt.Println(err.Error())
	}
}
