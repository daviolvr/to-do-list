package main

import (
	"log"
	"net/http"
	"to-do-list/src/db"
	"to-do-list/src/router"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	conn, err := db.ConnectToDB()
	if err != nil {
		log.Fatalf("Erro ao conectar: %v", err)
	}
	defer conn.Close()

	r := router.NewRouter(conn)

	log.Println("Servidor rodando em http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
