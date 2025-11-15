package main

import (
	"log"
	"hih-yadebi-backend/internal/config"
	"hih-yadebi-backend/internal/database"
)

func main() {
	log.Println("Testing config and database...")
	
	// Просто загружаем конфиг и подключаемся к БД
	dbConfig := config.LoadDB()
	db := database.Connect(dbConfig)
	defer db.Close()
	
	log.Println("✅ Everything works!")
}