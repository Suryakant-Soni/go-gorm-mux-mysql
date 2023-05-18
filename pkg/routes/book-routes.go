package routes

import (
	"go-gorm-mux-mysql/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterBookRoutes = func(router *mux.Router) {
	router.HandleFunc("/book", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book", controllers.GetAllBooks).Methods("GET")
	router.HandleFunc("/book/{bookid}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookid}", controllers.DeleteBook).Methods("DELETE")
	router.HandleFunc("/book/{bookid}", controllers.GetBookById).Methods("GET")
}
