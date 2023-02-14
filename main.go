package main

import (
	"github.com/eferronatto/api-go-gin/database"
	"github.com/eferronatto/api-go-gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
