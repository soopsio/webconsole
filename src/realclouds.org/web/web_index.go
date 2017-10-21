package web

import (
	"github.com/labstack/echo"
	"realclouds.org/middleware"
)

//Index 入口
type Index struct {
}

//MainPage *
func (i *Index) MainPage(c echo.Context) error {
	ctx := c.(*middleware.Context)

	return ctx.ToString("Hello RealClouds.")
}
