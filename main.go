package main

import (
	"log"
	"net/http"

	"github.com/erfanyousefi/simple-article/config"
	"github.com/erfanyousefi/simple-article/database"
	"github.com/erfanyousefi/simple-article/handlers"
	"github.com/erfanyousefi/simple-article/models"
	"github.com/erfanyousefi/simple-article/repository"
	"github.com/gorilla/mux"
)

func main() {
	//load environment variable configs
	cfg := config.LoadConfig()

	//connected to db
	database.ConnectGORM(cfg)
	log.Println("all connections are true")

	//migrations
	database.GORMDB.AutoMigrate(&models.Article{})

	//init handler
	repo := repository.ArticleRepository{}
	handler := handlers.ArticleHandler{Repo: repo}

	//routes
	router := mux.NewRouter()

	router.HandleFunc("/articles", handler.GetAllArticles).Methods("GET")
	router.HandleFunc("/articles/{id}", handler.GetArticleById).Methods("GET")
	router.HandleFunc("/articles", handler.CreateArticle).Methods("POST")
	router.HandleFunc("/articles/{id}", handler.UpdateArticle).Methods("PUT")
	router.HandleFunc("/articles/{id}", handler.DeleteArticle).Methods("DELETE")

	//run http server
	log.Printf("Server running on port: %s \n", cfg.ServerPort)
	log.Fatal(http.ListenAndServe(":"+cfg.ServerPort, router))
}
