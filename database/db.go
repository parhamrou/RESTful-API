package database

import (
	"fmt"
	"gorm.io/driver/postgres"
  	"gorm.io/gorm"
	"github.com/parhamrou/Music-CloudStorage/model"
)

const (
	host     = 	"localhost"
	user     = 	"parhamrou"
	password = 	"parham1381"
	dbname   =  "restful_api"
	port     =  5432
)


var db *gorm.DB


func Connect() error {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d", host, user, password, dbname, port)
	db, err = gorm.Open(postgres.Open(dsn))
	return err
}

// Add music to database
func AddMusic(music *model.Music) error {
	return db.Model(model.Music{}).Create(music).Error
}


func GetMusic(id string) (*model.Music, error) {
	var music *model.Music
	result := db.Model(model.Music{}).First(&music, id)
	return music, result.Error
}


func DeleteMusic(id string) error {
	return db.Delete(&model.Music{}, id).Error
}


func GetByAlbum(albumName string) ([]*model.Music, error) {
	var musics []*model.Music
	result := db.Find(&musics, "album = ?", albumName)
	return musics, result.Error
}


func GetByArtist(artistName string) ([]*model.Music, error) {
	var musics []*model.Music
	result := db.Find(&musics, "artist = ?", artistName)
	return musics, result.Error
}