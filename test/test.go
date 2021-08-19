package main

import (
	"net/http"

	"github.com/Mojashi/regi"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(regi.RegressionTestWithConfig(regi.RegressionTestConfig{
		GoldenDst:  "localhost:1324",
		CurrentDst: "localhost:1323",
		WebUIPort:  ":1325",
	}))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/hello", hello)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
