package main


import (
	"github.com/joho/godotenv"

	"silver-train/db"
	"silver-train/router/http"
)

func main() {
	godotenv.Load()
	db.Connect()
	httpRouter.Run()
}
