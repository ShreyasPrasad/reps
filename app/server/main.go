package main

import (
	"fmt"
	"os"

	routes "github.com/shreyasprasad/reps/app/server/routes"
	shared "github.com/shreyasprasad/reps/app/server/shared"
)

func main() {
	// initialize app state
	app := shared.Init()

	// setup routes, passing in the App state
	router := routes.SetupRoutes(app)
	router.Run(fmt.Sprintf(":%s", os.Getenv("API_PORT")))
}
