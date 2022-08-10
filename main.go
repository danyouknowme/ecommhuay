package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/danyouknowme/ecommerce-api/pkg/app"
	"github.com/danyouknowme/ecommerce-api/pkg/database"
	"github.com/danyouknowme/ecommerce-api/pkg/util"
	"github.com/gofiber/fiber/v2"
)

func main() {
	fiberApp := app.CreateFiberApp()

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Connot load config:", err.Error())
	}

	dbDriver, dbSource := config.DBDriver, config.DBSource
	database.ConnectDatabase(dbDriver, dbSource)

	fiberApp.Get("/", func(c *fiber.Ctx) error {
		panic("This panic is caught by fiber")
	})

	port := config.Port
	log.Printf("Server starting at port %s.", port)
	fiberApp.Listen(":" + port)
}
