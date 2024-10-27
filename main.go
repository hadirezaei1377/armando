package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	router := echo.New()

	router.GET("/health-check", healthCheck)
	router.GET("/serve-config-file", serveConfigFile)
	router.GET("/show-me-password", showMePassword)

	staticFilePath := getEnv("ARMANDO_STATIC_PATH", "./static-files")
	router.Static("/statics", staticFilePath)

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

func showMePassword(c echo.Context) error {
	key := getEnv("ARMANDO_SECRET_KEY", "no password")
	return c.JSON(http.StatusOK, echo.Map{
		"secret-key": key,
	})
}

func serveConfigFile(c echo.Context) error {
	filePath := getEnv("ARMANDO_CONFIG_FILE_PATH", "config.yml")

	f, err := os.Open(filePath)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	content, rErr := io.ReadAll(f)
	if rErr != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, rErr.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"data": string(content),
	})
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}

	return fallback
}
