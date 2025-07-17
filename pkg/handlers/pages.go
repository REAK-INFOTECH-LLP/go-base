package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/pkg/services"
	"github.com/mikestefanello/pagoda/pkg/ui/pages"
)

type Pages struct{}

func init() {
	Register(new(Pages))
}

func (h *Pages) Init(c *services.Container) error {
	return nil
}

func (h *Pages) Routes(g *echo.Group) {
	//g.GET("/", h.Home).Name = routenames.Home
	//g.GET("/about", h.About).Name = routenames.About
}

func (h *Pages) About(ctx echo.Context) error {
	return pages.About(ctx)
}
