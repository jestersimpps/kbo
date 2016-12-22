package main

import (
	"net/http"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/kbo/:vatNumber", getKboInfo)
	e.Logger.Fatal(e.Start(":8888"))
}

// e.GET("/kbo/:vatNumber", getKboInfo)
func getKboInfo(c echo.Context) error {
  	// kbo from path `kbo/:vatNumber`
  	vatNumber := c.Param("vatNumber")
	return c.String(http.StatusOK, vatNumber)
}
