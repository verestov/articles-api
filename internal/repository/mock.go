package repository

import (
	"Posts/internal/model"

	"github.com/stretchr/testify/mock"
)

// Убедимся, что MockArticleRepository реализует ArticleRepositoryInterface
var _ ArticleRepositoryInterface = (*MockArticleRepository)(nil)

// MockArticleRepository - мок-репозиторий для тестов
type MockArticleRepository struct {
	mock.Mock
}

func (m *MockArticleRepository) CreateArticle(article *model.Article) error {
	args := m.Called(article)
	return args.Error(0)
}

func (m *MockArticleRepository) GetAllArticles() ([]model.Article, error) {
	args := m.Called()
	return args.Get(0).([]model.Article), args.Error(1)
}

func (m *MockArticleRepository) GetArticleByID(id uint) (*model.Article, error) {
	args := m.Called(id)
	return args.Get(0).(*model.Article), args.Error(1)
}

func (m *MockArticleRepository) UpdateArticle(article *model.Article) error {
	args := m.Called(article)
	return args.Error(0)
}

func (m *MockArticleRepository) DeleteArticle(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockArticleRepository) GetFilteredAndSortedArticles(title, author, tag, sort string) ([]model.Article, error) {
	args := m.Called(title, author, tag, sort)
	return args.Get(0).([]model.Article), args.Error(1)
}
