package img_processing

import (
	"database/sql"
	"fmt"
	"log"

	dbs "web-scrapper/internal/db_structure"
)

func SaveImgToFile() {
	dbPath := "database/scrapped_data.db"

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		fmt.Println(err)
		log.Fatalf("Cannot open DB file: %v\n", err)
	}
	// defer db.Close()
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalf("Cannot close DB file: %v\n", err)
		}
	}(db)

	dbRow, err := db.Query("SELECT name, img FROM Scrapped")
	if err != nil {
		fmt.Println(err)
		log.Fatalf("Cannot handle DB request: %v\n", err)
	}
	// defer dbRow.Close()
	defer func(dbRow *sql.Rows) {
		err := dbRow.Close()
		if err != nil {
			log.Fatalf("Cannot close query: %v\n", err)
		}
	}(dbRow)

	for dbRow.Next() {
		var data dbs.ScrappedData
		err = dbRow.Scan(&data.Name, &data.Img)
		if err != nil {
			fmt.Println(err)
			log.Fatalln(err)
		}

		ProcessSaving(data)
	}
}
