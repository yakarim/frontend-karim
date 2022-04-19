package main

import (
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	log "github.com/sirupsen/logrus"
	"github.com/yakarim/frontend-karim/controller"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(encryptcookie.New(encryptcookie.Config{
		Key: encryptcookie.GenerateKey(),
	}))
	controller.Router(app)
	app.Listen(port())

}

func port() string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file")
	}
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = ":3000"
	} else if !strings.HasPrefix(":", port) {
		port = ":" + port
	}
	return port
}
