package main

import (
	handler "cloud-scrapper/internal/handler/alibaba"
	"log"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Create a new Echo instance
	e := echo.New()

	// Enable CORS for all origins or restrict to your frontend domain
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"}, // Frontend URL
		AllowMethods: []string{echo.GET, echo.POST},
	}))

	// Define your routes
	e.GET("/instance/alibaba", handler.GetInstanceInfoHandler)
	e.GET("/instance/alibaba/price", handler.GetPriceInfoHandler)

	// Start the server on port 8080
	e.Logger.Fatal(e.Start(":8080"))
}
