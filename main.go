package main

import (
    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
    "api-server/config"
    "api-server/handler"
)

func main() {
    e := echo.New()
    e.Use(middleware.Logger())
    e.GET("/echo/:msg",  handler.Echo())
    conf := config.Load_conf()
    e.Start(":" + conf.Port)
}
