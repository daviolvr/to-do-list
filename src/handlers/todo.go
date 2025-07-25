package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"to-do-list/src/models"

	"github.com/gorilla/mux"
)

// Retorna todas as tasks
func GetTodos(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT id, task, status FROM todos")
		if err != nil {
			http.Error(w, "Erro ao buscar tarefas", 500)
			return
		}
		defer rows.Close()

		var todos []models.ToDoList
		for rows.Next() {
			var todo models.ToDoList
			if err := rows.Scan(&todo.ID, &todo.Task, &todo.Status); err != nil {
				http.Error(w, "Erro ao ler tarefa", 500)
				return
			}
			todos = append(todos, todo)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(todos)
	}
}

func GetTodoByID(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Pega parâmetro "id" da URL
		vars := mux.Vars(r)
		idStr := vars["id"]

		// Converte string para int
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "ID inválido", http.StatusBadRequest)
			return
		}

		var todo models.ToDoList
		err = db.QueryRow("SELECT id, task, status FROM todos where id = $1", id).Scan(&todo.ID, &todo.Task, &todo.Status)

		if err == sql.ErrNoRows {
			http.Error(w, "Task não encontrada", http.StatusNotFound)
			return
		} else if err != nil {
			http.Error(w, "Erro ao buscar tarefa", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(todo)
	}
}

func CreateTodo(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var todo models.ToDoList
		if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
			http.Error(w, "Json inválido", 400)
			return
		}

		err := db.QueryRow(
			"INSERT INTO todos (task, status) VALUES ($1, $2) RETURNING id",
			todo.Task, todo.Status,
		).Scan(&todo.ID)

		if err != nil {
			http.Error(w, "Erro ao inserir tarefa", 500)
			return
		}

		w.Header().Set("Content-TYpe", "application/json")
		json.NewEncoder(w).Encode(todo)
	}
}

func DeleteTodo(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := mux.Vars(r)["id"]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "ID inválido", 400)
			return
		}

		_, err = db.Exec("DELETE FROM todos WHERE id = $1", id)
		if err != nil {
			http.Error(w, "Erro ao deletar tarefa", 500)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
