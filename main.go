package main

import (
	"log"

	"github.com/cmcd97/bytesize/app"
	"github.com/cmcd97/bytesize/auth"
	"github.com/cmcd97/bytesize/middleware"

	"github.com/joho/godotenv"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

func main() {
	if err := godotenv.Load(); err != nil {
		println("Error loading .env file. Please create one or check the spelling of the file name")
		log.Fatal(err)
	}

	pb := pocketbase.New()

	// serves static files from the provided public dir (if exists)
	pb.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.Static("/public", "public")

		authGroup := e.Router.Group("/auth", middleware.LoadAuthContextFromCookie(pb))
		auth.RegisterLoginRoutes(e, *authGroup)
		auth.RegisterRegisterRoutes(e, *authGroup)

		app.InitAppRoutes(e, pb)
		return nil
	})

	if err := pb.Start(); err != nil {
		log.Fatal(err)
	}
}
