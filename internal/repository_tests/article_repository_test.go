package repository_test

import (
	"Posts/internal/model"
	"Posts/internal/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	_ "github.com/stretchr/testify/mock"
)

func TestCreateArticle(t *testing.T) {
	mockRepo := new(repository.MockArticleRepository)
	article := &model.Article{
		Title:   "Test Article",
		Content: "This is a test article",
		Author:  "John Doe",
	}

	// Ожидаем, что метод CreateArticle будет вызван один раз с таким объектом
	mockRepo.On("CreateArticle", article).Return(nil)

	// Вызываем метод
	err := mockRepo.CreateArticle(article)

	// Проверяем, что ошибок не было
	assert.NoError(t, err)

	// Проверяем, что метод был вызван с нужными параметрами
	mockRepo.AssertExpectations(t)
}

func TestGetAllArticles(t *testing.T) {
	mockRepo := new(repository.MockArticleRepository)
	articles := []model.Article{
		{Title: "Article 1", Content: "Content 1", Author: "Author 1"},
		{Title: "Article 2", Content: "Content 2", Author: "Author 2"},
	}

	// Ожидаем, что метод GetAllArticles будет вызван
	mockRepo.On("GetAllArticles").Return(articles, nil)

	// Вызываем метод
	result, err := mockRepo.GetAllArticles()

	// Проверяем, что ошибок не было
	assert.NoError(t, err)

	// Проверяем, что возвращены правильные статьи
	assert.Equal(t, 2, len(result))
	assert.Equal(t, "Article 1", result[0].Title)

	// Проверяем, что метод был вызван
	mockRepo.AssertExpectations(t)
}

func TestUpdateArticle(t *testing.T) {
	mockRepo := new(repository.MockArticleRepository)
	article := &model.Article{
		ID:      1,
		Title:   "Updated Article",
		Content: "Updated content",
		Author:  "John Doe",
	}

	// Ожидаем, что метод UpdateArticle будет вызван
	mockRepo.On("UpdateArticle", article).Return(nil)

	// Вызываем метод
	err := mockRepo.UpdateArticle(article)

	// Проверяем, что ошибок не было
	assert.NoError(t, err)

	// Проверяем, что метод был вызван
	mockRepo.AssertExpectations(t)
}

func TestDeleteArticle(t *testing.T) {
	mockRepo := new(repository.MockArticleRepository)
	articleID := uint(1)

	// Ожидаем, что метод DeleteArticle будет вызван
	mockRepo.On("DeleteArticle", articleID).Return(nil)

	// Вызываем метод
	err := mockRepo.DeleteArticle(articleID)

	// Проверяем, что ошибок не было
	assert.NoError(t, err)

	// Проверяем, что метод был вызван
	mockRepo.AssertExpectations(t)
}
