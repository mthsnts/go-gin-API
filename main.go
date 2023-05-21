package main

import (
	"github.com/mthsnts/go-gin-API/database"
	"github.com/mthsnts/go-gin-API/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequest()
}
