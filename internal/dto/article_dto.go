package dto

import (
	"Posts/internal/model"
	"time"
)

type ArticleRequestDTO struct {
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Author  string   `json:"author"`
	Tags    []string `json:"tags"`
}

type ArticleResponseDTO struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Author    string    `json:"author"`
	Tags      []string  `json:"tags"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateArticleDTO struct {
	Title   string   `json:"title" validate:"required,min=3,max=100"`
	Content string   `json:"content" validate:"required,min=10"`
	Author  string   `json:"author" validate:"required"`
	Tags    []string `json:"tags" validate:"dive,required"`
}

func (dto *CreateArticleDTO) ToArticleModel() *model.Article {
	return &model.Article{
		Title:   dto.Title,
		Content: dto.Content,
		Author:  dto.Author,
		Tags:    dto.Tags,
	}
}

func ToArticleResponseDTO(article *model.Article) *ArticleResponseDTO {
	return &ArticleResponseDTO{
		ID:        article.ID,
		Title:     article.Title,
		Content:   article.Content,
		Author:    article.Author,
		Tags:      article.Tags,
		CreatedAt: article.CreatedAt,
	}
}
