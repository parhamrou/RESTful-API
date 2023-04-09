package main

import (
		"github.com/labstack/echo/v4"
	api "github.com/parhamrou/Music-CloudStorage/endpoints"
	db  "github.com/parhamrou/Music-CloudStorage/database"
)

func main() {
	e := echo.New()
	initialize(e)
	e.Logger.Printf("Listenning on port 8080...\n")
	e.Logger.Fatal(e.Start(":8080"))
}

// This method connects to the database and defines http methods and their handlers
func initialize(e *echo.Echo) error {
	if err := db.Connect(); err != nil {
		return err
	}
	e.POST("/api/music", api.AddMusic)
	e.GET("/api/music/:id", api.GetMusic)
	e.DELETE("api/music/:id", api.DeleteMusic)
	e.GET("/api/album/:album_name", api.GetByAlbum)
	e.GET("api/artist/:artist_name", api.GetByArtist)
	return nil
}