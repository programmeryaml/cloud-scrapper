package main

import (
	handler "cloud-scrapper/internal/handler/alibaba"
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	e := echo.New()
	e.GET("/instance/alibaba", handler.GetInstanceInfoHandler)
	e.GET("/instance/alibaba/price", handler.GetPriceInfoHandler)
	e.Logger.Fatal(e.Start(":8080"))
}
