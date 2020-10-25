package main

import (
    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
    "api-server/handler"
)

func main() {
    e := echo.New()
    e.Use(middleware.Logger())
    e.GET("/echo/:msg",  handler.Echo())
    e.Start(":1234")
}
