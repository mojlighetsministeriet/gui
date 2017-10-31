package main // import "github.com/mojlighetsministeriet/gui"

import (
	"net/http"
	"strings"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/mojlighetsministeriet/utils"
)

func startServer() {
	server := echo.New()

	server.Use()

	server.Use(middleware.Static("client"))
	server.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:    "client",
		HTML5:   true,
		Skipper: noHTML5IfAPICallSkipper,
	}))

	err := server.Start(":" + utils.GetEnv("PORT", "80"))
	if err != nil {
		panic(err)
	}
}

func noHTML5IfAPICallSkipper(context echo.Context) bool {
	if strings.HasPrefix(context.Path(), "/api/") || strings.HasPrefix(context.Path(), "/assets/") || strings.HasPrefix(context.Path(), "/src/") || strings.HasPrefix(context.Path(), "/bower_components/") {
		return true
	}

	return false
}

func respondEmptyBadRequest(context echo.Context) error {
	return context.JSON(http.StatusBadRequest, []byte("{\"message\":\"Bad Request\"}"))
}

func respondOK(context echo.Context, data interface{}) error {
	return context.JSON(http.StatusOK, data)
}

func respondEmptyOK(context echo.Context) error {
	return context.JSONBlob(http.StatusOK, []byte("{\"message\":\"OK\"}"))
}

func respondNotFound(context echo.Context) error {
	return context.JSONBlob(http.StatusNotFound, []byte("{\"message\":\"Not Found\"}"))
}

func respondInternalServerError(context echo.Context) error {
	return context.JSONBlob(http.StatusInternalServerError, []byte("{\"message\":\"Internal Server Error\"}"))
}

func main() {
	startServer()
}
