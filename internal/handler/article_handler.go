package handler

import (
	"Posts/internal/dto"
	"Posts/internal/repository"
	"encoding/json"
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
	articles, err := h.repo.GetAllArticles()
	if err != nil {
		http.Error(w, "Failed to fetch articles", http.StatusInternalServerError)
		return
	}

	var response []dto.ArticleResponseDTO
	for _, article := range articles {
		response = append(response, dto.ToArticleResponseDTO(&article))
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// Create
func (h *ArticleHandler) CreateArticleHandler(w http.ResponseWriter, r *http.Request) {
	var articleDTO dto.ArticleRequestDTO

	if err := json.NewDecoder(r.Body).Decode(&articleDTO); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	article := dto.ToArticleModel(articleDTO)

	if err := h.repo.CreateArticle(article); err != nil {
		http.Error(w, "Failed to create article", http.StatusInternalServerError)
		return
	}

	responseDTO := dto.ToArticleResponseDTO(article)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(responseDTO)
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
		http.Error(w, "Article not found", http.StatusNotFound)
		return
	}

	response := dto.ToArticleResponseDTO(article)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// Update Article
func (h *ArticleHandler) UpdateArticleHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var updateDTO dto.ArticleRequestDTO
	if err := json.NewDecoder(r.Body).Decode(&updateDTO); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	article := dto.ToArticleModel(updateDTO)
	article.ID = uint(id)

	if err := h.repo.UpdateArticle(article); err != nil {
		http.Error(w, "Failed to update article", http.StatusInternalServerError)
		return
	}

	response := dto.ToArticleResponseDTO(article)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (h *ArticleHandler) DeleteArticleHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if err := h.repo.DeleteArticle(uint(id)); err != nil {
		http.Error(w, "Failed to delete article", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
