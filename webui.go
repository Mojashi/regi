package regi

import (
	"log"
	"os"
	"os/exec"
	"strconv"
	"sync/atomic"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func serveWebUI(config RegressionTestConfig) {
	e := echo.New()
	e.GET("/api/diffs/:id", getDiff)
	e.GET("/api/diffs", getDiffs)
	e.GET("/api/enabled", getEnabled)
	e.POST("/api/enable", postEnableWatch)
	e.POST("/api/disable", postDisableWatch)

	if _, err := os.Stat(config.StaticFilePos); os.IsNotExist(err) {
		out, err := exec.Command("rm", "/var/tmp/regi", "-rf").Output()
		log.Print(string(out))
		if err != nil {
			log.Fatal(err)
		}
		out, err = exec.Command("git", "clone", "https://github.com/Mojashi/regi.git", "/var/tmp/regi").Output()
		log.Print(string(out))
		if err != nil {
			log.Fatal(err)
		}
		out, err = exec.Command("cp", "-r", "/var/tmp/regi/frontend/regi/build", config.StaticFilePos).Output()
		if err != nil {
			log.Print(config.StaticFilePos)
			log.Fatal(err)
		}
		log.Print(string(out))
		log.Print("successfully initialized!!!")
	}

	e.Use(middleware.Static(config.StaticFilePos))
	go func() {
		e.Logger.Fatal(e.Start(config.WebUIPort))
	}()
}

func getDiff(c echo.Context) error {
	id := c.Param("id")

	dlog := Log{}
	err := db.Get(&dlog, "SELECT * from difflog WHERE id=?", id)
	if err != nil {
		log.Print(err)
		return err
	}
	return c.JSON(200, dlog)
}

func getDiffs(c echo.Context) error {
	st, _ := strconv.Atoi(c.QueryParam("_start"))
	end, _ := strconv.Atoi(c.QueryParam("_end"))
	order := c.QueryParam("_order")
	// sort, err := strconv.Atoi(c.QueryParam("_sort"))

	logs := []Log{}
	err := db.Select(&logs, "SELECT * from difflog ORDER BY id "+order+" LIMIT ? OFFSET ?", end-st, st)
	if err != nil {
		log.Print(err)
		return err
	}
	var totalCount int
	err = db.Get(&totalCount, "SELECT COUNT(1) from difflog")
	if err != nil {
		log.Print(err)
		return err
	}
	c.Response().Header().Set("X-Total-Count", strconv.Itoa(totalCount))
	return c.JSON(200, logs)
}

func postEnableWatch(c echo.Context) error {
	log.Print("enabled diff watching")
	atomic.StoreInt64(&enabled, 1)
	return c.NoContent(200)
}

func getEnabled(c echo.Context) error {
	return c.JSON(200, struct {
		Enabled bool `json:"enabled"`
	}{enabled != 0})
}

func postDisableWatch(c echo.Context) error {
	log.Print("disabled diff watching")
	atomic.StoreInt64(&enabled, 0)
	return c.NoContent(200)
}
