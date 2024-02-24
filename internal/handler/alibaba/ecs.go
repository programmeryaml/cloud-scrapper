package handler

import (
	ecs "cloud-scrapper/internal/alibaba"
	model "cloud-scrapper/internal/repository/alibaba"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetInstanceInfoHandler(c echo.Context) error {
	client, err := ecs.ECS()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to initialize ECS client"})
	}

	instances, err := model.DescribeInstances(client)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to describe instances"})
	}

	return c.JSON(http.StatusOK, instances)
}
