package main


import (
	"silver-train/db"
	"silver-train/router/http"
)

func main() {
	db.Connect()
	httpRouter.Run()
}
