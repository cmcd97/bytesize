package handlers

import (
	"github.com/cmcd97/bytesize/app/views"
	"github.com/cmcd97/bytesize/lib"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/models"
)

func ProfileGet(c echo.Context) error {
	var record *models.Record = c.Get(apis.ContextAuthRecordKey).(*models.Record)
	return lib.Render(c, 200, views.Profile(record))
}
