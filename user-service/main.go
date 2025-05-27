package main

import (
	"github.com/MaxSkold/ecommerce-platform/user-service/internal/user/repository"
	"github.com/MaxSkold/ecommerce-platform/user-service/pkg/db"
	"log"
	"net/http"
)

func main() {
	dbConn, err := db.ConnectPSQL()
	if err != nil {
		log.Fatalf("Ошибка подключения к БД: %v", err)
	}
	defer dbConn.Close()

	log.Printf("Успешное подключение к БД")

	userRepo := repository.NewPGRepo(dbConn)
	_ = userRepo

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		if _, errW := w.Write([]byte("user-service OK")); errW != nil {
			log.Printf("<UNK> <UNK> <UNK>: %v", errW)
		}
	})

	log.Println("Сервис запущен на :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
