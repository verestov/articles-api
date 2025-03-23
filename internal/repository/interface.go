package repository

import "Posts/internal/model"

// ArticleRepositoryInterface определяет методы, которые должен реализовать репозиторий
type ArticleRepositoryInterface interface {
	CreateArticle(article *model.Article) error
	GetAllArticles() ([]model.Article, error)
	GetArticleByID(id uint) (*model.Article, error)
	UpdateArticle(article *model.Article) error
	DeleteArticle(id uint) error
	GetFilteredAndSortedArticles(title, author, tag, sort string) ([]model.Article, error)
}
