package handlers

import (
	"reak/base/pkg/routenames"
	"reak/base/pkg/services"
	"reak/base/pkg/ui/pages"

	"github.com/labstack/echo/v4"
)

type Pages struct{}

func init() {
	Register(new(Pages))
}

func (h *Pages) Init(c *services.Container) error {
	return nil
}

func (h *Pages) Routes(g *echo.Group) {
	g.GET("/", h.Home).Name = routenames.Home
}

func (h *Pages) Home(ctx echo.Context) error {
	return pages.Home(ctx)
}
