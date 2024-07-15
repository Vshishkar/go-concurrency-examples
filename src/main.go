package main

import (
	"src/client"
	"src/database"
)

func main() {
	database.RunDb()
	client.RunClient()
}
