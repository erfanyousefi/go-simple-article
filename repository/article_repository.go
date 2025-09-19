package repository

import (
	"github.com/erfanyousefi/simple-article/database"
	"github.com/erfanyousefi/simple-article/models"
)

type ArticleRepository struct{}

func (ArticleRepository) Create(article *models.Article) error {
	return database.GORMDB.Create(article).Error
}

func (ArticleRepository) Update(article *models.Article) error {
	return database.GORMDB.Save(&models.Article{}).Error
}

func (ArticleRepository) Delete(article *models.Article, id uint) error {
	return database.GORMDB.Delete(&models.Article{}, id).Error
}
func (ArticleRepository) GetByID(id uint) (models.Article, error) {
	var article models.Article
	err := database.GORMDB.First(&article, id).Error
	return article, err
}
func (ArticleRepository) GetAll() ([]models.Article, error) {
	var articles []models.Article
	err := database.GORMDB.Find(&articles).Error
	return articles, err
}
