package endpoints

import (
		"net/http"
		"github.com/labstack/echo/v4"
	db  "github.com/parhamrou/RESTful-API/database"
		"github.com/parhamrou/RESTful-API/model"
)


func AddMusic(c echo.Context) error {
	var music model.Music
	if err := c.Bind(&music); err != nil {
		return err
	}
	if err := db.AddMusic(&music); err != nil {
		return err
	}
	return c.String(http.StatusOK, "You music is saved!")
}


func GetMusic(c echo.Context) error {
	musicName := c.Param("id")
	music, err := db.GetMusic(musicName)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, music)
}


func DeleteMusic(c echo.Context) error {
	musicName := c.Param("id")
	if err := db.DeleteMusic(musicName); err != nil {
		return err
	}
	return c.String(http.StatusOK, "The music is deleted!") 
}


func GetByAlbum(c echo.Context) error {
	albumName := c.Param("album_name")
	album, err := db.GetByAlbum(albumName)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, album)
}


func GetByArtist(c echo.Context) error {
	artistName := c.Param("artist_name")
	musics, err := db.GetByArtist(artistName)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, musics)
}