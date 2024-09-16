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
		return c.Redirect(303, "/")
	})
	appGroup.GET("/", handlers.IndexGet)
	// appGroup.GET("/todos", TodosGet(e))
	// appGroup.GET("/todos/add", TodoAddGet)
	// appGroup.POST("/todos/add", TodoAddPost(e))
	// appGroup.POST("/todos/:id/delete", TodoDelete(e))
}
