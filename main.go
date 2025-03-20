package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Data struct {
	AccID   int
	Name    string
	Balance float64
}

func getMySQL() *sql.DB {
	db, err := sql.Open("mysql", "root:root@(127.0.0.1:3306)/Transaction_db?parseTime=true")
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
		return nil
	}

	if err = db.Ping(); err != nil {
		log.Fatal("Failed to ping database:", err)
		return nil
	}

	return db
}

func main() {
	db := getMySQL()
	if db == nil {
		return
	}
	defer db.Close()

	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS history (
		AccID   INT PRIMARY KEY,
		Name    VARCHAR(255),
		Balance DECIMAL(10,2)
	)`)
	if err != nil {
		log.Fatal("Failed to create table:", err)
	}

	data := []Data{
		{1, "Vishal", 1000.0},
		{2, "Jeel", 2000.0},
		{3, "Abhi", 3000.0},
	}

	for _, row := range data {
		_, err := db.Exec(`INSERT INTO history (AccID, Name, Balance) VALUES (?, ?, ?)`, row.AccID, row.Name, row.Balance)
		if err != nil {
			log.Println("Failed to insert record:", err)
		}
	}
}
