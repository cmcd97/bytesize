package app

import (
	"github.com/cmcd97/bytesize/app/handlers"
	"github.com/cmcd97/bytesize/middleware"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

func InitAppRoutes(e *core.ServeEvent, pb *pocketbase.PocketBase) {
	appGroup := e.Router.Group("/app", middleware.LoadAuthContextFromCookie(pb), middleware.AuthGuard)

	appGroup.GET("", func(c echo.Context) error {
		return c.Redirect(303, "profile")
	})
	appGroup.GET("/profile", handlers.ProfileGet)

	e.Router.GET("/", func(c echo.Context) error {
		return c.Redirect(303, "/app")
	})
}
