package handler

import (
	"Posts/internal/model"
	"Posts/internal/repository"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type ArticleHandler struct {
	repo *repository.ArticleRepository
}

func NewArticleHandler(repo *repository.ArticleRepository) *ArticleHandler {
	return &ArticleHandler{repo: repo}
}

// GetAll
func (h *ArticleHandler) GetAllArticlesHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем параметры из URL
	title := r.URL.Query().Get("title")
	author := r.URL.Query().Get("author")
	tag := r.URL.Query().Get("tag")
	sort := r.URL.Query().Get("sort")

	// Получаем статьи с фильтрацией и сортировкой
	articles, err := h.repo.GetFilteredAndSortedArticles(title, author, tag, sort)
	if err != nil {
		http.Error(w, "Failed to fetch articles", http.StatusInternalServerError)
		return
	}

	// Отправляем результат в формате JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(articles)
}

// Create
func (h *ArticleHandler) CreateArticleHandler(w http.ResponseWriter, r *http.Request) {
	var article model.Article

	//Парсим тело запроса
	if err := json.NewDecoder(r.Body).Decode(&article); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Сохранение статьи в БД через репозиторий
	if err := h.repo.CreateArticle(&article); err != nil {
		log.Println("DB error:", err)
		http.Error(w, "Failed to create article", http.StatusInternalServerError)
		return
	}

	//Возвращаем ответ
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(article)
}

// Get Article by ID
func (h *ArticleHandler) GetArticleByIDHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	article, err := h.repo.GetArticleByID(uint(id))
	if err != nil {
		http.Error(w, "Failed to get article", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(article)
}

// Update Article
func (h *ArticleHandler) UpdateArticleHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var updateArticle model.Article
	if err := json.NewDecoder(r.Body).Decode(&updateArticle); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	updateArticle.ID = uint(id)

	if err := h.repo.UpdateArticle(&updateArticle); err != nil {
		http.Error(w, "Failed to update", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updateArticle)
}

func (h *ArticleHandler) DeleteArticleHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	if err := h.repo.DeleteArticle(uint(id)); err != nil {
		http.Error(w, "Failed to delete", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
