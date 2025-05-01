package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	host := "localhost"
	port := "5433"
	user := "postgres"
	password := "1234"
	dbname := "todo-app"

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("❌ Error opening DB:", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("❌ Cannot connect to DB:", err)
	}
	fmt.Println("✅ Database connected!")

	createTables(DB)
}

func createTables(db *sql.DB) {
	schema := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		username VARCHAR(50) NOT NULL UNIQUE,
		email VARCHAR(100) NOT NULL UNIQUE,
		password TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS todos (
		id SERIAL PRIMARY KEY,
		user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
		title VARCHAR(100) NOT NULL,
		description TEXT,
		status VARCHAR(20) NOT NULL DEFAULT 'pending',
		priority VARCHAR(10) NOT NULL DEFAULT 'medium',
		due_date TIMESTAMP,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS tags (
		id SERIAL PRIMARY KEY,
		name VARCHAR(50) NOT NULL,
		user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		UNIQUE(name, user_id)
	);

	CREATE TABLE IF NOT EXISTS todo_tags (
		todo_id INTEGER NOT NULL REFERENCES todos(id) ON DELETE CASCADE,
		tag_id INTEGER NOT NULL REFERENCES tags(id) ON DELETE CASCADE,
		PRIMARY KEY (todo_id, tag_id)
	);
	`

	if _, err := db.Exec(schema); err != nil {
		log.Fatal("❌ Failed to create tables:", err)
	}
	fmt.Println("✅ Tables created (or already exist).")
}

func CloseDB() {
	if err := DB.Close(); err != nil {
		log.Fatal("❌ Error closing DB:", err)
	}
	fmt.Println("✅ Database connection closed.")
}
func GetDB() *sql.DB {
	return DB
}
