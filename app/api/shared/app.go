package shared

import (
	"database/sql"
	"fmt"
)

type App struct {
	DB *sql.DB
}

func Init() *App {
	db, err := sql.Open("postgres", "postgres://postgres:password@postgres/reps?sslmode=disable")
	if err != nil {
		panic(fmt.Sprintf("DB initialization error: %s", err.Error()))
	}

	if err = db.Ping(); err != nil {
		panic(fmt.Sprintf("DB connection test error: %s", err.Error()))
	}

	return &App{DB: db}
}
