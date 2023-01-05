package shared

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type App struct {
	DB *gorm.DB
}

func Init() *App {
	db_config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", "db", "postgres", "password", "reps", os.Getenv("DB_PORT"))
	db, err := gorm.Open(postgres.Open(db_config), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("DB initialization error: %s", err.Error()))
	}

	return &App{DB: db}
}
