package main

import (
	"Posts/internal/handler"
	"Posts/internal/repository"
	"Posts/internal/routes"
	"fmt"
	"log"
	"net/http"
)

func main() {
	db, err := repository.InitDB()
	if err != nil {
		log.Fatalf("Failed to init database: %v", err)
	}

	//Репозиторий и хэндлер
	articleRepo := repository.NewArticleRepository(db)
	articleHandler := handler.NewArticleHandler(articleRepo)

	//Роутинг
	r := routes.SetUpRouter(articleHandler)

	fmt.Printf("Started on port: 8080")
	http.ListenAndServe(":8080", r)
}
