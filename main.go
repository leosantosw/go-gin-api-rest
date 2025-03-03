package main

import (
	"github.com/leosantosw/database"
	"github.com/leosantosw/routes"
)

func main() {
	database.Connect()
	routes.HandleRequest()
}
