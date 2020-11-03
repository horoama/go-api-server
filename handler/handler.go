package handler

import (
    "net/http"
    "github.com/labstack/echo"
)

func Echo() echo.HandlerFunc {
    return func(c echo.Context) error {
        result := c.Param("msg")
        return c.String(http.StatusOK, result)
    }
}
