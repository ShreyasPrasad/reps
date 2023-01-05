package shared

import (
	"fmt"
	"log"
	"os"

	"go.temporal.io/sdk/client"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type App struct {
	DB             *gorm.DB
	TemporalClient *client.Client
}

func Init() *App {
	db := initDB()
	client := initTemporalClient()
	return &App{DB: db, TemporalClient: client}
}

func initDB() *gorm.DB {
	db_config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", "db", "postgres", "password", "reps", os.Getenv("DB_PORT"))
	db, err := gorm.Open(postgres.Open(db_config), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("DB initialization error: %s", err.Error()))
	}

	return db
}

func initTemporalClient() *client.Client {
	// Create the client object just once per process
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}

	return &c
}
