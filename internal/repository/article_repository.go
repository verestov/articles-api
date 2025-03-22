package repository

import (
	"Posts/internal/model"
	"fmt"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type ArticleRepository struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) *ArticleRepository {
	return &ArticleRepository{db: db}
}

// GetAll
func (r *ArticleRepository) GetAllArticles() ([]model.Article, error) {
	var articles []model.Article
	if err := r.db.Find(&articles).Error; err != nil {
		return nil, err
	}

	return articles, nil
}

// Create
func (r *ArticleRepository) CreateArticle(article *model.Article) error {
	return r.db.Create(article).Error
}

// GetById
func (r *ArticleRepository) GetArticleByID(id uint) (*model.Article, error) {
	var article model.Article
	if err := r.db.First(&article, id).Error; err != nil {
		return nil, err
	}
	return &article, nil
}

// Update
func (r *ArticleRepository) UpdateArticle(article *model.Article) error {
	// Обновляем только те поля, которые были изменены
	return r.db.Model(article).Updates(map[string]interface{}{
		"title":   article.Title,
		"content": article.Content,
		"author":  article.Author,
		"tags":    article.Tags,
	}).Error
}

// Delete
func (r *ArticleRepository) DeleteArticle(id uint) error {
	return r.db.Delete(&model.Article{}, id).Error
}

// Метод для получения с фильтрацией и сортировкой
func (r *ArticleRepository) GetFilteredAndSortedArticles(title, author, tag, sort string) ([]model.Article, error) {
	var articles []model.Article
	// Строим запрос
	query := r.db.Model(model.Article{})

	// Фильтрация по title
	if title != "" {
		query = query.Where("title ILIKE ?", "%"+title+"%")
	}

	// Фильтрация по author
	if author != "" {
		query = query.Where("author ILIKE ?", "%"+author+"%")
	}

	// Фильтраци по tags
	if tag != "" {
		query = query.Where("tags @> ?", pq.Array([]string{tag}))
	}

	// Сортировка
	if sort != "" {
		if sort[0] == '-' {
			// По возрастанию
			query = query.Order(fmt.Sprintf("%s DESC", sort[1:]))
		} else {
			// По убыванию
			query = query.Order(fmt.Sprintf("%s ASC", sort))
		}
	}

	if err := query.Find(&articles).Error; err != nil {
		return nil, err
	}

	return articles, nil
}
