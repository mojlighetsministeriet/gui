package main // import "github.com/mojlighetsministeriet/gui"

import (
	"os"
	"regexp"
	"strings"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/mojlighetsministeriet/utils"
	"github.com/mojlighetsministeriet/utils/server"
)

func main() {
	useTLS := true
	if os.Getenv("TLS") == "disable" {
		useTLS = false
	}
	bodyLimit := utils.GetEnv("BODY_LIMIT", "5M")

	server := server.NewServer(useTLS, true, bodyLimit)

	server.Use(middleware.Static("static"))

	hasFileExtensionPattern := regexp.MustCompile("/[^\\./]+\\.[^\\./]+$")
	server.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:  "static",
		HTML5: true,
		Skipper: func(context echo.Context) bool {
			path := context.Path()
			if strings.HasPrefix(path, "/api/") || hasFileExtensionPattern.MatchString(path) {
				return true
			}

			return false
		},
		Browse: false,
	}))

	server.Listen(":" + utils.GetEnv("PORT", "443"))
}
