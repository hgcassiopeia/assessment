package drivers

import (
	"database/sql"
	"log"
	"os"
)

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Connect to database error", err)
		return nil, err
	}

	return db, nil
}

func InitTable(db *sql.DB) error {
	createTb := `
	CREATE TABLE IF NOT EXISTS expenses (
		id SERIAL PRIMARY KEY,
		title TEXT,
		amount FLOAT,
		note TEXT,
		tags TEXT[]
	);`

	_, err := db.Exec(createTb)
	if err != nil {
		log.Fatal("Failed to create table", err)
		return err
	}

	return nil
}
