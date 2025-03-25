package routes

import (
	"Posts/internal/handler"
	"net/http"

	"github.com/go-chi/chi"
)

func SetUpRouter(articleHandler *handler.ArticleHandler) *chi.Mux {
	r := chi.NewRouter()

	//Health
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	r.Get("/articles", articleHandler.GetAllArticlesHandler)
	r.Post("/articles", articleHandler.CreateArticleHandler)
	r.Get("/articles/{id}", articleHandler.GetArticleByIDHandler)
	r.Post("/articles/{id}", articleHandler.UpdateArticleHandler)
	r.Delete("/articles/{id}", articleHandler.DeleteArticleHandler)
	return r
}
