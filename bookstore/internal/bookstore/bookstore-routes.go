package bookstore

import (
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", CreateBook).Methods("POST")
	router.HandleFunc("/book/", GetBooks).Methods("GET")
	router.HandleFunc("/book/{bookId}", GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", DeleteBook).Methods("DELETE")
}
