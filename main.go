package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	echo "github.com/labstack/echo/v4"
)

func main() {
	router := echo.New()

	router.GET("/health-check", healthCheck)

	port := getEnv("ARMANDO_HTTP_PORT", "8080")

	if err := router.Start(fmt.Sprintf(":%s", port)); err != nil {
		log.Fatal(fmt.Sprintf("could'nt start http server: %v", err))
	}
}

func healthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"message": "everything is good!",
	})
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}

	return fallback
}
