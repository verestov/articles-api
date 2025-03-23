// internal/handler_tests/article_handler_test.go
package handler_tests

import (
	"Posts/internal/dto"
	"Posts/internal/handler"
	"Posts/internal/repository"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateArticleHandler(t *testing.T) {
	// Создаем мок-репозиторий
	mockRepo := new(repository.MockArticleRepository)
	articleHandler := handler.NewArticleHandler(mockRepo)

	// Подготавливаем тестовые данные
	testArticle := dto.CreateArticleDTO{
		Title:   "Test Title",
		Author:  "Test Author",
		Content: "Test Content",
	}

	// Настраиваем ожидания для мока
	mockRepo.On("CreateArticle", mock.AnythingOfType("*model.Article")).Return(nil)

	// Создаем новый HTTP-запрос
	reqBody, err := json.Marshal(testArticle)
	if err != nil {
		t.Fatalf("Failed to marshal request body: %v", err)
	}

	req := httptest.NewRequest(http.MethodPost, "/articles", bytes.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Вызовем хендлер
	articleHandler.CreateArticleHandler(w, req)

	// Проверим статус-код
	assert.Equal(t, http.StatusCreated, w.Code)

	// Проверим тело ответа
	var response dto.ArticleResponseDTO
	err = json.NewDecoder(w.Body).Decode(&response)
	if err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	// Проверим корректность возвращаемых данных
	assert.Equal(t, "Test Title", response.Title)
	assert.Equal(t, "Test Author", response.Author)
	assert.Equal(t, "Test Content", response.Content)

	// Проверим, что все ожидаемые вызовы были выполнены
	mockRepo.AssertExpectations(t)
}

func TestCreateArticleHandler_ValidationError(t *testing.T) {
	// Создаем мок-репозиторий
	mockRepo := new(repository.MockArticleRepository)
	articleHandler := handler.NewArticleHandler(mockRepo)

	// Подготавливаем тестовые данные с ошибкой (отсутствует обязательное поле)
	invalidArticle := dto.CreateArticleDTO{
		Title:   "",
		Author:  "Test Author",
		Content: "Test Content",
	}

	// Создаем новый HTTP-запрос
	reqBody, err := json.Marshal(invalidArticle)
	if err != nil {
		t.Fatalf("Failed to marshal request body: %v", err)
	}

	req := httptest.NewRequest(http.MethodPost, "/articles", bytes.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Вызовем хендлер
	articleHandler.CreateArticleHandler(w, req)

	// Проверим статус-код ошибки
	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Проверим, что тело ответа содержит ошибку валидации
	responseBody := w.Body.String()
	assert.True(t, strings.Contains(responseBody, "Validation error"))

	// Проверим, что метод CreateArticle не был вызван
	mockRepo.AssertNotCalled(t, "CreateArticle")
}
