package main

import (
	"os"

	"github.com/raviMukti/gin-rest-starter-project/app"
)

func main() {

	app.Init()
	router := app.SetupRouter()

	router.Run(":" + os.Getenv("APP_PORT"))
}
