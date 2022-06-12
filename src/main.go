package main

import (
	"challenge/src/database"
	"challenge/src/routes"
)

func main() {
	database.ConnectWithDatabase()
	routes.HandleRequest()
}
