package main

import (
	"fiber-go-exercise/pkg/config"
	db "fiber-go-exercise/pkg/database"
	"fiber-go-exercise/pkg/handler"
	"fiber-go-exercise/pkg/router"
	"fiber-go-exercise/utils"
	"log"

	"github.com/caarlos0/env/v6"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	utils.InitLogger()
	cfg := config.InitConfig()

	if err := env.Parse(cfg); err != nil {
		log.Fatal("Unable to parse envar, error: ", err)
	}

	newDB, err := db.New(cfg.DBUsername, cfg.DBPassword, cfg.DBPort, cfg.DBHost, cfg.DBName)
	if err != nil {
		log.Fatal("Unable connect to database, error: ", err)
	}
	db.AutoMigrateDB(newDB)

	h := handler.New(newDB)

	app := fiber.New(fiber.Config{
		EnablePrintRoutes: true,
	})
	app.Use(cors.New())

	router.SetupRoutes(h, app)

	log.Fatal(app.Listen(":" + cfg.HTTPPort))

}
