package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/danyouknowme/ecommerce-api/pkg/app"
	"github.com/danyouknowme/ecommerce-api/pkg/database"
	"github.com/danyouknowme/ecommerce-api/pkg/routes"
	"github.com/danyouknowme/ecommerce-api/pkg/util"
)

func main() {
	fiberApp := app.CreateFiberApp()

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Connot load config:", err.Error())
	}

	dbDriver, dbSource := config.DBDriver, config.DBSource
	database.ConnectDatabase(dbDriver, dbSource)

	routes.SetupRouter(fiberApp)

	port := config.Port
	log.Printf("Server starting at port %s.", port)
	fiberApp.Listen(":" + port)
}
