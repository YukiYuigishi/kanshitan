package main

import (
	"fmt"
	"github.com/YukiYuigishi/kanshitan-sub/pub/machine"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/api/machine", func(c echo.Context) error {
		r := new(machine.PostRequestMachine)
		log.Println(c.Request().Header)
		if err := c.Bind(r); err != nil {
			log.Println(err)
			return c.String(http.StatusBadRequest, err.Error())
		}
		m, err := machine.NewMachine(*r)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		fmt.Println("name", m.Name)
		fmt.Println("IP", m.IP)
		fmt.Println("uptime", m.UpTime.Seconds())
		fmt.Println("Status", m.Status)
		return c.JSON(http.StatusCreated, r)
	})
	e.Logger.Fatal(e.Start(":1323"))
}
