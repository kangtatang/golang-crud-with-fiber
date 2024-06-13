package main

import (
	"fmt"
	"karyawan/config"
	"karyawan/internal/models"
	"karyawan/internal/routes"
	"karyawan/pkg/db"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.InitConfig()
	db.InitDB()

	// Migrasi tabel
	err := db.DB.AutoMigrate(&models.Employee{}, &models.Position{}, &models.Department{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	app := fiber.New()

	// Middleware untuk menangani error
	app.Use(func(c *fiber.Ctx) error {
		err := c.Next()
		if err != nil {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			c.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)
			return c.Status(code).SendString(err.Error())
		}
		return nil
	})

	routes.SetupRoutes(app)

	log.Printf("Starting server at port %d...", config.AppConfig.App.Port)
	if err := app.Listen(fmt.Sprintf(":%d", config.AppConfig.App.Port)); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
