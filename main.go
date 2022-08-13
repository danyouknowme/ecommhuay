package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/danyouknowme/ecommerce/pkg/app"
	"github.com/danyouknowme/ecommerce/pkg/database"
	"github.com/danyouknowme/ecommerce/pkg/routes"
	"github.com/danyouknowme/ecommerce/pkg/util"

	"github.com/gofiber/fiber/v2/middleware/cors"

	swagger "github.com/arsmn/fiber-swagger/v2"
	_ "github.com/danyouknowme/ecommerce/docs"
)

func main() {
	fiberApp := app.CreateFiberApp()

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Connot load config:", err.Error())
	}

	fiberApp.Use(cors.New())

	fiberApp.Get("/swagger/*", swagger.HandlerDefault)

	dbDriver, dbSource := config.DBDriver, config.DBSource
	database.ConnectDatabase(dbDriver, dbSource)

	routes.SetupRouter(fiberApp, config.TokenSymmetricKey)

	port := config.Port
	log.Printf("Server starting at port %s.", port)
	fiberApp.Listen(":" + port)
}
