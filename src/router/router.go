package router

import (
	"database/sql"
	"net/http"
	"to-do-list/src/handlers"

	"github.com/gorilla/mux"
)

func NewRouter(db *sql.DB) http.Handler {
	r := mux.NewRouter()

	// Rotas
	r.HandleFunc("/todos/", handlers.GetTodos(db)).Methods("GET")
	r.HandleFunc("/todos/", handlers.CreateTodo(db)).Methods("POST")
	r.HandleFunc("/todos/{id}/", handlers.GetTodoByID(db)).Methods("GET")
	r.HandleFunc("/todos/{id}/", handlers.DeleteTodo(db)).Methods("DELETE")

	return r
}
