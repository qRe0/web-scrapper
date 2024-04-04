package db_processing

import (
	"database/sql"
	"errors"
	"log"

	dbs "web-scrapper/internal/db_structure"
)

func WriteDataToDatabase(arr []dbs.ScrappedData) {
	dbPath := "internal/database/scrapped_data.db"

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatalf("Cannot open DB file: %v\n", err)
	}

	// defer db.Close()
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalf("Cannot close DB file: %v\n", err)
		}
	}(db)

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS Scrapped (
        name TEXT,
        price TEXT,
        url TEXT PRIMARY KEY,
        img TEXT
    )`)
	if err != nil {
		log.Fatalf("Error creating table: %v\n", err)
	}

	existsQuery, err := db.Prepare("SELECT 1 FROM Scrapped WHERE url = ?")
	if err != nil {
		log.Fatalf("Error preparing exists query: %v\n", err)
	}
	// defer existsQuery.Close()
	defer func(existsQuery *sql.Stmt) {
		err := existsQuery.Close()
		if err != nil {
			log.Fatalf("Cannot exit exist query: %v\n", err)
		}
	}(existsQuery)

	tx, err := db.Begin()
	if err != nil {
		log.Fatalf("Error beginning transaction: %v\n", err)
	}
	// defer recover()
	defer func() {
		if err := recover(); err != nil {
			err := tx.Rollback()
			if err != nil {
				log.Fatalf("Cannot recover: %v\n", err)
			}
			log.Fatalf("Recovered from panic: %v\n", err)
		}
	}()

	insertStatement, err := tx.Prepare("INSERT INTO Scrapped (name, price, url, img) VALUES (?, ?, ?, ?)")
	if err != nil {
		log.Fatalf("Error preparing insert statement: %v\n", err)
	}
	// defer insertStatement.Close()
	defer func(insertStatement *sql.Stmt) {
		err := insertStatement.Close()
		if err != nil {
			log.Fatalf("Cannot close transaction statement: %v\n", err)
		}
	}(insertStatement)

	for _, pokemon := range arr {
		var exists bool
		err := existsQuery.QueryRow(pokemon.Url).Scan(&exists)
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			log.Fatalf("Error checking existence: %v\n", err)
		}
		if !exists {
			_, err := insertStatement.Exec(pokemon.Name, pokemon.Price, pokemon.Url, pokemon.Img)
			if err != nil {
				log.Fatalf("Error inserting data: %v\n", err)
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Fatalf("Error committing transaction: %v\n", err)
	}
}
