package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/velvetriddles/fullstack-hidden-chapters/internal/delivery/rest"
	"github.com/velvetriddles/fullstack-hidden-chapters/internal/delivery/rest/handlers"
	"github.com/velvetriddles/fullstack-hidden-chapters/internal/repository/postgres"
	"github.com/velvetriddles/fullstack-hidden-chapters/internal/usecase"
)

func main() {
	db, _ := sql.Open("postgres", "user=myuser password=123 dbname=mydb sslmode=disable host=localhost port=5432")

	defer db.Close()

	bookRepo := postgres.NewBookRepository(db)
	bookUseCase := usecase.NewBookUseCase(bookRepo)
	bookHandler := handlers.NewBookHandler(bookUseCase)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: rest.InitRoutes(bookHandler),
	}
	log.Println("server starts")
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
