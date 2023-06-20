package main

import (
	"adebin/endpoints"
	"adebin/store"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main(){

	// Initialise redis connection
	scl := store.NewClient(store.Config{
		Address: "127.0.0.1:6379",
		Password: "",
		Database: 0,
	})

	e := echo.New()

	e.HideBanner = true
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
	e.Use(middleware.RemoveTrailingSlash())

	endpoint := endpoints.Handler{
		Store: scl,
	}

	api := e.Group("/api")

	// TODO - fetch endpoint to retrieve the stored note
	api.GET("/fetch", endpoint.Fetch)

	// TODO - push endpoint to create a note on adebin

	e.Logger.Fatal(e.Start("0.0.0.0:8080"))

}