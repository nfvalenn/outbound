package main

import (
	"log"
	"os"
	"warehouse/handler"
	"warehouse/repository"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	//"github.com/gofiber/fiber/v2/middleware/cors"
	//"gorm.io/gorm/logger"
)

func main() {
	app := fiber.New()

	//app.Use(cors.New)
	//app.Use(logger.New())

	os.Setenv("DATABASE_URL", "postgres://postgres:123@localhost:5432/Outbound")

	outboundRepository := repository.NewOutboundRepository(&gorm.DB{})
	OutboundHandler := handler.NewOutboundHandler(&outboundRepository, app)
	OutboundHandler.SetupRoutes()
	
	log.Println("server is running on port 8080")
	app.Listen(":8080")
}