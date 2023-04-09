# RESTful-API
A simple RESTful API written in GO using **Echo** and **Gorm** frameworks. 
This app gets musics information, saves them in Postgres, and gets them back using http methods.
## API
* Use POST method with with address `api/music` with json body with this format to add music:
```
{
	"id": "your id",
	"title": "your title",
	"album": "your album",
	"artist": "your artist"
}
```

* Use Get method with with address `api/music/:id` to get a music by its ID.

* Use DELETE method with with address `api/music:id` to delete a music from database by its ID.

* Use GET method with with address `api/album/:album_name` to get all musics in a album.

* Use GET method with with address `api/artist/:artist_name` to get an artist's all musics with name.
