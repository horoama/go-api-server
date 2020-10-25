package handler

import (
    "net/http"
    "github.com/labstack/echo"
)

func Echo() echo.HandlerFunc {
    return func(c echo.Context) error {
        return c.String(http.StatusOK, c.Param("msg"))
    }
}
