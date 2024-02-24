package handler

import (
	bss "cloud-scrapper/internal/alibaba"
	model "cloud-scrapper/internal/repository/alibaba"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetPriceInfoHandler(c echo.Context) error {
	// Initialize the BSS client
	client, err := bss.BSS()
	if err != nil {
		// If there's an error initializing the BSS client, return an internal server error
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to initialize BSS client"})
	}

	// Retrieve price information using the DescribePrices function
	prices, err := model.DescribePrices(client)
	if err != nil {
		// If there's an error retrieving the price information, return an internal server error
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve price information"})
	}

	// Successfully retrieved price information, return it in the response
	return c.JSON(http.StatusOK, prices)
}
