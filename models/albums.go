package models

import "go.mongodb.org/mongo-driver/mongo"

// Upercase variables

type Album struct {
	ID     int     `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

type AlbumBson struct {
	ID     int     `bson:"id"`
	Title  string  `bson:"title"`
	Artist string  `bson:"artist"`
	Price  float64 `bson:"price"`
}


// var Albums = []Album {
// 	{ID: 1, Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
// 	{ID: 2, Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
// 	{ID: 3, Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
// }

func AlbumsCollection(db mongo.Client) *mongo.Collection {
	var coll = db.Database("golang-albums").Collection("albums");
	return coll;
}