package model

// This is a struct for all musics. The value field is actually base64URL of music binary 
type Music struct {
	ID 			string  `json:"id"`
	Title		string 	`json:"title"`
	Album 		string 	`json:"album"`
	Artist		string  `json:"artist"`
}

