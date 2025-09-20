package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/erfanyousefi/simple-article/models"
	"github.com/erfanyousefi/simple-article/repository"
	"github.com/gorilla/mux"
)

type ArticleHandler struct {
	Repo repository.ArticleRepository
}

func (handler ArticleHandler) CreateArticle(res http.ResponseWriter, req *http.Request) {
	var article models.Article
	if err := json.NewDecoder(req.Body).Decode(&article); err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	if err := handler.Repo.Create(&article); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	res.WriteHeader(http.StatusCreated)
	json.NewEncoder(res).Encode(article)
}

func (handler ArticleHandler) GetArticleById(res http.ResponseWriter, req *http.Request) {
	idStr := mux.Vars(req)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(res, "Invalid article ID", http.StatusBadRequest)
		return
	}
	article, err := handler.Repo.GetByID(uint(id))
	if err != nil {
		http.Error(res, err.Error(), http.StatusNotFound)
		return
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(article)
}

func (handler ArticleHandler) GetAllArticles(res http.ResponseWriter, req *http.Request) {
	articles, err := handler.Repo.GetAll()
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(articles)
}

func (handler ArticleHandler) UpdateArticle(res http.ResponseWriter, req *http.Request) {
	idStr := mux.Vars(req)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	var article models.Article
	if err := json.NewDecoder(req.Body).Decode(&article); err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	article.ID = uint(id)
	if err := handler.Repo.Update(&article); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(res).Encode(article)
}

func (handler ArticleHandler) DeleteArticle(res http.ResponseWriter, req *http.Request) {
	idStr := mux.Vars(req)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(res, "Invalid article ID", http.StatusBadRequest)
		return
	}
	article, err := handler.Repo.GetByID(uint(id))
	if err != nil {
		http.Error(res, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(res).Encode(article)

}
