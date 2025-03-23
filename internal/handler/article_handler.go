package handler

import (
	"Posts/internal/dto"
	"Posts/internal/repository"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type ArticleHandler struct {
	repo repository.ArticleRepositoryInterface
}

func NewArticleHandler(repo repository.ArticleRepositoryInterface) *ArticleHandler {
	return &ArticleHandler{repo: repo}
}

// GetAll
func (h *ArticleHandler) GetAllArticlesHandler(w http.ResponseWriter, r *http.Request) {
	articles, err := h.repo.GetAllArticles()
	if err != nil {
		http.Error(w, "Failed to fetch articles", http.StatusInternalServerError)
		return
	}

	var response []*dto.ArticleResponseDTO
	for _, article := range articles {
		response = append(response, dto.ToArticleResponseDTO(&article))
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// Create
func (h *ArticleHandler) CreateArticleHandler(w http.ResponseWriter, r *http.Request) {
	var createDTO dto.CreateArticleDTO
	if err := json.NewDecoder(r.Body).Decode(&createDTO); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := validate.Struct(createDTO); err != nil {
		http.Error(w, fmt.Sprintf("Validation error: %v", err), http.StatusBadRequest)
		return
	}

	article := createDTO.ToArticleModel()

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

	var updateDTO dto.CreateArticleDTO
	if err := json.NewDecoder(r.Body).Decode(&updateDTO); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := validate.Struct(updateDTO); err != nil {
		http.Error(w, fmt.Sprintf("Validation error: %v", err), http.StatusBadRequest)
		return
	}

	article := updateDTO.ToArticleModel()
	article.ID = uint(id)

	if err := h.repo.UpdateArticle(article); err != nil {
		http.Error(w, "Failed to update article", http.StatusInternalServerError)
		return
	}

	response := dto.ToArticleResponseDTO(article)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// Delete
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
