package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"go-project/database"
)

func main() {
	Environment()
	db, error := database.DBConnection()
	fmt.Println("database", db, error)
}

//Environment variables
func Environment() {
	err := godotenv.Load()
	if err != nil {
		err := godotenv.Load(".env")
		if err != nil {
			panic(".env wasn't found.")
		}
	}
}
