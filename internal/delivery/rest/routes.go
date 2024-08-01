package rest

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/velvetriddles/fullstack-hidden-chapters/internal/delivery/rest/handlers"
)

func InitRoutes(bookHandler *handlers.BookHandler) *mux.Router {
	r := mux.NewRouter()

	books := r.PathPrefix("/books").Subrouter()
	{
		books.HandleFunc("/{id:[0-9]+}", bookHandler.GetBook).Methods(http.MethodGet)
		books.HandleFunc("/{id:[0-9]+}",bookHandler.UpdateBook).Methods(http.MethodPut)
		books.HandleFunc("/{id:[0-9]+}", bookHandler.DeleteBook).Methods(http.MethodDelete)
		// books.HandleFunc("").Methods(http.MethodGet)
		books.HandleFunc("", bookHandler.CreateBook).Methods(http.MethodPost)
	}
	return r
}
