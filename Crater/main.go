package main

import (
	"os"

	router "./router"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	router.Init(r)

	return r
}

func setupConfigs() {
	os.Setenv("CURRENTDOMAIN", "https://ba8d-102-32-68-204.ngrok.io")
	os.Setenv("WEBSERVER_PORT", "8080")
}

func main() {


	r := setupRouter()

	setupConfigs()

	// panic(http.ListenAndServe("8000", nil))

	r.Run(":" + os.Getenv("WEBSERVER_PORT"))
}
